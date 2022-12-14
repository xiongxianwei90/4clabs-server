package nftgo

import (
	"4clabs-server/app/4clabs-server/internal/adapter/data/model"
	"4clabs-server/app/4clabs-server/internal/adapter/data/query"
	"4clabs-server/app/4clabs-server/internal/adapter/driving/repo/nft"
	"4clabs-server/app/4clabs-server/internal/conf"
	"4clabs-server/app/4clabs-server/internal/data"
	"4clabs-server/app/4clabs-server/internal/domain/entity"
	"4clabs-server/app/4clabs-server/internal/ports"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"time"
)

var _ ports.Query = &Service{}

type Service struct {
	apiKey          string
	cacheNfts       *nft.Nft
	ContractAddress string
	Rawurl          string
	data            *data.Data
}

func NewService(bootstrap *conf.Bootstrap, cacheNft *nft.Nft, data *data.Data) *Service {
	return &Service{apiKey: bootstrap.ThirdParty.Nftgo.ApiKey, cacheNfts: cacheNft, ContractAddress: bootstrap.ThirdParty.Contract.Address, Rawurl: bootstrap.ThirdParty.Contract.Rawurl, data: data}
}

type Nft struct {
	Blockchain      string   `json:"blockchain"`
	CollectionName  string   `json:"collection_name"`
	CollectionSlug  string   `json:"collection_slug"`
	ContractAddress string   `json:"contract_address"`
	TokenId         string   `json:"token_id"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Image           string   `json:"image"`
	AnimationUrl    string   `json:"animation_url"`
	OwnerAddresses  []string `json:"owner_addresses"`
	Traits          []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"traits"`
	Rarity struct {
		Score float64 `json:"score"`
		Rank  int     `json:"rank"`
		Total int     `json:"total"`
	} `json:"rarity"`
}

type AddressPortfolioResponse struct {
	Total  int `json:"total"`
	Assets []struct {
		Nft      Nft `json:"nft"`
		Quantity int `json:"quantity"`
	} `json:"assets"`
}
type NftMetrics struct {
	LastUpdated int `json:"last_updated"`
	SaleNum7D   int `json:"sale_num_7d"`
	SaleNumAll  int `json:"sale_num_all"`
	MaxPrice    struct {
		TxHash               string  `json:"tx_hash"`
		PriceToken           float64 `json:"price_token"`
		TokenSymbol          string  `json:"token_symbol"`
		TokenContractAddress string  `json:"token_contract_address"`
		PriceUsd             float64 `json:"price_usd"`
		Time                 int     `json:"time"`
	} `json:"max_price"`
	MinPrice struct {
		TxHash               string  `json:"tx_hash"`
		PriceToken           float64 `json:"price_token"`
		TokenSymbol          string  `json:"token_symbol"`
		TokenContractAddress string  `json:"token_contract_address"`
		PriceUsd             float64 `json:"price_usd"`
		Time                 int     `json:"time"`
	} `json:"min_price"`
	LastPrice struct {
		TxHash               string  `json:"tx_hash"`
		PriceToken           float64 `json:"price_token"`
		TokenSymbol          string  `json:"token_symbol"`
		TokenContractAddress string  `json:"token_contract_address"`
		PriceUsd             float64 `json:"price_usd"`
		Time                 int     `json:"time"`
	} `json:"last_price"`
	PastOwners []struct {
		Address     string `json:"address"`
		HoldingTime int    `json:"holding_time"`
	} `json:"past_owners"`
	CreateTime         int `json:"create_time"`
	StartHoldingTime   int `json:"start_holding_time"`
	LongestHoldingTime int `json:"longest_holding_time"`
}

