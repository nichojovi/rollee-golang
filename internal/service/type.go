package service

import (
	"context"

	"github.com/nichojovi/rollee-test/cmd/config"
	"github.com/nichojovi/rollee-test/internal/entity"
	"github.com/nichojovi/rollee-test/internal/repository"
)

type (
	userService struct {
		cfg      *config.MainConfig
		userRepo repository.UserRepository
	}

	fibonacciService struct {
		cfg *config.MainConfig
	}
)

type (
	UserService interface {
		GetUserAuth(ctx context.Context, username, password string) (*entity.User, error)
		GetUserByID(ctx context.Context, userID int64) (*entity.User, error)
		InsertUser(ctx context.Context, userData entity.User) error
		UpdateUserPhone(ctx context.Context, userID int64, phone string) error
		DeleteUserByID(ctx context.Context, userID int64) error
	}

	FibonacciService interface {
		GetFibonacci(n int) int
	}
)
