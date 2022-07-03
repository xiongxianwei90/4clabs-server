package service

import (
	auth "4clabs-server/api/auth/v1"
	apibase "4clabs-server/api/base/v1"
	nft "4clabs-server/api/nft/v1"
	v1 "4clabs-server/api/service/v1"
	ticketPb "4clabs-server/api/tickets/v1"
	"4clabs-server/app/4clabs-server/internal/adapter/assembler"
	"4clabs-server/app/4clabs-server/internal/usecase"
	authUtils "4clabs-server/pkg/auth"
	"context"
)

type Service struct {
	v1.UnimplementedNftServer
	addressUc *usecase.Address
	nftUc     *usecase.Nft
	auth      *usecase.Auth
	authUtils *authUtils.ContextUtils
	ticket    *usecase.Ticket
	comic     *usecase.Comics
}

func NewService(addressUc *usecase.Address, nftUc *usecase.Nft, auth *usecase.Auth, authUtils *authUtils.ContextUtils, ticket *usecase.Ticket, comic *usecase.Comics) *Service {
	return &Service{addressUc: addressUc, nftUc: nftUc, auth: auth, authUtils: authUtils, ticket: ticket, comic: comic}
}

func (s *Service) CreateComic(ctx context.Context, req *nft.ComicCreateRequest) (*nft.ComicCreateResponse, error) {
	comics, nextScore, total, hasMore, err := s.comic.List(ctx, s.authUtils.GetAddress(ctx),
}

// commic works
func (s *Service) ListComicWorks(ctx context.Context, req *nft.ListComicWorkRequest) (*nft.ListComicWorkResponse, error) {
	comics, nextScore, total, hasMore, err := s.comic.List(ctx, req.Address, req.BaseListRequest.Limit, req.BaseListRequest.LastScore)
	if err != nil {
		return nil, err
	}
	return &nft.ListComicWorkResponse{
		BaseListResponse: &apibase.BaseListResponse{
			LastScore: nextScore,
			HasMore:   hasMore,
			Total:     total,
		},
		ComicWorks: assembler.CoverComicToHttpDto(comics...),
	}, nil
}

func (s *Service) RegisterNft(ctx context.Context, req *nft.RegisterNftRequest) (*nft.RegisterNftResponse, error) {
	if err := s.nftUc.Register(ctx, req.TokenId, req.ContractAddress, req.UserAddress); err != nil {
		return nil, err
	}
	return &nft.RegisterNftResponse{}, nil
}

// register nfts
func (s *Service) ListRegsiterNfts(ctx context.Context, req *nft.ListRegisterNftRequest) (*nft.ListRegisterNftResponse, error) {
	nfts, nextScore, total, hasMore, err := s.nftUc.ListRegistedNfts(ctx, req.Address, req.BaseListRequest.Limit, req.BaseListRequest.LastScore)
	if err != nil {
		return nil, err
	}
	return &nft.ListRegisterNftResponse{
		BaseListResponse: &apibase.BaseListResponse{
			LastScore: nextScore,
			HasMore:   hasMore,
			Total:     total,
		},
		Summaries: assembler.CoverNftToHttpDto(nfts...),
	}, nil
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
	if res.Token, res.Registered, err = s.auth.SignToLogin(ctx, req.Address, req.Sign, req.Message); err != nil {
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
