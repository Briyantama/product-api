package dtos

import "github.com/shopspring/decimal"

type UserResponse struct {
	EmailUser string `json:"email_user"`
}

type ProductResponse struct {
	Name  string          `json:"name"`
	Price decimal.Decimal `json:"price"`
}

type VendorResponse struct {
	Name     string            `json:"name"`
	Products []ProductResponse `json:"products,omitempty"`
}

type UserWithVendorsResponse struct {
	User    UserResponse     `json:"user"`
	Vendors []VendorResponse `json:"vendors"`
}
