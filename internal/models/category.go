package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"type:character varying(30);not null"`
}
