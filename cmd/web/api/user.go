package api

import (
	"encoding/json"
	"net/http"

	"github.com/nichojovi/rollee-test/internal/entity"
	"github.com/nichojovi/rollee-test/internal/utils/response"
	opentracing "github.com/opentracing/opentracing-go"
)

func (a *API) GetUser(w http.ResponseWriter, r *http.Request) *response.JSONResponse {
	span, ctx := opentracing.StartSpanFromContext(r.Context(), "api.GetUser")
	defer span.Finish()

	var request entity.User
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}

	user, err := a.userService.GetUserByID(ctx, request.ID)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrInternalServerError).SetMessage(err.Error())
	}

	return response.NewJSONResponse().SetData(user)
}

func (a *API) InsertUser(w http.ResponseWriter, r *http.Request) *response.JSONResponse {
	span, ctx := opentracing.StartSpanFromContext(r.Context(), "api.InsertUser")
	defer span.Finish()

	var request entity.User
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}

	err = a.userService.InsertUser(ctx, request)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrInternalServerError).SetMessage(err.Error())
	}

	return response.NewJSONResponse()
}

func (a *API) UpdateUserPhone(w http.ResponseWriter, r *http.Request) *response.JSONResponse {
	span, ctx := opentracing.StartSpanFromContext(r.Context(), "api.UpdateUser")
	defer span.Finish()

	var request entity.User
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}

	err = a.userService.UpdateUserPhone(ctx, request.ID, request.Phone)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrInternalServerError).SetMessage(err.Error())
	}

	return response.NewJSONResponse()
}

func (a *API) DeleteUser(w http.ResponseWriter, r *http.Request) *response.JSONResponse {
	span, ctx := opentracing.StartSpanFromContext(r.Context(), "api.DeleteUser")
	defer span.Finish()

	var request entity.User
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrBadRequest).SetMessage(err.Error())
	}

	err = a.userService.DeleteUserByID(ctx, request.ID)
	if err != nil {
		return response.NewJSONResponse().SetError(response.ErrInternalServerError).SetMessage(err.Error())
	}

	return response.NewJSONResponse()
}
