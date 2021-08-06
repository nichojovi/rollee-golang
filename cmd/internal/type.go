package internal

import (
	"github.com/nichojovi/rollee-test/internal/service"
)

type (
	Service struct {
		User      service.UserService
		Fibonacci service.FibonacciService
	}
)
