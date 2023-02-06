package models

import "gorm.io/gorm"

type BlockchainStatus string

var (
	BlockchainStatusActive   = BlockchainStatus("active")
	BlockchainStatusDisabled = BlockchainStatus("disabled")
)

type Blockchain struct {
	gorm.Model
	Name   string           `gorm:"type:character varying(20);not null"`
	Status BlockchainStatus `gorm:"type:character varying(10);not null"`
}
