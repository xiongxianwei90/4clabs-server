package usecase

import (
	"4clabs-server/app/4clabs-server/internal/domain/entity"
	"4clabs-server/app/4clabs-server/internal/ports"
	"context"
)

type Comics struct {
	ports.Comic
	trade ports.Trade
}

func NewComics(comic ports.Comic, trade ports.Trade) *Comics {
	return &Comics{Comic: comic, trade: trade}
}

func (c *Comics) List(ctx context.Context, userAddress string, limit uint32, nextScore int64) ([]entity.Comic, int64, uint32, bool, error) {
	return c.Comic.List(ctx, userAddress, limit, nextScore)
}
func (c *Comics) Create(ctx context.Context, comic entity.Comic) error {
	return c.Comic.Create(ctx, comic)
}

func (c *Comics) GetComicNftList(ctx context.Context, limit uint32, nextScore int64) ([]entity.ComicNft, int64, uint32, bool, error) {
	return c.Comic.GetComicNfts(ctx, limit, nextScore)
}

func (c *Comics) NftPurchase(ctx context.Context, tokenId string, buyerAddress string) error {
	return c.trade.NftPurchase(ctx, tokenId, buyerAddress)
}

func (c *Comics) GetComicNftListByComicId(ctx context.Context, id string, limit uint32, nextScore int64) ([]entity.ComicNft, int64, uint32, bool, error) {
	return c.Comic.GetComicNftsByComicId(ctx, id, limit, nextScore)
}
