// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameComic = "comics"

// Comic mapped from table <comics>
type Comic struct {
	ID              int32     `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id"`
	TokenID         string    `gorm:"column:token_id;type:varchar(200);not null;index:idx_contract_address_token_id,priority:2" json:"token_id"`
	Name            string    `gorm:"column:name;type:varchar(200);not null" json:"name"`
	ContractAddress string    `gorm:"column:contract_address;type:varchar(200);not null;index:idx_contract_address_token_id,priority:1" json:"contract_address"`
	UserAddress     string    `gorm:"column:user_address;type:varchar(200);not null;index:idx_user_address,priority:1" json:"user_address"`
	MintLimit       int32     `gorm:"column:mint_limit;type:int(11);not null" json:"mint_limit"`
	MintPrice       float64   `gorm:"column:mint_price;type:double;not null" json:"mint_price"`
	ImageURIs       string    `gorm:"column:imageURIs;type:varchar(1000)" json:"imageURIs"`
	CreatedAt       time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName Comic's table name
func (*Comic) TableName() string {
	return TableNameComic
}
