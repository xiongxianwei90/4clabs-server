package repo

import (
	"4clabs-server/app/4clabs-server/internal/adapter/data/model"
	"4clabs-server/app/4clabs-server/internal/adapter/data/query"
	"4clabs-server/app/4clabs-server/internal/adapter/driving/nftgo"
	"4clabs-server/app/4clabs-server/internal/data"
	"4clabs-server/app/4clabs-server/internal/domain/entity"
	"4clabs-server/app/4clabs-server/internal/ports"
	"context"
	"github.com/pkg/errors"
	"time"
)

var _ ports.Comic = &Comic{}

type Comic struct {
	data  *data.Data
	nftgo *nftgo.Service
}

func (c Comic) List(ctx context.Context, userAddress string, limit uint32, nextScore int64) ([]entity.Comic, int64, uint32, bool, error) {
	rnft := query.Use(c.data.DB).Comic
	datas, err := rnft.WithContext(ctx).
		Where(rnft.UserAddress.Eq(userAddress)).
		Where(rnft.CreatedAt.Lt(time.Unix(nextScore, 0))).
		Order(rnft.CreatedAt.Desc()).Limit(int(limit + 1)).Find()
	if err != nil {
		return nil, 0, 0, false, err
	}
	total, err := rnft.WithContext(ctx).Where(rnft.UserAddress.Eq(userAddress)).Count()
	if err != nil {
		return nil, 0, 0, false, err
	}

	if len(datas) == 0 {
		return nil, 0, 0, false, nil
	}

	hasMore := len(datas) > int(limit)
	if hasMore {
		datas = datas[:len(datas)-1]
	}
	lastScore := datas[len(datas)-1].CreatedAt.Unix()

	var infos []struct {
		ContractAddress string
		TokenId         string
	}
	for _, d := range datas {
		infos = append(infos, struct {
			ContractAddress string
			TokenId         string
		}{ContractAddress: d.ContractAddress, TokenId: d.TokenID})
	}

	nfts, err := c.nftgo.BatchGetNftSummary(ctx, infos)
	if err != nil {
		return nil, 0, 0, false, err
	}

	var result []entity.Comic
	for _, d := range datas {
		for _, nft := range nfts {
			if d.ContractAddress == nft.ContractAddress && d.TokenID == nft.TokenId {
				result = append(result, entity.Comic{
					Origin:       nft,
					MintLimit:    uint32(d.MintLimit),
					MintPrice:    d.MintPrice,
					Name:         d.Name,
					MetadataJson: d.Medata,
					CreatedAt:    d.CreatedAt,
					UserAddress:  d.UserAddress,
				})
			}
		}
	}

	return result, lastScore, uint32(total), hasMore, nil
}

func (c Comic) Create(ctx context.Context, comic entity.Comic) error {
	db := query.Use(c.data.DB).Comic
	if err := db.WithContext(ctx).Create(&model.Comic{
		TokenID:         comic.Origin.TokenId,
		ContractAddress: comic.Origin.ContractAddress,
		UserAddress:     comic.UserAddress,
		MintLimit:       int32(comic.MintLimit),
		MintPrice:       comic.MintPrice,
		Medata:          comic.MetadataJson,
	}); err != nil {
		return errors.Wrapf(err, "model : %#v", comic)
	}
	return nil
}

func NewComic(data *data.Data, nftgo *nftgo.Service) *Comic {
	return &Comic{data: data, nftgo: nftgo}
}
