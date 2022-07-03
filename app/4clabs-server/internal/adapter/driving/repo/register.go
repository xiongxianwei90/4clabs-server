package repo

import (
	"4clabs-server/app/4clabs-server/internal/adapter/data/model"
	"4clabs-server/app/4clabs-server/internal/adapter/data/query"
	"4clabs-server/app/4clabs-server/internal/data"
	"4clabs-server/app/4clabs-server/internal/ports"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Register struct {
	data *data.Data
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

func (r Register) Register(ctx context.Context, nftTokenId string, nftContractAddress string, userAddress string) error {
	ok, err := r.Registered(ctx, nftTokenId, nftContractAddress, userAddress)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	rNft := query.Use(r.data.DB).RegisterNft
	if err := rNft.WithContext(ctx).Create(&model.RegisterNft{
		TokenID:         nftTokenId,
		ContractAddress: nftContractAddress,
		UserAddress:     userAddress,
	}); err != nil {
		return errors.Wrapf(err, "token : %s, contract : %s, address : %s", nftTokenId, nftContractAddress, userAddress)
	}
	return nil
}

func (r Register) Registered(ctx context.Context, nftTokenId string, nftContractAddress string, userAddress string) (bool, error) {
	rNft := query.Use(r.data.DB).RegisterNft
	if _, err := rNft.WithContext(ctx).First(); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, errors.Wrapf(err, "token : %s, contract : %s, address : %s", nftTokenId, nftContractAddress, userAddress)
	}
	return true, nil
}

var _ ports.Register = &Register{}

func NewRegister(data *data.Data) *Register {
	return &Register{data: data}
}
