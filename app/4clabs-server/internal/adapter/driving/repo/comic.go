package repo

import (
	"4clabs-server/app/4clabs-server/internal/adapter/data/model"
	"4clabs-server/app/4clabs-server/internal/adapter/data/query"
	"4clabs-server/app/4clabs-server/internal/adapter/driving/nftgo"
	"4clabs-server/app/4clabs-server/internal/data"
	"4clabs-server/app/4clabs-server/internal/domain/entity"
	"4clabs-server/app/4clabs-server/internal/ports"
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var _ ports.Comic = &Comic{}

type Comic struct {
	data  *data.Data
	nftgo *nftgo.Service
}

func (c Comic) GetComicAboutMine(ctx context.Context, userAdderss string, limit uint32, score int64) ([]entity.Comic, int64, uint32, bool, error) {
	return nil, 0, 0, false, nil
}

func (c Comic) List(ctx context.Context, userAddress string, limit uint32, nextScore int64) ([]entity.Comic, int64, uint32, bool, error) {
	if nextScore == 0 {
		nextScore = time.Now().Unix()
	}
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

	comicMap := make(map[string]entity.Nft)
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
	for _, item := range nfts {
		comicMap[item.TokenId] = item
	}
	var result []entity.Comic
	for _, d := range datas {
		nft := comicMap[d.TokenID]
		imageUris := strings.Split(d.ImageURIs, ",")
		result = append(result, entity.Comic{
			Origin:      nft,
			ComicId:     fmt.Sprint(int(d.ID)),
			MintLimit:   uint32(d.MintLimit),
			MintPrice:   d.MintPrice,
			Name:        d.Name,
			CreatedAt:   d.CreatedAt,
			UserAddress: d.UserAddress,
			ImageUris:   imageUris,
		})
	}

	return result, lastScore, uint32(total), hasMore, nil
}

func (c Comic) Create(ctx context.Context, comic entity.Comic) error {
	nftgo.ComicUpdate(ctx, c.data.DB, c.nftgo.Rawurl, c.nftgo.ContractAddress, true)
	nftgo.ComicNftUpdate(ctx, c.data.DB, c.nftgo.Rawurl, c.nftgo.ContractAddress, true)
	return nil
}

func (c Comic) GetComicNfts(ctx context.Context, limit uint32, lastScore int64) ([]entity.ComicNft, int64, uint32, bool, error) {
	var comicNfts []entity.ComicNft

	rnft := query.Use(c.data.DB).ComicsNft
	query := rnft.WithContext(ctx).
		Where(rnft.Owner.EqCol(rnft.Author))
	datas, err := query.
		Offset(int(lastScore)).
		Limit(int(limit)).Preload(rnft.Comic).Find()
	if err != nil {
		return nil, 0, 0, false, err
	}
	totalSupply, err := query.Count()
	if err != nil {
		return nil, 0, 0, false, err
	}

	var infos []struct {
		ContractAddress string
		TokenId         string
	}
	worksComic := make(map[int]*model.Comic)

	for index, d := range datas {
		infos = append(infos, struct {
			ContractAddress string
			TokenId         string
		}{ContractAddress: d.Comic.ContractAddress, TokenId: d.Comic.TokenID})
		worksComic[index] = &d.Comic
	}

	origins, err := c.nftgo.BatchGetNftSummary(ctx, infos)
	originNfts := make(map[string]entity.Nft)

	for _, item := range origins {
		originNfts[item.TokenId] = item
	}

	for tokenId, item := range datas {
		id := strconv.FormatInt(int64(item.ID), 10)
		comicNfts = append(comicNfts, entity.ComicNft{
			TokenId: id,
			Owner:   item.Owner,
			Comic: entity.Comic{
				Origin:      originNfts[worksComic[tokenId].TokenID],
				MintLimit:   uint32(item.Comic.MintLimit),
				MintPrice:   item.Comic.MintPrice,
				Name:        item.Comic.Name,
				CreatedAt:   item.Comic.CreatedAt,
				UserAddress: item.Comic.UserAddress,
				ImageUris:   strings.Split(item.Comic.ImageURIs, ","),
			},
		})
	}
	nextLastScore := int64(limit) + lastScore
	return comicNfts, nextLastScore, uint32(totalSupply), nextLastScore < totalSupply, nil
}

func (c Comic) GetComicNftsByComicId(ctx context.Context, comicId string, limit uint32, lastScore int64) ([]entity.ComicNft, int64, uint32, bool, error) {
	var comicNfts []entity.ComicNft

	rnft := query.Use(c.data.DB).ComicsNft
	query := rnft.WithContext(ctx).Where(rnft.ComicsID.Eq(comicId)).Where(rnft.Owner.EqCol(rnft.Author))
	datas, err := query.Offset(int(lastScore)).
		Limit(int(limit)).Preload(rnft.Comic).Find()
	totalSupply, err := query.Count()

	if err != nil {
		return nil, 0, 0, false, err
	}
	var infos []struct {
		ContractAddress string
		TokenId         string
	}
	worksComic := make(map[int]*model.Comic)

	for index, d := range datas {
		infos = append(infos, struct {
			ContractAddress string
			TokenId         string
		}{ContractAddress: d.Comic.ContractAddress, TokenId: d.Comic.TokenID})
		worksComic[index] = &d.Comic
	}

	origins, err := c.nftgo.BatchGetNftSummary(ctx, infos)
	originNfts := make(map[string]entity.Nft)

	for _, item := range origins {
		originNfts[item.TokenId] = item
	}

	for tokenId, item := range datas {
		id := strconv.FormatInt(int64(item.ID), 10)
		comicNfts = append(comicNfts, entity.ComicNft{
			TokenId: id,
			Owner:   item.Owner,
			Comic: entity.Comic{
				Origin:      originNfts[worksComic[tokenId].TokenID],
				ComicId:     fmt.Sprint(worksComic[tokenId].ID),
				MintLimit:   uint32(item.Comic.MintLimit),
				MintPrice:   item.Comic.MintPrice,
				Name:        item.Comic.Name,
				CreatedAt:   item.Comic.CreatedAt,
				UserAddress: item.Comic.UserAddress,
				ImageUris:   strings.Split(item.Comic.ImageURIs, ","),
			},
		})
	}
	nextLastScore := int64(limit) + lastScore
	return comicNfts, nextLastScore, uint32(totalSupply), nextLastScore < totalSupply, nil
}

func NewComic(data *data.Data, nftgo *nftgo.Service) *Comic {
	return &Comic{data: data, nftgo: nftgo}
}
