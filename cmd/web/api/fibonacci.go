package api

import (
	"encoding/json"
	"net/http"

	"github.com/nichojovi/rollee-test/internal/entity"
	"github.com/nichojovi/rollee-test/internal/utils/response"
	opentracing "github.com/opentracing/opentracing-go"
)

func (a *API) GetFibonacci(w http.ResponseWriter, r *http.Request) *response.JSONResponse {
	span, _ := opentracing.StartSpanFromContext(r.Context(), "api.GetFibonacci")
	defer span.Finish()

	var request entity.Fibonacci
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}

	n := a.fibonacciService.GetFibonacci(request.N)

	return response.NewJSONResponse().SetData(n)
}
