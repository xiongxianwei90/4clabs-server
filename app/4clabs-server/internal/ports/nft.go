package ports

import (
	"4clabs-server/app/4clabs-server/internal/domain/entity"
	"context"
)

type Query interface {
	GetAddressNfts(ctx context.Context, address string, limit uint32, lastScore int64) ([]entity.Nft, int64, uint32, bool, error)
	GetNftDetail(ctx context.Context, contractAddress, tokenId string) (entity.NftDetail, error)
}

type Register interface {
	Register(ctx context.Context, nftTokenId string, nftContractAddress string, userAddress string) error
	Registered(ctx context.Context, nftTokenId string, nftContractAddress string, userAddress string) (bool, error)
	UserRegistered(ctx context.Context, userAddress string) (bool, error)
	ListRegistedNfts(ctx context.Context, userAddress string, limit uint32, nextScore int64) ([]entity.Nft, int64, uint32, bool, error)
}
