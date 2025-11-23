package model

import (
	"time"
)

type GenericEntity struct {
	ID         int64     `gorm:"primaryKey;column:id" json:"id"`
	Deleted    bool      `gorm:"column:deleted;not null" json:"deleted"`
	DateCreate time.Time `gorm:"column:date_create;not null;autoCreateTime" json:"dateCreate"`
	DateUpdate time.Time `gorm:"column:date_update;autoUpdateTime" json:"dateUpdate"`
}

// database table name.
func (GenericEntity) TableName() string {
	return "generic_entity"
}
