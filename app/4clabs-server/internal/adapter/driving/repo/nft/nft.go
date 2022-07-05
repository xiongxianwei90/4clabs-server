package nft

import (
	"4clabs-server/app/4clabs-server/internal/adapter/data/query"
	"4clabs-server/app/4clabs-server/internal/data"
	"context"
)

type Iden struct {
	ContractAddress string
	TokenId         string
}

type Nft struct {
	data *data.Data
}

func NewNft(data *data.Data) *Nft {
	return &Nft{data: data}
}

func (n *Nft) Image(ctx context.Context, idens ...string) (map[string]string, error) {
	q := query.Use(n.data.DB).Nft
	datas, err := q.WithContext(ctx).Where(q.ContractAddressTokenID.In(idens...)).Find()
	if err != nil {
		return nil, err
	}
	result := make(map[string]string)
	for _, r := range datas {
		result[r.ContractAddressTokenID] = r.Image
	}
	return result, nil
}
