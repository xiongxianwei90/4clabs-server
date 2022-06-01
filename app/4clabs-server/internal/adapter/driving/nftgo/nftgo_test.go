package nftgo

import (
	"context"
	"testing"
)

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
	ss, err := s.GetNftDetail(context.TODO(), "0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d", "123")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ss)
}
