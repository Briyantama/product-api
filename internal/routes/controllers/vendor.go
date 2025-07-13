package controllers

import (
	"net/http"
	"test-case-vhiweb/internal/dtos"
	"test-case-vhiweb/internal/routes/usecase"
	"test-case-vhiweb/internal/utils"

	"github.com/gin-gonic/gin"
)

type VendorController struct {
	usecase usecase.VendorUsecase
}

func NewVendorController(usecase usecase.VendorUsecase) *VendorController {
	return &VendorController{usecase}
}

func (vc *VendorController) RegisterVendor(c *gin.Context) {
	var input dtos.VendorRegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	userID := c.MustGet("userID").(uint)

	err := vc.usecase.RegisterVendor(c, input.NameVendor, userID)
	if err != nil {
		c.Error(err)
		return
	}

	utils.JSONResponse(c, http.StatusCreated, gin.H{"message": "Vendor registered"})
}

func (vc *VendorController) GetVendorsByUserID(c *gin.Context) {
	var email string
	userID := c.MustGet("userID").(uint)

	vendors, err := vc.usecase.GetVendorsByUser(c, userID)
	if err != nil {
		c.Error(err)
		return
	}

	var vendorResponses []dtos.VendorResponse
	for _, v := range vendors {
		email = v.User.Email
		vendorResponses = append(vendorResponses, dtos.VendorResponse{
			Name: v.Name,
		})
	}

	res := dtos.UserWithVendorsResponse{
		User: dtos.UserResponse{
			EmailUser: email,
		},
		Vendors: vendorResponses,
	}

	utils.JSONResponse(c, http.StatusOK, res)
}
