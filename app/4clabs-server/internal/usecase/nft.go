package usecase

import (
	"4clabs-server/app/4clabs-server/internal/domain/entity"
	"4clabs-server/app/4clabs-server/internal/ports"
	"context"
)

type Nft struct {
	q        ports.Query
	trade    ports.Trade
	register ports.Register
}

func NewNft(q ports.Query, trade ports.Trade, register ports.Register) *Nft {
	return &Nft{q: q, trade: trade, register: register}
}

func (n *Nft) ListRegistedNfts(ctx context.Context, userAddress string, limit uint32, nextScore int64) ([]*entity.Nft, int64, uint32, bool, error) {
	return n.register.ListRegistedNfts(ctx, userAddress, limit, nextScore)
}
func (n *Nft) Register(ctx context.Context, nfts []entity.BaseNft, userAddress string) error {
	return n.register.Register(ctx, nfts, userAddress)
}
func (n *Nft) Registered(ctx context.Context, nftTokenId string, nftContractAddress string, userAddress string) (bool, error) {
	return n.register.Registered(ctx, nftTokenId, nftContractAddress, userAddress)
}
func (n *Nft) GetNftDetail(ctx context.Context, contractAddress, tokenId string) (entity.NftDetail, error) {
	return n.q.GetNftDetail(ctx, contractAddress, tokenId)
}
func (n *Nft) GetComicNftList(ctx context.Context, limit uint32, nextScore int64) ([]entity.ComicNft, int64, uint32, bool, error) {
	return n.q.GetComicNfts(ctx, limit, nextScore)
}
func (n *Nft) NftPurchase(ctx context.Context, tokenId string, buyerAddress string) error {
	return n.trade.NftPurchase(ctx, tokenId, buyerAddress)
}
