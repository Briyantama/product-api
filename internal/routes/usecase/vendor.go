package usecase

import (
	"context"
	"test-case-vhiweb/internal/constants"
	"test-case-vhiweb/internal/models"
	"test-case-vhiweb/internal/routes/repository"
)

type VendorUsecase interface {
	RegisterVendor(ctx context.Context, name string, userID uint) error
	GetVendorsByUser(ctx context.Context, userID uint) ([]models.Vendor, error)
}

type vendorUsecase struct {
	repo repository.VendorRepository
	tx   repository.WithTx
}

func NewVendorUsecase(
	repo repository.VendorRepository,
	tx repository.WithTx,
) VendorUsecase {
	return &vendorUsecase{repo: repo, tx: tx}
}

func (u *vendorUsecase) RegisterVendor(ctx context.Context, name string, userID uint) error {
	vendor := &models.Vendor{
		Name:   name,
		UserID: userID,
	}

	err := u.tx.WithTx(ctx, func(ctx context.Context) error {
		vendorExist, err := u.repo.GetVendorByName(ctx, name)
		if err != nil {
			return err
		}

		if vendorExist != nil {
			return constants.New(
				constants.ERRCONFLICT,
				constants.ErrVendorAlreadyExists,
			)
		}

		err = u.repo.Create(ctx, vendor)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func (u *vendorUsecase) GetVendorsByUser(ctx context.Context, userID uint) ([]models.Vendor, error) {
	return u.repo.GetByUserID(ctx, userID)
}
