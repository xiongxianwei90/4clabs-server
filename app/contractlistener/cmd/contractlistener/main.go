package main

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
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

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
	)
}

func main() {
	//flag.Parse()
	//logger := log.With(log.NewStdLogger(os.Stdout),
	//	"ts", log.DefaultTimestamp,
	//	"caller", log.DefaultCaller,
	//	"service.id", id,
	//	"service.name", Name,
	//	"service.version", Version,
	//	"trace_id", tracing.TraceID(),
	//	"span_id", tracing.SpanID(),
	//)
	//c := config.New(
	//	config.WithSource(
	//		file.NewSource(fmt.Sprintf("./configs/config.%s.yaml", environment)),
	//	),
	//)
	//
	//defer c.Close()
	//
	//if err := c.Load(); err != nil {
	//	panic(err)
	//}
	//
	//var bc conf.Bootstrap
	//if err := c.Scan(&bc); err != nil {
	//	panic(err)
	//}
	//authUtitls := auth.NewJwtUtils(
	//	bc.Server.Http.Jwt.Key,
	//	bc.Server.Http.Jwt.Timeout.Seconds,
	//	bc.Server.Http.Jwt.MinRefresh.Seconds,
	//)
	//
	//app, cleanup, err := wireApp(&bc, logger, authUtitls, auth.NewUtils())
	//if err != nil {
	//	panic(err)
	//}
	//defer cleanup()
	//
	//// start and wait for stop signal
	//if err := app.Run(); err != nil {
	//	panic(err)
	//}

	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws/v3/011c5e6aa1fc44daab9ab5e51baeb9fb")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x78cebaf8498a34451e50dcda7f0ee568db89ae39")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(client.BlockNumber(context.TODO()))

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}

}
