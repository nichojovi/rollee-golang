package repository

import (
	"context"
	"database/sql"

	"github.com/nichojovi/rollee-test/internal/entity"
	opentracing "github.com/opentracing/opentracing-go"
)

const (
	getAllUserInfoQuery = "SELECT id, username, password, full_name, email, phone FROM user"
	insertUserQuery     = "INSERT INTO user(username, password, full_name, email, phone) VALUES (?, ?, ?, ?, ?)"
	updateUserPhoneByID = "UPDATE user SET phone = ? WHERE id = ?"
	deleteUserByID      = "DELETE FROM user WHERE id = ?"
)

func (ur *userRepo) GetUserAuth(ctx context.Context, username, password string) (*entity.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userRepo.GetUserAuth")
	defer span.Finish()

	query := getAllUserInfoQuery + " WHERE username = ? and password = ?"

	result := new(entity.User)
	err := ur.db.GetSlave().GetContext(ctx, result, query, username, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

func (ur *userRepo) GetUserByID(ctx context.Context, id int64) (*entity.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userRepo.GetUserByID")
	defer span.Finish()

	query := getAllUserInfoQuery + " WHERE id = ?"

	result := new(entity.User)
	err := ur.db.GetSlave().GetContext(ctx, result, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

func (ur *userRepo) InsertUser(ctx context.Context, data entity.User) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userRepo.InsertUser")
	defer span.Finish()

	args := []interface{}{
		data.Username,
		data.Password,
		data.FullName,
		data.Email,
		data.Phone,
	}

	_, err := ur.db.GetMaster().ExecContext(ctx, insertUserQuery, args...)
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepo) UpdateUserPhoneByID(ctx context.Context, id int64, phone string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userRepo.UpdateUserPhoneByID")
	defer span.Finish()

	args := []interface{}{
		phone,
		id,
	}

	_, err := ur.db.GetMaster().ExecContext(ctx, updateUserPhoneByID, args...)
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepo) DeleteUserByID(ctx context.Context, id int64) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userRepo.DeleteUserByID")
	defer span.Finish()

	_, err := ur.db.GetMaster().ExecContext(ctx, deleteUserByID, id)
	if err != nil {
		return err
	}

	return nil
}
