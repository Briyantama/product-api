package repository

import (
	"context"
	"database/sql"
	"test-case-vhiweb/internal/constants"
	"test-case-vhiweb/internal/logger"
	"test-case-vhiweb/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	db := ExtractTx(ctx, r.db)

	err := db.Create(user).Error
	if err != nil {
		logger.Log.Error(err)
		return constants.New(
			constants.ERRBADREQUEST,
			constants.ErrUserRegistrationFail,
		)
	}

	return nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	db := ExtractTx(ctx, r.db)

	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		logger.Log.Error(err)
		if err == sql.ErrNoRows {
			return nil, constants.New(
				constants.ERRNOTFOUND,
				constants.ErrUserNotFound,
			)
		}
		return nil, constants.New(
			constants.ERRBADREQUEST,
			constants.ErrUserInvalid,
		)
	}

	return &user, nil
}
