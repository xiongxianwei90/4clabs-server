package nftgo

import (
	"4clabs-server/internal/domain/entity"
	"4clabs-server/internal/domain/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

type T struct {
	LastUpdated        int     `json:"last_updated"`
	HoldingValueUsd    float64 `json:"holding_value_usd"`
	EstHoldingValueUsd float64 `json:"est_holding_value_usd"`
	PnlUsd             float64 `json:"pnl_usd"`
	BuyVolumeUsd       float64 `json:"buy_volume_usd"`
	SellVolumeUsd      float64 `json:"sell_volume_usd"`
	ActivityNum        int     `json:"activity_num"`
	MintNum            struct {
		H  int `json:"24h"`
		D  int `json:"7d"`
		D1 int `json:"30d"`
		M  int `json:"3M"`
	} `json:"mint_num"`
	BuyNum struct {
		H  int `json:"24h"`
		D  int `json:"7d"`
		D1 int `json:"30d"`
		M  int `json:"3M"`
	} `json:"buy_num"`
	SellNum struct {
		H  int `json:"24h"`
		D  int `json:"7d"`
		D1 int `json:"30d"`
		M  int `json:"3M"`
	} `json:"sell_num"`
	BurnNum struct {
		H  int `json:"24h"`
		D  int `json:"7d"`
		D1 int `json:"30d"`
		M  int `json:"3M"`
	} `json:"burn_num"`
	SendNum struct {
		H  int `json:"24h"`
		D  int `json:"7d"`
		D1 int `json:"30d"`
		M  int `json:"3M"`
	} `json:"send_num"`
	ReceiveNum struct {
		H  int `json:"24h"`
		D  int `json:"7d"`
		D1 int `json:"30d"`
		M  int `json:"3M"`
	} `json:"receive_num"`
	NftNum        int `json:"nft_num"`
	CollectionNum int `json:"collection_num"`
}

var _ service.Query = &Service{}

type Service struct {
	apiKey string
}
type AddressPortfolioResponse struct {
	Total  int `json:"total"`
	Assets []struct {
		Nft struct {
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
		} `json:"nft"`
		Quantity int `json:"quantity"`
	} `json:"assets"`
}

func (s *Service) GetAddressNfts(ctx context.Context, address string, limit uint32, lastScore int64) ([]entity.Nft, int64, uint32, bool, error) {
	urlStr := fmt.Sprintf("https://api.nftgo.dev/eth/v1/address/%s/portfolio?offset=%d&limit=%d", address, lastScore, limit)
	reader, err := s.baseGet(ctx, urlStr)
	defer reader.Close()
	if err != nil {
		return nil, 0, 0, false, err
	}
	addressPortfolio := AddressPortfolioResponse{}
	if err := json.NewDecoder(reader).Decode(&addressPortfolio); err != nil {
		return nil, 0, 0, false, errors.Wrapf(err, "address : %s", address)
	}

	var result []entity.Nft
	for _, a := range addressPortfolio.Assets {
		var ts []entity.Traits
		for _, t := range a.Nft.Traits {
			ts = append(ts, entity.Traits{
				Type:  t.Type,
				Value: t.Value,
			})
		}
		result = append(result, entity.Nft{
			Blockchain:      a.Nft.Blockchain,
			CollectionName:  a.Nft.CollectionName,
			CollectionSlug:  a.Nft.CollectionSlug,
			ContractAddress: a.Nft.ContractAddress,
			TokenId:         a.Nft.TokenId,
			Name:            a.Nft.Name,
			Description:     a.Nft.Description,
			Image:           a.Nft.Image,
			AnimationUrl:    a.Nft.AnimationUrl,
			OwnerAddresses:  a.Nft.OwnerAddresses,
			Traits:          nil,
			Rarity: entity.Rarity{
				Score: a.Nft.Rarity.Score,
				Rank:  a.Nft.Rarity.Rank,
				Total: a.Nft.Rarity.Total,
			},
		})
	}
	return result, lastScore + int64(limit), uint32(addressPortfolio.Total), false, nil
}

func (s *Service) baseGet(ctx context.Context, url string) (io.ReadCloser, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	req.Header.Set("X-API-KEY", s.apiKey)
	res, err := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		return nil, errors.Wrapf(err, "url : %s", url)
	}
	return res.Body, nil
}
