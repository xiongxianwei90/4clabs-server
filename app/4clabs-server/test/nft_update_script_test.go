package test

import (
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

func TestRegisterNftUpdate(t *testing.T) {
	dataData, bc, err := GetEnvironment()
	cacheNfts := nft.NewNft(dataData)
	s := nftgo.NewService(&bc, cacheNfts, dataData)
	register := repo.NewRegister(s, dataData)
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
			nftgo.RegisterNftUpdate(context.TODO(), dataData.DB, item.TokenId, item.ContractAddress, detail.Stat.LastPrice.PriceToken)
		}
		if !hasMore {
			break
		}
	}

}

func TestComicNftAllUpdate(t *testing.T) {
	dataData, bc, err := GetEnvironment()
	err = nftgo.ComicUpdate(context.TODO(), dataData.DB, bc.ThirdParty.Contract.Rawurl, bc.ThirdParty.Contract.Address, false)
	//err = ComicNftUpdate(context.TODO(), dataData.DB, bc.ThirdParty.Contract.Rawurl, bc.ThirdParty.Contract.Address, false)
	if err != nil {
		t.Fail()
	}
}

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
