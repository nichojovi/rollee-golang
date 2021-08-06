package service

import (
	"github.com/nichojovi/rollee-test/cmd/config"
	"github.com/nichojovi/rollee-test/internal/repository"
)

func NewUserService(user repository.UserRepository, cfg *config.MainConfig) UserService {
	return &userService{
		cfg:      cfg,
		userRepo: user,
	}
}

func NewFibonacciService(cfg *config.MainConfig) FibonacciService {
	return &fibonacciService{
		cfg: cfg,
	}
}
