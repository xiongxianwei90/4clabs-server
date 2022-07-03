package entity

import "time"

type Comic struct {
	Origin      Nft
	MintLimit   uint32
	MintPrice   float64 // eth
	Name        string
	ImageUrls   []string
	CreatedAt   time.Time
	UserAddress string
}
