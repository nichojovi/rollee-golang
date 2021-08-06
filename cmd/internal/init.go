package internal

import (
	"github.com/nichojovi/rollee-test/cmd/config"
	"github.com/nichojovi/rollee-test/internal/repository"
	"github.com/nichojovi/rollee-test/internal/service"
	"github.com/nichojovi/rollee-test/internal/utils/database"
)

func GetService(db *database.Store, config *config.MainConfig) *Service {
	//REPO
	userRepository := repository.NewUserRepository(db, config)

	//SERVICE
	userService := service.NewUserService(userRepository, config)
	fibonacciService := service.NewFibonacciService(config)

	return &Service{
		User:      userService,
		Fibonacci: fibonacciService,
	}
}