func (s *Service) GetNftDetail(ctx context.Context, contractAddress, tokenId string) (entity.NftDetail, error) {
	var eg errgroup.Group
	ssChan := make(chan entity.Nft, 2)
	metricsChan := make(chan entity.NftStat, 2)
	eg.Go(func() error {
		ss, err := s.GetNftSummary(ctx, contractAddress, tokenId)
		if err != nil {
			return err
		}
		ssChan <- ss
		return nil
	})
	eg.Go(func() error {
		urlStr := fmt.Sprintf("https://api.nftgo.dev/eth/v1/nft/%s/%s/metrics", contractAddress, tokenId)
		reader, err := s.baseGet(ctx, urlStr)
		if err != nil {
			return err
		}
		defer reader.Close()
		metrics := NftMetrics{}
		if err := json.NewDecoder(reader).Decode(&metrics); err != nil {
			return errors.Wrapf(err, "address : %s, id : %s", contractAddress, tokenId)
		}
		metricsChan <- s.coverMetrics(metrics)
		return nil
	})
	if err := eg.Wait(); err != nil {
		return entity.NftDetail{}, err
	}
	close(metricsChan)
	close(ssChan)
	return entity.NftDetail{
		Summary: <-ssChan,
		Stat:    <-metricsChan,
	}, nil
}

