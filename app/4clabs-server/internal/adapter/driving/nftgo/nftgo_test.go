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
