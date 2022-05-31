package usecase

import (
	"4clabs-server/app/4clabs-server/internal/domain/entity"
	"4clabs-server/app/4clabs-server/internal/domain/service"
	"context"
)

type Address struct {
	q service.Query
}

func NewAddress(q service.Query) *Address {
	return &Address{q: q}
}

func (a *Address) GetAddressNfts(ctx context.Context, address string, limit uint32, lastScore int64) ([]entity.Nft, int64, uint32, bool, error) {
	return a.q.GetAddressNfts(ctx, address, limit, lastScore)
}
