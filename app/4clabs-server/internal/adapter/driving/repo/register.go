package repo

import (
	"4clabs-server/api/nft/v1"
	"4clabs-server/app/4clabs-server/internal/adapter/data/model"
	"4clabs-server/app/4clabs-server/internal/adapter/data/query"
	"4clabs-server/app/4clabs-server/internal/adapter/driving/nftgo"
	"4clabs-server/app/4clabs-server/internal/adapter/driving/repo/nft"
	"4clabs-server/app/4clabs-server/internal/data"
	"4clabs-server/app/4clabs-server/internal/domain/entity"
	"4clabs-server/app/4clabs-server/internal/ports"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Register struct {
	nftgo     *nftgo.Service
	data      *data.Data
	cacheNfts *nft.Nft
}

func NewRegister(nftgo *nftgo.Service, data *data.Data, cacheNfts *nft.Nft) *Register {
	return &Register{nftgo: nftgo, data: data, cacheNfts: cacheNfts}
}

func (r Register) ListRegistedNfts(ctx context.Context, userAddress string, sort *v1.Sort, limit uint32, nextScore int64) ([]*entity.Nft, int64, uint32, bool, error) {

	rnft := query.Use(r.data.DB).RegisterNft

	var datas []*model.RegisterNft
	query := rnft.WithContext(ctx)

	var err error
	if userAddress != "" {
		query = query.
			Where(rnft.UserAddress.Eq(userAddress))
	}

	total, err := query.Count()
	if err != nil {
		return nil, 0, 0, false, err
	}
	query = query.Offset(int(nextScore)).Limit(int(limit))

	if sort != nil {
		if sort.ByPrice == 2 {
			query = query.Order(rnft.Price.Desc())
		} else {
			query = query.Order(rnft.Price)
		}
	} else {
		query.Order(rnft.CreatedAt.Desc())
	}

	datas, err = query.Find()

	if err != nil {
		return nil, 0, 0, false, err
	}

	if len(datas) == 0 {
		return nil, 0, 0, false, nil
	}

	var infos []struct {
		ContractAddress string
		TokenId         string
	}
	var ids []string
	for _, d := range datas {
		infos = append(infos, struct {
			ContractAddress string
			TokenId         string
		}{ContractAddress: d.ContractAddress, TokenId: d.TokenID})

		ids = append(ids, fmt.Sprintf("%s_%s", d.ContractAddress, d.TokenID))
	}

	images, err := r.cacheNfts.Image(ctx, ids...)

	// 每次从nftgo 获取会有频次限制。数据存入数据库
	var result []*entity.Nft
	for _, d := range datas {
		image := d.Image
		if im, ok := images[fmt.Sprintf("%s_%s", d.ContractAddress, d.TokenID)]; ok {
			image = im
		}
		result = append(result, &entity.Nft{
			CollectionName:  d.CollectionName,
			ContractAddress: d.ContractAddress,
			TokenId:         d.TokenID,
			Name:            d.Name,
			OwnerAddresses:  []string{d.UserAddress},
			Image:           image,
		})
	}

	return result, nextScore + int64(limit), uint32(total), nextScore+int64(limit) < total, nil
}

func (r Register) UserRegistered(ctx context.Context, userAddress string) (bool, error) {
	rnft := query.Use(r.data.DB).RegisterNft
	if _, err := rnft.WithContext(ctx).Where(rnft.UserAddress.Eq(userAddress)).First(); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, errors.Wrapf(err, "address : %s", userAddress)
	}
	return true, nil
}

func (r Register) Register(ctx context.Context, nfts []entity.BaseNft, userAddress string) error {
	var needCreated []*model.RegisterNft
	for _, n := range nfts {
		ok, err := r.Registered(ctx, n.TokenId, n.ContractAddress, userAddress)
		if err != nil {
			return err
		}

		detail, _ := r.nftgo.GetNftDetail(ctx, n.ContractAddress, n.TokenId)

		if !ok {
			needCreated = append(needCreated, &model.RegisterNft{
				TokenID:         n.TokenId,
				Name:            detail.Summary.Name,
				CollectionName:  detail.Summary.CollectionName,
				ContractAddress: n.ContractAddress,
				UserAddress:     userAddress,
				Price:           detail.Stat.LastPrice.PriceToken,
				Image:           detail.Summary.Image,
			})
		}
	}
	rNft := query.Use(r.data.DB).RegisterNft
	if err := rNft.WithContext(ctx).Create(needCreated...); err != nil {
		return errors.Wrapf(err, "model : %#v, address : %s", nfts, userAddress)
	}
	return nil
}

func (r Register) Registered(ctx context.Context, nftTokenId string, nftContractAddress string, userAddress string) (bool, error) {
	rNft := query.Use(r.data.DB).RegisterNft
	if _, err := rNft.WithContext(ctx).Where(rNft.ContractAddress.Eq(nftContractAddress), rNft.TokenID.Eq(nftTokenId)).First(); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, errors.Wrapf(err, "token : %s, contract : %s, address : %s", nftTokenId, nftContractAddress, userAddress)
	}
	return true, nil
}

var _ ports.Register = &Register{}
