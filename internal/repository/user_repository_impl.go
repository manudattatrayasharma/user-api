package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	db "user-api/db/sqlc"
)

type userRepository struct {
	q *db.Queries
}

func NewUserRepository(dbConn *sql.DB) UserRepository {
	return &userRepository{
		q: db.New(dbConn),
	}
}

func (r *userRepository) Create(
	ctx context.Context,
	name string,
	dob time.Time,
) (db.User, error) {

	result, err := r.q.CreateUser(ctx, db.CreateUserParams{
		Name: name,
		Dob:  dob, // ✅ time.Time
	})
	if err != nil {
		return db.User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return db.User{}, err
	}

	return r.q.GetUserByID(ctx, id)
}

func (r *userRepository) GetByID(
	ctx context.Context,
	id int64,
) (db.User, error) {

	user, err := r.q.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return db.User{}, err
		}
		return db.User{}, err
	}
	return user, nil
}

func (r *userRepository) Update(
	ctx context.Context,
	id int64,
	name string,
	dob time.Time,
) (db.User, error) {

	err := r.q.UpdateUser(ctx, db.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob, // ✅ time.Time
	})
	if err != nil {
		return db.User{}, err
	}

	return r.q.GetUserByID(ctx, id)
}

func (r *userRepository) Delete(
	ctx context.Context,
	id int64,
) error {
	return r.q.DeleteUser(ctx, id)
}

func (r *userRepository) List(
	ctx context.Context,
	limit, offset int32,
) ([]db.User, error) {
	return r.q.ListUsers(ctx, db.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
}
