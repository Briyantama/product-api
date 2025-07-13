package dtos

import "github.com/shopspring/decimal"

type ProductRegisterRequest struct {
	VendorRegisterRequest
	NameProducts string          `json:"name_product" binding:"required"`
	Price        decimal.Decimal `json:"price" binding:"required"`
}

type VendorRegisterRequest struct {
	NameVendor string `json:"name_vendor" binding:"required"`
}
