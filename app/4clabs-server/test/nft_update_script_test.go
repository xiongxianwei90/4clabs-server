package test

import (
	"4clabs-server/app/4clabs-server/internal/adapter/data/model"
	"4clabs-server/app/4clabs-server/internal/adapter/data/query"
	"4clabs-server/app/4clabs-server/internal/adapter/driving/nftgo"
	"4clabs-server/app/4clabs-server/internal/adapter/driving/repo"
	"4clabs-server/app/4clabs-server/internal/adapter/driving/repo/nft"
	"4clabs-server/app/4clabs-server/internal/conf"
	"4clabs-server/app/4clabs-server/internal/data"
	"4clabs-server/app/4clabs-server/internal/domain/entity"
	forClabs "4clabs-server/app/4clabs-server/pkg/contract"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"math/big"
	"os"
	"strings"
	"testing"
)

func GetEnvironment() (*data.Data, conf.Bootstrap, error) {
	var environment = "production"
	c := config.New(
		config.WithSource(
			file.NewSource(fmt.Sprintf("/Users/xiongwei/GolandProjects/4clabs-server/app/4clabs-server/configs/config.%s.yaml", environment)),
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
	var (
		id, _   = os.Hostname()
		Name    string
		Version string
	)

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	dataData, _, err := data.NewData(&bc, logger)
	return dataData, bc, err
}

// 更新已经登记过的nft 价格图片等
func TestRegisterNftUpdate(t *testing.T) {
	dataData, bc, err := GetEnvironment()
	cacheNfts := nft.NewNft(dataData)
	s := nftgo.NewService(&bc, cacheNfts, dataData)
	register := repo.NewRegister(s, dataData, cacheNfts)
	if err != nil {
		t.Fail()
	}
	var i int64 = 0
	var list []*entity.Nft
	var hasMore = true
	for {
		list, i, _, hasMore, err = register.ListRegistedNfts(context.TODO(), "", nil, 50, i)
		for _, item := range list {
			detail, err := s.GetNftDetail(context.TODO(), item.ContractAddress, item.TokenId)
			if err != nil {
				t.Fail()
			}
			if detail.Summary.Name == "" {
				continue
			}
			nftgo.RegisterNftUpdate(
				context.TODO(),
				dataData.DB,
				detail.Summary.TokenId,
				detail.Summary.Name,
				detail.Summary.CollectionName,
				detail.Summary.ContractAddress,
				detail.Stat.LastPrice.PriceToken,
				detail.Summary.Image,
			)
		}
		if !hasMore {
			break
		}
	}

}

func TestComicNftAllUpdate(t *testing.T) {
	dataData, bc, err := GetEnvironment()
	err = nftgo.ComicUpdate(context.TODO(), dataData.DB, bc.ThirdParty.Contract.Rawurl, bc.ThirdParty.Contract.Address, false)
	err = nftgo.ComicNftUpdate(context.TODO(), dataData.DB, bc.ThirdParty.Contract.Rawurl, bc.ThirdParty.Contract.Address, false)
	if err != nil {
		t.Fail()
	}
}

// tokenUri解析
func TestTokenId(t *testing.T) {
	_, bc, _ := GetEnvironment()
	client, err := ethclient.Dial(bc.ThirdParty.Contract.Rawurl)
	if err != nil {
		t.Fail()
	}
	contract := common.HexToAddress(bc.ThirdParty.Contract.Address)
	instance, err := forClabs.NewForClabs(contract, client)
	tokenUri, err := instance.TokenURI(nil, big.NewInt(0))
	tokenUri = strings.Replace(tokenUri, "data:application/json;base64,", "", 1)
	println(tokenUri)
	decoder := Base64Decode(tokenUri)
	println(decoder)
}

// 验证是否登记过
func TestIsAuthorized(t *testing.T) {
	_, bc, _ := GetEnvironment()
	client, err := ethclient.Dial(bc.ThirdParty.Contract.Rawurl)
	if err != nil {
		t.Fail()
	}
	contract := common.HexToAddress(bc.ThirdParty.Contract.Address)
	instance, err := forClabs.NewForClabs(contract, client)
	// 5438 1041 2413 0xD10b3AAe17C7a0950Dca17D0F21D23dae368c283
	isAuthorized, err := instance.IsContractAndTokenAuthorized(nil, common.HexToAddress("0x9401518f4EBBA857BAA879D9f76E1Cc8b31ed197"), big.NewInt(380))
	println(isAuthorized)
}

// 从合约更新所有登记过的nft
func TestGetContractAndTokenAuthorized(t *testing.T) {
	dataData, bc, err := GetEnvironment()
	cacheNfts := nft.NewNft(dataData)
	s := nftgo.NewService(&bc, cacheNfts, dataData)
	register := repo.NewRegister(s, dataData, cacheNfts)
	client, err := ethclient.Dial(bc.ThirdParty.Contract.Rawurl)
	if err != nil {
		t.Fail()
	}

	contract := common.HexToAddress(bc.ThirdParty.Contract.Address)
	instance, err := forClabs.NewForClabs(contract, client)
	filterQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(bc.ThirdParty.Contract.FromBlock),
		Addresses: []common.Address{
			common.HexToAddress(bc.ThirdParty.Contract.Address),
		},
	}
	logs, err := client.FilterLogs(context.Background(), filterQuery)
	if err != nil {
		log.Fatal(err)
	}

	// 或者只筛选FilterContractAndTokenAuthorized 日志更简便
	//filterOpts := &bind.FilterOpts{
	//	Start:   uint64(bc.ThirdParty.Contract.FromBlock),
	//	Context: context.Background(),
	//}
	//auther, err := instance.FilterContractAndTokenAuthorized(filterOpts)
	//for auther.Next() {
	//	event := auther.Event
	//	println(fmt.Printf(event.Holder.String()))
	//}

	authorized := make(map[string]*forClabs.ForClabsContractAndTokenAuthorized)
	for _, item := range logs {
		forclabsLog, _ := instance.ParseContractAndTokenAuthorized(item)
		if forclabsLog == nil {
			continue
		}
		key := fmt.Sprintf("%s_%s", strings.ToUpper(forclabsLog.ContractAddress.String()), forclabsLog.TokenId.String())
		authorized[key] = forclabsLog
	}
	println("合约授权数量", len(authorized))

	var i int64 = 0
	var list []*entity.Nft
	var hasMore = true

	for {
		list, i, _, hasMore, err = register.ListRegistedNfts(context.TODO(), "", nil, 1000, i)
		for _, item := range list {
			key := fmt.Sprintf("%s_%s", strings.ToUpper(item.ContractAddress), item.TokenId)
			_, exists := authorized[key]
			if exists {
				delete(authorized, key)
				continue
			}
		}
		if !hasMore {
			break
		}
	}

	println("需要更新的数量", len(authorized))

	for _, item := range authorized {
		println("update --", item.ContractAddress.String(), item.TokenId.String())

		detail, err := s.GetNftDetail(context.TODO(), item.ContractAddress.String(), item.TokenId.String())
		if err != nil {
			t.Fail()
		}
		registerNft := query.Use(dataData.DB).RegisterNft

		registerNft.WithContext(context.TODO()).Create(&model.RegisterNft{
			TokenID:         item.TokenId.String(),
			Name:            detail.Summary.Name,
			CollectionName:  detail.Summary.CollectionName,
			ContractAddress: item.ContractAddress.String(),
			UserAddress:     item.Holder.String(),
			Price:           detail.Stat.LastPrice.PriceToken,
			Image:           detail.Summary.Image,
		})
	}
}

func TestParseLog(t *testing.T) {

}

func Base64Decode(str string) string {
	reader := strings.NewReader(str)
	decoder := base64.NewDecoder(base64.RawStdEncoding, reader)
	// 以流式解码
	buf := make([]byte, 1024)
	// 保存解码后的数据
	dst := ""
	for {
		n, err := decoder.Read(buf)
		dst += string(buf[:n])
		if n == 0 || err != nil {
			break
		}
	}
	return dst
}
