package models

import "gorm.io/gorm"

type Collection struct {
	gorm.Model
	Name         string `gorm:"type:character varying(20);not null"`
	Desc         string `gorm:"type:character varying"`
	Url          string `gorm:"type:character varying(50)"`
	LogoURL      string `gorm:"type:character varying"`
	FeaturedURL  string `gorm:"type:character varying"`
	BannerURL    string `gorm:"type:character varying"`
	CategoryID   int64  `gorm:"type:bigint;not null"`
	BlockchainID int64  `gorm:"type:bigint;not null"`
}
