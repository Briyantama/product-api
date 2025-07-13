package repository

import (
	"context"
	"database/sql"
	"test-case-vhiweb/internal/constants"
	"test-case-vhiweb/internal/logger"
	"test-case-vhiweb/internal/models"

	"gorm.io/gorm"
)

type VendorRepository interface {
	Create(ctx context.Context, vendor *models.Vendor) error
	GetByUserID(ctx context.Context, userID uint) ([]models.Vendor, error)
	GetByVendorID(ctx context.Context, vendorID uint) (*models.Vendor, error)
	GetVendorByName(ctx context.Context, name string) (*models.Vendor, error)
}

type vendorRepository struct {
	db *gorm.DB
}

func NewVendorRepository(db *gorm.DB) VendorRepository {
	return &vendorRepository{db: db}
}

func (r *vendorRepository) Create(ctx context.Context, vendor *models.Vendor) error {
	db := ExtractTx(ctx, r.db)

	err := db.Create(vendor).Error
	if err != nil {
		logger.Log.Error(err)
		return constants.New(
			constants.ERRBADREQUEST,
			constants.ErrVendorRegistrationFail,
		)
	}

	return nil
}

func (r *vendorRepository) GetByUserID(ctx context.Context, userID uint) ([]models.Vendor, error) {
	var vendors []models.Vendor
	db := ExtractTx(ctx, r.db)

	err := db.Preload("User").Where("user_id = ?", userID).Find(&vendors).Error
	if err != nil {
		logger.Log.Error(err)
		if err == sql.ErrNoRows {
			return nil, constants.New(
				constants.ERRNOTFOUND,
				constants.ErrVendorNotFound,
			)
		}

		return nil, constants.New(
			constants.ERRBADREQUEST,
			constants.ErrGetVendorByUserIDFail,
		)
	}

	return vendors, nil
}

func (r *vendorRepository) GetVendorByName(ctx context.Context, name string) (*models.Vendor, error) {
	var vendor *models.Vendor
	db := ExtractTx(ctx, r.db)

	err := db.Preload("User").Where("name = ?", name).First(&vendor).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logger.Log.Error(err)
		return nil, constants.New(
			constants.ERRBADREQUEST,
			constants.ErrVendorAlreadyExists,
		)
	}

	return vendor, nil
}

func (r *vendorRepository) GetByVendorID(ctx context.Context, vendorID uint) (*models.Vendor, error) {
	var vendor *models.Vendor
	db := ExtractTx(ctx, r.db)

	err := db.Preload("User").Where("id = ?", vendorID).Find(&vendor).Error
	if err != nil {
		logger.Log.Error(err)
		if err == sql.ErrNoRows {
			return nil, constants.New(
				constants.ERRNOTFOUND,
				constants.ErrVendorNotFound,
			)
		}

		return nil, constants.New(
			constants.ERRBADREQUEST,
			constants.ErrGetVendorByUserIDFail,
		)
	}

	return vendor, nil
}
