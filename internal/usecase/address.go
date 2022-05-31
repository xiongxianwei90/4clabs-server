package usecase

import (
	"4clabs-server/internal/domain/entity"
	"context"
)

type Address struct {
}

func (a *Address) GetAddressNfts(ctx context.Context, address string, limit uint32, lastScore int64) ([]entity.Nft, int64, bool, error) {
}
