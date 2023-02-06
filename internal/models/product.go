package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Data         string `gorm:"type: character varying;not null"`
	Name         string `gorm:"type:character varying(30);not null"`
	Desc         string `gorm:"type:character varying"`
	CollectionID int64  `gorm:"type:bigint;not null"`
	BlockchainID int64  `gorm:"type:bigint;not null"`
}
