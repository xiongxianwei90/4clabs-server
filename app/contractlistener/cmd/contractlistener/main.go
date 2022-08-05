package main

import (
	script "4clabs-server/api/script/v1"
	forClabs "4clabs-server/app/4clabs-server/pkg/contract"
	"4clabs-server/app/contractlistener/internal/conf"
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/pkg/errors"
	"io"
	"math/big"
	"net/http"
	"os"
	"time"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "4clabs-contract-service"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	environment = os.Getenv("environment")
	id, _       = os.Hostname()
)

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(fmt.Sprintf("./configs/config.%s.yaml", environment)),
		),
	)

	defer c.Close()
	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, _, err := wireApp(&bc, logger)
	if err != nil {
		panic(err)
	}

	begin := time.Now()
	if err := app.Run(); err != nil {
		panic(err)
	}
	fmt.Printf("the task cost : %s\n", time.Now().Sub(begin).String())

}

type ContractScriptApplication struct {
	bc *conf.Bootstrap
}

func NewCronJobApplication(bc *conf.Bootstrap) *ContractScriptApplication {
	return &ContractScriptApplication{bc: bc}
}

func (app *ContractScriptApplication) Run() error {
	//root := context.Background()
	println()
	startLog := ":'######:::'#######::'##::: ##:'########:'########:::::'###:::::'######::'########:\n'##... ##:'##.... ##: ###:: ##:... ##..:: ##.... ##:::'## ##:::'##... ##:... ##..::\n ##:::..:: ##:::: ##: ####: ##:::: ##:::: ##:::: ##::'##:. ##:: ##:::..::::: ##::::\n ##::::::: ##:::: ##: ## ## ##:::: ##:::: ########::'##:::. ##: ##:::::::::: ##::::\n ##::::::: ##:::: ##: ##. ####:::: ##:::: ##.. ##::: #########: ##:::::::::: ##::::\n ##::: ##: ##:::: ##: ##:. ###:::: ##:::: ##::. ##:: ##.... ##: ##::: ##:::: ##::::\n. ######::. #######:: ##::. ##:::: ##:::: ##:::. ##: ##:::: ##:. ######::::: ##::::\n:......::::.......:::..::::..:::::..:::::..:::::..::..:::::..:::......::::::..:::::\n:::'##:::::::'####::'######::'########:'########:'##::: ##:'########:'########:::::\n::: ##:::::::. ##::'##... ##:... ##..:: ##.....:: ###:: ##: ##.....:: ##.... ##::::\n::: ##:::::::: ##:: ##:::..::::: ##:::: ##::::::: ####: ##: ##::::::: ##:::: ##::::\n::: ##:::::::: ##::. ######::::: ##:::: ######::: ## ## ##: ######::: ########:::::\n::: ##:::::::: ##:::..... ##:::: ##:::: ##...:::: ##. ####: ##...:::: ##.. ##::::::\n::: ##:::::::: ##::'##::: ##:::: ##:::: ##::::::: ##:. ###: ##::::::: ##::. ##:::::\n::: ########:'####:. ######::::: ##:::: ########: ##::. ##: ########: ##:::. ##::::\n:::........::....:::......::::::..:::::........::..::::..::........::..:::::..:::::"
	println(startLog)
	println()
	client, err := ethclient.Dial(app.bc.ThirdParty.Contract.Rawurl)
	if err != nil {
		return err
	}

	contract := common.HexToAddress(app.bc.ThirdParty.Contract.Address)
	instance, err := forClabs.NewForClabs(contract, client)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(app.bc.ThirdParty.Contract.FromBlock),
		Addresses: []common.Address{contract},
	}

	// 如果需要监听单独事件可使用
	// instance.WatchContractAndTokenAuthorized()
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println("TxHash:", vLog.TxHash)
			cata, _ := instance.ParseContractAndTokenAuthorized(vLog)
			if cata != nil {
				request := script.ScriptRegisterRequest{
					ContractAddress: cata.ContractAddress.String(),
					TokenId:         cata.TokenId.String(),
					UserAddress:     cata.Holder.String(),
				}
				jsonBytes, _ := json.Marshal(request)
				println("register data:", string(jsonBytes))
				url := fmt.Sprintf("%s/v1/script/register/update", app.bc.ThirdParty.ScriptUrl)
				_, err = app.basePost(context.Background(), url, jsonBytes)
				if err != nil {
					println("post error: ", fmt.Sprint(err))
				}
			}
			fmt.Println()
		}
	}
	println(" ----------- end -----------")
	return nil
}

func (s *ContractScriptApplication) basePost(ctx context.Context, url string, json []byte) (io.ReadCloser, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, errors.Wrapf(err, "url : %s", url)
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(fmt.Errorf("status code %d is not ok", res.StatusCode), "url : %s", url)
	}
	return res.Body, nil
}
