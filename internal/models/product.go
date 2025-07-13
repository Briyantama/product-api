package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string          `json:"name" gorm:"unique"`
	Price    decimal.Decimal `json:"price"`
	VendorID uint            `json:"vendor_id"`
	Vendor   Vendor          `json:"vendor" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
