// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID        int32     `gorm:"column:id;type:int(20) unsigned;primaryKey;autoIncrement:true" json:"id"`
	Address   string    `gorm:"column:address;type:varchar(100);not null;uniqueIndex:uniq_idx_address,priority:1" json:"address"`
	Nonce     string    `gorm:"column:nonce;type:varchar(200);not null" json:"nonce"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
