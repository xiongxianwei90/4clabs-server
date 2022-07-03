package usecase

import (
	"4clabs-server/app/4clabs-server/internal/domain/entity"
	"4clabs-server/app/4clabs-server/internal/ports"
	"context"
)

type Comics struct {
	ports.Comic
}

func NewComics(comic ports.Comic) *Comics {
	return &Comics{Comic: comic}
}

func (c *Comics) List(ctx context.Context, userAddress string, limit uint32, nextScore int64) ([]entity.Comic, int64, uint32, bool, error) {
	return c.Comic.List(ctx, userAddress, limit, nextScore)
}
