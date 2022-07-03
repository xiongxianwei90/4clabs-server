package usecase

import (
	"4clabs-server/app/4clabs-server/internal/domain/entity"
	"4clabs-server/app/4clabs-server/internal/ports"
	"context"
)

type Nft struct {
	q        ports.Query
	register ports.Register
}

func NewNft(q ports.Query, register ports.Register) *Nft {
	return &Nft{q: q, register: register}
}

func (n *Nft) ListRegistedNfts(ctx context.Context, userAddress string, limit uint32, nextScore int64) ([]entity.Nft, int64, uint32, bool, error) {
	return n.register.ListRegistedNfts(ctx, userAddress, limit, nextScore)
}
func (n *Nft) Register(ctx context.Context, nftTokenId string, nftContractAddress string, userAddress string) error {
	return n.register.Register(ctx, nftTokenId, nftContractAddress, userAddress)
}
func (n *Nft) Registered(ctx context.Context, nftTokenId string, nftContractAddress string, userAddress string) (bool, error) {
	return n.register.Registered(ctx, nftTokenId, nftContractAddress, userAddress)
}
func (n *Nft) GetNftDetail(ctx context.Context, contractAddress, tokenId string) (entity.NftDetail, error) {
	return n.q.GetNftDetail(ctx, contractAddress, tokenId)
}
