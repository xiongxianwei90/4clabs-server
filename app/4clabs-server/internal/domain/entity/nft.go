package entity

import "time"

type Rarity struct {
	Score float64 `json:"score"`
	Rank  int     `json:"rank"`
	Total int     `json:"total"`
}
type Traits struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type OwnerStat struct {
	Address     string
	HoldingTime time.Duration
}
type PriceInfo struct {
	TxHash               string
	PriceToken           float64
	TokenSymbol          string
	TokenContractAddress string
	Time                 time.Time
}

type NftStat struct {
	LastUpdated        uint64
	SaleNum7d          uint64
	SaleNumAll         uint64
	MaxPrice           PriceInfo
	MinPrice           PriceInfo
	PastOwners         []OwnerStat
	CreateTime         time.Time
	StartHoldingTime   time.Time
	LongestHoldingTime time.Duration
}

type NftDetail struct {
	Summary Nft
	Stat    NftStat
}
type Nft struct {
	Blockchain      string
	CollectionName  string
	CollectionSlug  string
	ContractAddress string
	TokenId         string
	Name            string
	Description     string
	Image           string
	AnimationUrl    string
	OwnerAddresses  []string
	Traits          []Traits
	Rarity          Rarity
}
