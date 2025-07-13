package controllers

import (
	"net/http"
	"strconv"
	"test-case-vhiweb/internal/dtos"
	"test-case-vhiweb/internal/routes/usecase"
	"test-case-vhiweb/internal/utils"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	uc usecase.ProductUsecase
}

func NewProductController(uc usecase.ProductUsecase) *ProductController {
	return &ProductController{uc}
}

func (p *ProductController) Create(c *gin.Context) {
	var product dtos.ProductRegisterRequest
	if err := c.ShouldBindJSON(&product); err != nil {
		c.Error(err)
		return
	}

	userID := c.MustGet("userID").(uint)

	if err := p.uc.CreateProduct(c, userID, product); err != nil {
		c.Error(err)
		return
	}

	utils.JSONResponse(c, http.StatusCreated, gin.H{"message": "Product created"})
}

func (p *ProductController) GetProductByUserID(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	email := ""
	vendorMap := map[string]*dtos.VendorResponse{}

	products, err := p.uc.GetProductsByUserID(c, userID)
	if err != nil {
		c.Error(err)
		return
	}

	for _, p := range products {
		email = p.Vendor.User.Email

		vendorName := p.Vendor.Name

		v, exists := vendorMap[vendorName]
		if !exists {
			v = &dtos.VendorResponse{
				Name:     vendorName,
				Products: []dtos.ProductResponse{},
			}
			vendorMap[vendorName] = v
		}

		v.Products = append(v.Products, dtos.ProductResponse{
			Name:  p.Name,
			Price: p.Price,
		})
	}

	vendors := make([]dtos.VendorResponse, 0, len(vendorMap))
	for _, v := range vendorMap {
		vendors = append(vendors, *v)
	}

	res := dtos.UserWithVendorsResponse{
		User: dtos.UserResponse{
			EmailUser: email,
		},
		Vendors: vendors,
	}

	utils.JSONResponse(c, http.StatusOK, res)
}

func (p *ProductController) GetProductByVendorID(c *gin.Context) {
	idStr := c.Query("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	ID := uint(id)
	vendorMap := map[string]*dtos.VendorResponse{}

	products, err := p.uc.GetProductsByVendorID(c, ID)
	if err != nil {
		c.Error(err)
		return
	}

	for _, p := range products {
		vendorName := p.Vendor.Name

		v, exists := vendorMap[vendorName]
		if !exists {
			v = &dtos.VendorResponse{
				Name:     vendorName,
				Products: []dtos.ProductResponse{},
			}
			vendorMap[vendorName] = v
		}

		v.Products = append(v.Products, dtos.ProductResponse{
			Name:  p.Name,
			Price: p.Price,
		})
	}

	res := make([]dtos.VendorResponse, 0, len(vendorMap))
	for _, v := range vendorMap {
		res = append(res, *v)
	}

	utils.JSONResponse(c, http.StatusOK, res)
}

func (p *ProductController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}

	var product dtos.UpdateProduct
	if err := c.ShouldBindJSON(&product); err != nil {
		c.Error(err)
		return
	}

	product.ID = id
	if err := p.uc.UpdateProduct(c, &product); err != nil {
		c.Error(err)
		return
	}

	utils.JSONResponse(c, http.StatusCreated, gin.H{"message": "Product updated"})
}

func (p *ProductController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(err)
		return
	}

	if err := p.uc.DeleteProduct(c, uint(id)); err != nil {
		c.Error(err)
		return
	}

	utils.JSONResponse(c, http.StatusCreated, gin.H{"message": "Product deleted"})
}
