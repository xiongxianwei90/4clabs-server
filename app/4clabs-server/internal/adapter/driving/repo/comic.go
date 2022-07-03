package repo

import (
	"4clabs-server/app/4clabs-server/internal/adapter/data/model"
	"4clabs-server/app/4clabs-server/internal/adapter/data/query"
	"4clabs-server/app/4clabs-server/internal/adapter/driving/nftgo"
	"4clabs-server/app/4clabs-server/internal/data"
	"4clabs-server/app/4clabs-server/internal/domain/entity"
	"4clabs-server/app/4clabs-server/internal/ports"
	"context"
)

var _ ports.Comic = &Comic{}

type Comic struct {
	data  *data.Data
	nftgo *nftgo.Service
}

func (c Comic) List(ctx context.Context, userAddress string, limit uint32, nextScore int64) ([]entity.Comic, int64, uint32, bool, error) {
	query.Use(c.data.DB).Comic

	panic("implement me")
}

func (c Comic) Create(ctx context.Context, comic entity.Comic) {
	db := query.Use(c.data.DB).Comic
	db.WithContext(ctx).Create(&model.Comic{
		TokenID:         comic.Origin.TokenId,
		ContractAddress: comic.Origin.ContractAddress,
		UserAddress:     comic.UserAddress,
		MintLimit:       int32(comic.MintLimit),
		MintPrice:       comic.MintPrice,
		Medata:          comic.,
	})

	panic("implement me")
}

func NewComic(data *data.Data, nftgo *nftgo.Service) *Comic {
	return &Comic{data: data, nftgo: nftgo}
}
