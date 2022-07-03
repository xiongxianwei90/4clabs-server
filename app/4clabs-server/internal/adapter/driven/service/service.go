package service

import (
	auth "4clabs-server/api/auth/v1"
	apibase "4clabs-server/api/base/v1"
	nft "4clabs-server/api/nft/v1"
	v1 "4clabs-server/api/service/v1"
	ticketPb "4clabs-server/api/tickets/v1"
	"4clabs-server/app/4clabs-server/internal/adapter/assembler"
	"4clabs-server/app/4clabs-server/internal/usecase"
	"context"
)

type Service struct {
	v1.UnimplementedNftServer
	addressUc *usecase.Address
	nftUc     *usecase.Nft
	auth      *usecase.Auth
	ticket    *usecase.Ticket
}

func NewService(addressUc *usecase.Address, nftUc *usecase.Nft, auth *usecase.Auth, ticket *usecase.Ticket) *Service {
	return &Service{addressUc: addressUc, nftUc: nftUc, auth: auth, ticket: ticket}
}

// ticket WL
func (s *Service) InTicketsWLRequest(ctx context.Context, req *ticketPb.CanMintRequest) (*ticketPb.CantMintResponse, error) {
	res := &ticketPb.CantMintResponse{}
	ok, level, err := s.ticket.Query(ctx, req.Address)
	if err != nil {
		return nil, err
	}
	res.Level = level
	res.InWl = ok
	return res, nil
}

func (s *Service) GetNftDetail(ctx context.Context, req *nft.GetNftDetailRequest) (*nft.GetNftDetailResponse, error) {
	detail, err := s.nftUc.GetNftDetail(ctx, req.ContractAddress, req.TokenId)
	if err != nil {
		return nil, err
	}
	res := &nft.GetNftDetailResponse{
		Detail: assembler.CoverNftDetailToHttpDto(detail),
	}
	return res, nil
}
func (s *Service) SignToLogin(ctx context.Context, req *auth.VerifySignToLoginSignRequest) (*auth.VerifySignToLoginSighResponse, error) {
	res := &auth.VerifySignToLoginSighResponse{}
	var err error
	if res.Token, err = s.auth.SignToLogin(ctx, req.Address, req.Sign, req.Message); err != nil {
		return nil, err
	}
	return res, nil
}
func (s *Service) FetchNonce(ctx context.Context, req *auth.FetchSignMessageRequest) (*auth.FetchSignMessageResponse, error) {
	res := &auth.FetchSignMessageResponse{}
	var err error
	if res.SignMessage, err = s.auth.GetSignMessage(ctx, req.Address); err != nil {
		return nil, err
	}
	return res, nil
}
func (s *Service) GetAddressNfts(ctx context.Context, req *nft.GetAddressNftsRequest) (*nft.GetAddressNftResponse, error) {
	nfts, nextScore, total, hasMore, err := s.addressUc.GetAddressNfts(ctx, req.Address, req.BaseListRequest.Limit, req.BaseListRequest.LastScore)
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
