package repo

import (
	"4clabs-server/app/4clabs-server/internal/adapter/data/query"
	"4clabs-server/app/4clabs-server/internal/adapter/driving/nftgo"
	"4clabs-server/app/4clabs-server/internal/conf"
	"4clabs-server/app/4clabs-server/internal/data"
	"4clabs-server/app/4clabs-server/internal/domain/entity"
	"4clabs-server/app/4clabs-server/internal/ports"
	"context"
)

var _ ports.Script = (*Script)(nil)

type Script struct {
	data            *data.Data
	nftgo           *nftgo.Service
	register        *Register
	ContractAddress string
	Rawurl          string
}

func NewScript(bootstrap *conf.Bootstrap, data *data.Data, nftgo *nftgo.Service, register *Register) *Script {
	return &Script{
		data:            data,
		nftgo:           nftgo,
		register:        register,
		ContractAddress: bootstrap.ThirdParty.Contract.Address,
		Rawurl:          bootstrap.ThirdParty.Contract.Rawurl,
	}
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

func (s *Script) ComicWorksUpdate(ctx context.Context, isIncrement bool) error {
	err := nftgo.ComicUpdate(ctx, s.data.DB, s.Rawurl, s.ContractAddress, isIncrement)
	err = nftgo.ComicNftUpdate(ctx, s.data.DB, s.Rawurl, s.ContractAddress, isIncrement)
	return err
}

func (s *Script) ComicWorksSold(ctx context.Context, to string, tokenId int64) error {
	comicsNfts := query.Use(s.data.DB).ComicsNft
	_, err := comicsNfts.WithContext(ctx).
		Where(comicsNfts.ID.Eq(int32(tokenId))).
		Update(comicsNfts.Owner, to)
	return err
}
