package service

import (
	"context"

	"github.com/nichojovi/rollee-test/internal/entity"
	"github.com/opentracing/opentracing-go"
)

func (us *userService) GetUserAuth(ctx context.Context, username, password string) (*entity.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userService.GetUserAuth")
	defer span.Finish()

	var user *entity.User
	user, err := us.userRepo.GetUserAuth(ctx, username, password)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (us *userService) GetUserByID(ctx context.Context, userID int64) (*entity.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userService.GetUserByID")
	defer span.Finish()

	var user *entity.User
	user, err := us.userRepo.GetUserByID(ctx, userID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (us *userService) InsertUser(ctx context.Context, userData entity.User) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userService.InsertUser")
	defer span.Finish()

	err := us.userRepo.InsertUser(ctx, userData)
	if err != nil {
		return err
	}

	return nil
}

func (us *userService) UpdateUserPhone(ctx context.Context, userID int64, phone string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userService.UpdateUserPhone")
	defer span.Finish()

	err := us.userRepo.UpdateUserPhoneByID(ctx, userID, phone)
	if err != nil {
		return err
	}

	return nil
}

func (us *userService) DeleteUserByID(ctx context.Context, userID int64) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userService.DeleteUserByID")
	defer span.Finish()

	err := us.userRepo.DeleteUserByID(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}
