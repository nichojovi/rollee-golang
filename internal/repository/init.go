package repository

import (
	"github.com/nichojovi/rollee-test/cmd/config"
	"github.com/nichojovi/rollee-test/internal/utils/database"
)

func NewUserRepository(db *database.Store, config *config.MainConfig) UserRepository {
	return &userRepo{
		db:  db,
		cfg: config,
	}
}
