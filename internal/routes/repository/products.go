package repository

import (
	"context"
	"database/sql"
	"test-case-vhiweb/internal/constants"
	"test-case-vhiweb/internal/dtos"
	"test-case-vhiweb/internal/logger"
	"test-case-vhiweb/internal/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(ctx context.Context, product *models.Product) error
	FindByVendor(ctx context.Context, vendorID uint) ([]models.Product, error)
	Update(ctx context.Context, product *dtos.UpdateProduct) error
	Delete(ctx context.Context, id uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) Create(ctx context.Context, product *models.Product) error {
	db := ExtractTx(ctx, r.db)

	err := db.Create(product).Error
	if err != nil {
		logger.Log.Error(err)
		return constants.New(
			constants.ERRBADREQUEST,
			constants.ErrProductCreationFail,
		)
	}

	return nil
}

func (r *productRepository) FindByVendor(ctx context.Context, vendorID uint) ([]models.Product, error) {
	var products []models.Product
	db := ExtractTx(ctx, r.db)

	err := db.Preload("Vendor.User").Where("vendor_id = ?", vendorID).
		Find(&products).Error
	if err != nil {
		logger.Log.Error(err)
		if err == sql.ErrNoRows {
			return nil, constants.New(
				constants.ERRNOTFOUND,
				constants.ErrProductNotFound,
			)
		}

		return nil, constants.New(
			constants.ERRBADREQUEST,
			constants.ErrProductInvalidPayload,
		)
	}

	return products, err
}

func (r *productRepository) Update(ctx context.Context, product *dtos.UpdateProduct) error {
	db := ExtractTx(ctx, r.db)

	err := db.Model(&models.Product{}).
		Where("id = ?", product.ID).
		Updates(map[string]interface{}{
			"name":  product.Name,
			"price": product.Price,
		}).Error

	if err != nil {
		logger.Log.Error(err)
		return constants.New(
			constants.ERRBADREQUEST,
			constants.ErrProductUpdateFail,
		)
	}

	return nil
}

func (r *productRepository) Delete(ctx context.Context, id uint) error {
	db := ExtractTx(ctx, r.db)

	err := db.Delete(&models.Product{}, id).Error
	if err != nil {
		logger.Log.Error(err)
		return constants.New(
			constants.ERRBADREQUEST,
			constants.ErrProductDeleteFail,
		)
	}

	return nil
}
