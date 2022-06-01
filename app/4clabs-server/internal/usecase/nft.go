package usecase

import (
	"4clabs-server/app/4clabs-server/internal/domain/entity"
	"4clabs-server/app/4clabs-server/internal/ports"
	"context"
)

type Nft struct {
	q ports.Query
}

func (n *Nft) GetNftDetail(ctx context.Context, contractAddress, tokenId string) (entity.NftDetail, error) {
	return n.q.GetNftDetail(ctx, contractAddress, tokenId)
}
