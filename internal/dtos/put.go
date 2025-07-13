package dtos

import "github.com/shopspring/decimal"

type UpdateProduct struct {
	ID    int
	Name  string          `json:"name" gorm:"unique"`
	Price decimal.Decimal `json:"price"`
}
