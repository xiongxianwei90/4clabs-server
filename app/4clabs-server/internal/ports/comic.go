package ports

import (
	"4clabs-server/app/4clabs-server/internal/domain/entity"
	"context"
)

type Comic interface {
	List(ctx context.Context, userAddress string, limit uint32, nextScore int64) ([]entity.Comic, int64, uint32, bool, error)
	Create(ctx context.Context, comic entity.Comic) error
	GetComicNfts(ctx context.Context, limit uint32, lastScore int64) ([]entity.ComicNft, int64, uint32, bool, error)
	GetComicNftsByComicId(ctx context.Context, comicId string, limit uint32, score int64) ([]entity.ComicNft, int64, uint32, bool, error)
	GetComicAboutMine(ctx context.Context, userAddress string, limit uint32, score int64) ([]entity.Comic, int64, uint32, bool, error)
}

type Trade interface {
	NftPurchase(ctx context.Context, tokenId string, buyerAddress string) error
}
