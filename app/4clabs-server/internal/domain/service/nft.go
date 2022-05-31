package service

import (
	"4clabs-server/app/4clabs-server/internal/domain/entity"
	"context"
)

type Query interface {
	GetAddressNfts(ctx context.Context, address string, limit uint32, lastScore int64) ([]entity.Nft, int64, uint32, bool, error)
}
