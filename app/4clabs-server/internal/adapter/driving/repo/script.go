package repo

import (
	"4clabs-server/app/4clabs-server/internal/adapter/driving/nftgo"
	"4clabs-server/app/4clabs-server/internal/data"
	"4clabs-server/app/4clabs-server/internal/domain/entity"
	"4clabs-server/app/4clabs-server/internal/ports"
	"context"
)

var _ ports.Script = (*Script)(nil)

type Script struct {
	data     *data.Data
	nftgo    *nftgo.Service
	register *Register
}

func NewScript(data *data.Data, nftgo *nftgo.Service, register *Register) *Script {
	return &Script{data: data, nftgo: nftgo, register: register}
}

func (s *Script) RegisterUpdate(ctx context.Context, contactAddress string, tokenId string, userAddress string) error {
	registered, err := s.register.Registered(ctx, tokenId, contactAddress, userAddress)
	if err != nil {
		return err
	}
	if registered {
		return nil
	}
	var nfts []entity.BaseNft
	nfts = append(nfts, entity.BaseNft{
		ContractAddress: contactAddress,
		TokenId:         tokenId,
	})

	err = s.register.Register(ctx, nfts, userAddress)

	return err
}
