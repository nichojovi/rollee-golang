package api

import (
	"net/http"

	"github.com/nichojovi/rollee-test/cmd/internal"
	"github.com/nichojovi/rollee-test/internal/service"
	"github.com/nichojovi/rollee-test/internal/utils/auth"
	"github.com/nichojovi/rollee-test/internal/utils/response"
	"github.com/nichojovi/rollee-test/internal/utils/router"
)

type Options struct {
	Prefix         string
	DefaultTimeout int
	AuthService    auth.AuthService
	Service        *internal.Service
}

type API struct {
	options          *Options
	authService      auth.AuthService
	userService      service.UserService
	fibonacciService service.FibonacciService
}

func New(o *Options) *API {
	return &API{
		options:          o,
		authService:      o.AuthService,
		userService:      o.Service.User,
		fibonacciService: o.Service.Fibonacci,
	}
}

func (a *API) Register() {
	r := router.New(&router.Options{Timeout: a.options.DefaultTimeout, Prefix: a.options.Prefix})

	// Test
	r.GET("/ping", a.Ping)

	// CRUD
	r.GET("/user", a.authService.Authorize(a.GetUser))
	r.POST("/insert-user", a.authService.Authorize(a.InsertUser))
	r.PUT("/update-user-phone", a.authService.Authorize(a.UpdateUserPhone))
	r.DELETE("/delete-user", a.authService.Authorize(a.DeleteUser))

	// Fibonaci
	r.GET("/fibonacci", a.authService.Authorize(a.GetFibonacci))
}

func (a *API) Ping(w http.ResponseWriter, r *http.Request) *response.JSONResponse {
	return response.NewJSONResponse().SetMessage("pong")
}
