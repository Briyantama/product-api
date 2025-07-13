package usecase

import (
	"context"
	"test-case-vhiweb/internal/constants"
	"test-case-vhiweb/internal/models"
	"test-case-vhiweb/internal/routes/repository"
	"test-case-vhiweb/internal/utils"
)

type UserUsecase interface {
	Register(ctx context.Context, user *models.User) error
	Login(ctx context.Context, email, password string) (*uint, error)
}

type userUsecase struct {
	repo repository.UserRepository
	tx   repository.WithTx
}

func NewUserUsecase(
	r repository.UserRepository,
	tx repository.WithTx,
) UserUsecase {
	return &userUsecase{repo: r, tx: tx}
}

func (u *userUsecase) Register(ctx context.Context, user *models.User) error {
	return u.tx.WithTx(ctx, func(txCtx context.Context) error {
		hashed, err := utils.GeneratePassword(user.Password)
		if err != nil {
			return constants.New(
				constants.ERRBADREQUEST,
				constants.ErrUserInvalid,
			)
		}

		user.Password = string(hashed)
		err = u.repo.CreateUser(txCtx, user)
		if err != nil {
			return err
		}

		return nil
	})
}
func (u *userUsecase) Login(ctx context.Context, email, password string) (*uint, error) {
	var res *uint

	err := u.tx.WithTx(ctx, func(txCtx context.Context) error {
		user, err := u.repo.GetUserByEmail(txCtx, email)
		if err != nil {
			return err
		}

		err = utils.CompareHashAndPassword(user.Password, password)
		if err != nil {
			return constants.New(
				constants.ERRBADREQUEST,
				constants.ErrUserInvalid,
			)
		}

		res = &user.ID
		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