func (s *Service) BatchGetNftSummary(ctx context.Context, infos []struct {
	ContractAddress string
	TokenId         string
}) ([]entity.Nft, error) {
	nfts := make(chan entity.Nft, len(infos))
	var eg errgroup.Group
	for _, i := range infos {
		i := i
		eg.Go(func() error {
			nft, _ := s.GetNftSummary(ctx, i.ContractAddress, i.TokenId)
			// nftgo 429 ????????????
			//if err != nil {
			//	return err
			//}
			nfts <- nft
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}
	close(nfts)
	var result []entity.Nft
	var ids []string
	for n := range nfts {
		result = append(result, n)
		ids = append(ids, fmt.Sprintf("%s_%s", n.ContractAddress, n.TokenId))
	}
	images, err := s.cacheNfts.Image(ctx, ids...)
	if err != nil {
		return nil, err
	}
	for idx, n := range result {
		n := n
		if im, ok := images[fmt.Sprintf("%s_%s", n.ContractAddress, n.TokenId)]; ok {
			n.Image = im
		}
		result[idx] = n
	}
	return result, nil
}

func (s *Service) GetNftSummary(ctx context.Context, contractAddress, tokenId string) (entity.Nft, error) {
	urlStr := fmt.Sprintf("https://api.nftgo.dev/eth/v1/nft/%s/%s/info", contractAddress, tokenId)
	reader, err := s.baseGet(ctx, urlStr)

	if err != nil {
		return entity.Nft{}, err
	}
	defer reader.Close()
	nft := Nft{}
	if err := json.NewDecoder(reader).Decode(&nft); err != nil {
		return entity.Nft{}, errors.Wrapf(err, "address : %s, token : %s ", contractAddress, tokenId)
	}
	return *s.coverNftToEntity(nft)[0], nil
}

func (s *Service) GetAddressNfts(ctx context.Context, address string, limit uint32, lastScore int64) ([]*entity.Nft, int64, uint32, bool, error) {
	urlStr := fmt.Sprintf("https://api.nftgo.dev/eth/v1/address/%s/portfolio?offset=%d&limit=%d", address, lastScore, limit)
	reader, err := s.baseGet(ctx, urlStr)
	if reader == nil {
		return nil, 0, 0, false, nil
	}

	defer reader.Close()
	if err != nil {
		return nil, 0, 0, false, err
	}
	addressPortfolio := AddressPortfolioResponse{}
	if err := json.NewDecoder(reader).Decode(&addressPortfolio); err != nil {
		return nil, 0, 0, false, errors.Wrapf(err, "address : %s", address)
	}

	var result []*entity.Nft
	registerMap := make(map[string]bool)
	var tokenIds []string
	for _, a := range addressPortfolio.Assets {
		result = append(result, s.coverNftToEntity(a.Nft)...)
		tokenIds = append(tokenIds, a.Nft.TokenId)
	}
	rnft := query.Use(s.data.DB).RegisterNft

	var registerData []*model.RegisterNft
	registerData, err = rnft.WithContext(ctx).Where(rnft.TokenID.In(tokenIds...)).Find()
	if registerData != nil {
		for _, data := range registerData {
			registerMap[data.TokenID] = true
		}
		for _, nftItem := range result {
			nftItem.Registered = registerMap[nftItem.TokenId]
		}
	}
	return result, lastScore + int64(limit), uint32(addressPortfolio.Total), lastScore+int64(limit) < int64(addressPortfolio.Total), nil
}

func (s *Service) coverMetrics(detail NftMetrics) entity.NftStat {
	var PastOwners []entity.OwnerStat
	for _, w := range detail.PastOwners {
		PastOwners = append(PastOwners, entity.OwnerStat{
			Address:     w.Address,
			HoldingTime: time.Duration(int64(w.HoldingTime)),
		})
	}
	return entity.NftStat{
		LastUpdated: uint64(detail.LastUpdated),
		SaleNum7d:   uint64(detail.SaleNum7D),
		SaleNumAll:  uint64(detail.SaleNumAll),
		MaxPrice: entity.PriceInfo{
			TxHash:               detail.MaxPrice.TxHash,
			PriceToken:           detail.MaxPrice.PriceToken,
			TokenSymbol:          detail.MaxPrice.TokenSymbol,
			TokenContractAddress: detail.MaxPrice.TokenContractAddress,
			PriceUsd:             detail.MaxPrice.PriceUsd,
			Time:                 time.Unix(int64(detail.MaxPrice.Time), 0),
		},
		MinPrice: entity.PriceInfo{
			TxHash:               detail.MinPrice.TxHash,
			PriceToken:           detail.MinPrice.PriceToken,
			TokenSymbol:          detail.MinPrice.TokenSymbol,
			TokenContractAddress: detail.MinPrice.TokenContractAddress,
			PriceUsd:             detail.MinPrice.PriceUsd,
			Time:                 time.Unix(int64(detail.MinPrice.Time), 0),
		},
		LastPrice: entity.PriceInfo{
			TxHash:               detail.LastPrice.TxHash,
			PriceToken:           detail.LastPrice.PriceToken,
			TokenSymbol:          detail.LastPrice.TokenSymbol,
			TokenContractAddress: detail.LastPrice.TokenContractAddress,
			PriceUsd:             detail.LastPrice.PriceUsd,
			Time:                 time.Unix(int64(detail.LastPrice.Time), 0),
		},
		PastOwners:         PastOwners,
		CreateTime:         time.Unix(int64(detail.CreateTime), 0),
		StartHoldingTime:   time.Unix(int64(detail.StartHoldingTime), 0),
		LongestHoldingTime: time.Duration(detail.LongestHoldingTime),
	}
}
func (s *Service) coverNftToEntity(nfts ...Nft) []*entity.Nft {
	var result []*entity.Nft
	for _, nft := range nfts {
		var ts []entity.Traits
		for _, t := range nft.Traits {
			ts = append(ts, entity.Traits{
				Type:  t.Type,
				Value: t.Value,
			})
		}
		result = append(result, &entity.Nft{
			Blockchain:      nft.Blockchain,
			CollectionName:  nft.CollectionName,
			CollectionSlug:  nft.CollectionSlug,
			ContractAddress: nft.ContractAddress,
			TokenId:         nft.TokenId,
			Name:            nft.Name,
			Description:     nft.Description,
			Image:           nft.Image,
			AnimationUrl:    nft.AnimationUrl,
			OwnerAddresses:  nft.OwnerAddresses,
			Traits:          nil,
			Rarity: entity.Rarity{
				Score: nft.Rarity.Score,
				Rank:  nft.Rarity.Rank,
				Total: nft.Rarity.Total,
			},
		})
	}
	return result
}
func (s *Service) baseGet(ctx context.Context, url string) (io.ReadCloser, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	req.Header.Set("X-API-KEY", s.apiKey)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "url : %s", url)
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(fmt.Errorf("status code %d is not ok", res.StatusCode), "url : %s", url)
	}
	return res.Body, nil
}
