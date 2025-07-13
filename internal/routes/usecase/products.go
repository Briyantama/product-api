package usecase

import (
	"context"
	"test-case-vhiweb/internal/dtos"
	"test-case-vhiweb/internal/models"
	"test-case-vhiweb/internal/routes/repository"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, userID uint, product dtos.ProductRegisterRequest) error
	GetProductsByUserID(ctx context.Context, userID uint) ([]models.Product, error)
	GetProductsByVendorID(ctx context.Context, userID uint) ([]models.Product, error)
	UpdateProduct(ctx context.Context, product *dtos.UpdateProduct) error
	DeleteProduct(ctx context.Context, id uint) error
}

type productUsecase struct {
	repoProduct repository.ProductRepository
	repoVendor  repository.VendorRepository
	tx          repository.WithTx
}

func NewProductUsecase(
	repoProduct repository.ProductRepository,
	repoVendor repository.VendorRepository,
	tx repository.WithTx,
) ProductUsecase {
	return &productUsecase{
		repoProduct: repoProduct,
		repoVendor:  repoVendor,
		tx:          tx,
	}
}

func (uc *productUsecase) CreateProduct(ctx context.Context, userID uint, dto dtos.ProductRegisterRequest) error {
	return uc.tx.WithTx(ctx, func(txCtx context.Context) error {
		product := &models.Product{
			Name:  dto.NameProducts,
			Price: dto.Price,
		}

		vendor, err := uc.repoVendor.GetVendorByName(txCtx, dto.NameVendor)
		product.VendorID = vendor.ID

		err = uc.repoProduct.Create(ctx, product)
		if err != nil {
			return err
		}

		return nil
	})
}

func (uc *productUsecase) GetProductsByUserID(ctx context.Context, userID uint) ([]models.Product, error) {
	var products []models.Product

	err := uc.tx.WithTx(ctx, func(txCtx context.Context) error {
		vendor, err := uc.repoVendor.GetByUserID(ctx, userID)
		if err != nil {
			return err
		}

		for _, v := range vendor {
			res, err := uc.repoProduct.FindByVendor(ctx, v.ID)
			if err != nil {
				return err
			}

			products = append(products, res...)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (uc *productUsecase) GetProductsByVendorID(ctx context.Context, userID uint) ([]models.Product, error) {
	var products []models.Product

	err := uc.tx.WithTx(ctx, func(txCtx context.Context) error {
		vendor, err := uc.repoVendor.GetByVendorID(ctx, userID)
		if err != nil {
			return err
		}

		res, err := uc.repoProduct.FindByVendor(ctx, vendor.ID)
		if err != nil {
			return err
		}

		products = append(products, res...)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (uc *productUsecase) UpdateProduct(ctx context.Context, product *dtos.UpdateProduct) error {
	return uc.repoProduct.Update(ctx, product)
}

func (uc *productUsecase) DeleteProduct(ctx context.Context, id uint) error {
	return uc.repoProduct.Delete(ctx, id)
}
