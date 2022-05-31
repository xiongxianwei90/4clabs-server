package service

import (
	apibase "4clabs-server/api/base/v1"
	nft "4clabs-server/api/nft/v1"
	v1 "4clabs-server/api/service/v1"
	"4clabs-server/app/4clabs-server/internal/adapter/assembler"
	"4clabs-server/app/4clabs-server/internal/usecase"
	"context"
)

type Service struct {
	v1.UnimplementedNftServer
	uc *usecase.Address
}

func NewService(uc *usecase.Address) *Service {
	return &Service{uc: uc}
}

func (s *Service) GetAddressNfts(ctx context.Context, req *nft.GetAddressNftsRequest) (*nft.GetAddressNftResponse, error) {
	nfts, nextScore, total, hasMore, err := s.uc.GetAddressNfts(ctx, req.Address, req.BaseListRequest.Limit, req.BaseListRequest.LastScore)
	if err != nil {
		return nil, err
	}
	res := &nft.GetAddressNftResponse{
		BaseListResponse: &apibase.BaseListResponse{
			LastScore: nextScore,
			HasMore:   hasMore,
			Total:     total,
		},
		Summaries: assembler.CoverNftToHttpDto(nfts...),
	}
	return res, nil
}
