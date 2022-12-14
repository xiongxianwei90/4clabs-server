// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameNft = "nfts"

// Nft mapped from table <nfts>
type Nft struct {
	ID                     int32     `gorm:"column:id;type:int(11) unsigned;primaryKey;autoIncrement:true" json:"id"`
	ContractAddressTokenID string    `gorm:"column:contract_address_token_id;type:varchar(200);not null" json:"contract_address_token_id"`
	Image                  string    `gorm:"column:image;type:varchar(200);not null" json:"image"`
	CreatedAt              time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt              time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName Nft's table name
func (*Nft) TableName() string {
	return TableNameNft
}
