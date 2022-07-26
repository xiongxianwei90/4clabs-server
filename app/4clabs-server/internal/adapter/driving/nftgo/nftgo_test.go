package nftgo

import (
	"4clabs-server/app/4clabs-server/internal/conf"
	"4clabs-server/app/4clabs-server/internal/data"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"os"
	"testing"
)

func TestComicNftAllUpdate(t *testing.T) {
	var environment = "local"
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
	err = ComicUpdate(context.TODO(), dataData.DB, bc.ThirdParty.Contract.Rawurl, bc.ThirdParty.Contract.Address, false)
	//err = ComicNftUpdate(context.TODO(), dataData.DB, bc.ThirdParty.Contract.Rawurl, bc.ThirdParty.Contract.Address, false)
	if err != nil {
		t.Fail()
	}
}

func TestAddressGetNft(t *testing.T) {
	s := Service{apiKey: "a12366dfa8e44fc99d7667765cf3608e"}
	_, _, _, _, err := s.GetAddressNfts(context.TODO(), "0x735f781bd2B322CD0557e689cA912A3dA279B902", 20, 0)
	if err != nil {
		t.Fail()
	}
}
func TestGetSummary(t *testing.T) {
	s := Service{apiKey: "a12366dfa8e44fc99d7667765cf3608e"}
	ss, err := s.GetNftSummary(context.TODO(), "0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d", "123")
	if err != nil {
		t.Fail()
	}
	t.Log(ss)
}

func TestGetDetail(t *testing.T) {
	s := Service{apiKey: "a12366dfa8e44fc99d7667765cf3608e"}
	ss, err := s.GetNftDetail(context.TODO(), "0x09233d553058c2f42ba751c87816a8e9fae7ef10", "5726")
	if err != nil {
		t.Fatalf("err : %s", err)
	}
	t.Log(ss)
}
