package repo

import (
	"4clabs-server/app/4clabs-server/internal/adapter/data/model"
	"4clabs-server/app/4clabs-server/internal/adapter/data/query"
	"4clabs-server/app/4clabs-server/internal/adapter/driving/nftgo"
	"4clabs-server/app/4clabs-server/internal/data"
	"4clabs-server/app/4clabs-server/internal/domain/entity"
	"4clabs-server/app/4clabs-server/internal/ports"
	"context"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

type Register struct {
	nftgo *nftgo.Service
	data  *data.Data
}

func NewRegister(nftgo *nftgo.Service, data *data.Data) *Register {
	return &Register{nftgo: nftgo, data: data}
}

func (r Register) ListRegistedNfts(ctx context.Context, userAddress string, limit uint32, nextScore int64) ([]entity.Nft, int64, uint32, bool, error) {
	if nextScore == 0 {
		nextScore = math.MaxInt64
	}
	rnft := query.Use(r.data.DB).RegisterNft

	var datas []*model.RegisterNft
	var err error
	if userAddress != "" {
		datas, err = rnft.WithContext(ctx).
			Where(rnft.UserAddress.Eq(userAddress)).
			Where(rnft.CreatedAt.Lt(time.Unix(nextScore, 0))).
			Order(rnft.CreatedAt.Desc()).Limit(int(limit + 1)).Find()
	} else {
		datas, err = rnft.WithContext(ctx).
			Where(rnft.CreatedAt.Lt(time.Unix(nextScore, 0))).
			Order(rnft.CreatedAt.Desc()).Limit(int(limit + 1)).Find()
	}

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

	nfts, err := r.nftgo.BatchGetNftSummary(ctx, infos)
	if err != nil {
		return nil, 0, 0, false, err
	}

	var result []entity.Nft
	for _, d := range datas {
		for _, nft := range nfts {
			if d.ContractAddress == nft.ContractAddress && d.TokenID == nft.TokenId {
				result = append(result, nft)
			}
		}
	}

	return result, lastScore, uint32(total), hasMore, nil
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
		if !ok {
			needCreated = append(needCreated, &model.RegisterNft{
				TokenID:         n.TokenId,
				ContractAddress: n.ContractAddress,
				UserAddress:     userAddress,
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
