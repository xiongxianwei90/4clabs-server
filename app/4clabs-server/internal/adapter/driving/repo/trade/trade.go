package trade

import (
	"4clabs-server/app/4clabs-server/internal/adapter/driving/nftgo"
	"4clabs-server/app/4clabs-server/internal/data"
	"4clabs-server/app/4clabs-server/internal/ports"
	"context"
)

var _ ports.Trade = &Trade{}

type Trade struct {
	data  *data.Data
	nftgo *nftgo.Service
}

func NewTrade(data *data.Data, nftgo *nftgo.Service) *Trade {
	return &Trade{data: data, nftgo: nftgo}
}

func (t Trade) NftPurchase(ctx context.Context, tokenId string, buyerAddress string) error {
	err := nftgo.ComicNftUpdateById(ctx, t.data.DB, tokenId, buyerAddress)
	return err
}
