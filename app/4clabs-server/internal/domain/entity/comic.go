package entity

import "time"

type Comic struct {
	Origin       Nft
	MintLimit    uint32
	MintPrice    float64 // eth
	Name         string
	MetadataJson string
	CreatedAt    time.Time
	UserAddress  string
}
