package repository

import (
	"context"

	"github.com/nichojovi/rollee-test/cmd/config"
	"github.com/nichojovi/rollee-test/internal/entity"
	"github.com/nichojovi/rollee-test/internal/utils/database"
)

type (
	userRepo struct {
		db  *database.Store
		cfg *config.MainConfig
	}
)

type (
	UserRepository interface {
		GetUserAuth(ctx context.Context, username, password string) (*entity.User, error)
		GetUserByID(ctx context.Context, id int64) (*entity.User, error)
		InsertUser(ctx context.Context, data entity.User) error
		UpdateUserPhoneByID(ctx context.Context, id int64, phone string) error
		DeleteUserByID(ctx context.Context, id int64) error
	}
)
