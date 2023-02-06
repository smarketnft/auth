package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ActivityAction string

var (
	ActivityActionMinted = ActivityAction("minted")
	ActivityActionList   = ActivityAction("list")
)

type Activity struct {
	gorm.Model
	Action ActivityAction  `gorm:"type:character varying(20);not null"`
	Price  decimal.Decimal `gorm:"type:numeric(32,16)"`
	Amount int64           `gorm:"type:bigint"`
	FromID int64           `gorm:"bigint"`
	ToID   int64           `gorm:"bigint"`
}
