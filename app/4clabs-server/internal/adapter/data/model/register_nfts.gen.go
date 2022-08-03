// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameRegisterNft = "register_nfts"

// RegisterNft mapped from table <register_nfts>
type RegisterNft struct {
	ID              int32     `gorm:"column:id;type:int(11) unsigned;primaryKey;autoIncrement:true" json:"id"`
	TokenID         string    `gorm:"column:token_id;type:varchar(200);not null;uniqueIndex:uniq_idx_token_id_contract_address_user_address,priority:1;uniqueIndex:uniq_idx_contract_address_token_id,priority:2" json:"token_id"`
	Name            string    `gorm:"column:name;type:varchar(200)" json:"name"`
	CollectionName  string    `gorm:"column:collection_name;type:varchar(200)" json:"collection_name"`
	ContractAddress string    `gorm:"column:contract_address;type:varchar(200);not null;uniqueIndex:uniq_idx_token_id_contract_address_user_address,priority:2;uniqueIndex:uniq_idx_contract_address_token_id,priority:1" json:"contract_address"`
	UserAddress     string    `gorm:"column:user_address;type:varchar(200);not null;uniqueIndex:uniq_idx_token_id_contract_address_user_address,priority:3;index:idx_user_address,priority:1" json:"user_address"`
	Price           float64   `gorm:"column:price;type:double;index:idx_price,priority:1" json:"price"`
	Image           string    `gorm:"column:image;type:varchar(200)" json:"image"`
	CreatedAt       time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName RegisterNft's table name
func (*RegisterNft) TableName() string {
	return TableNameRegisterNft
}
