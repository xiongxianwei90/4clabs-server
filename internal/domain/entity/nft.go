package entity

type Rarity struct {
	Score float64 `json:"score"`
	Rank  int     `json:"rank"`
	Total int     `json:"total"`
}
type Traits struct {
	Type  string `json:"type"`
	Value string `json:"value"`
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
