package repository

import (
	"context"
	"time"

	db "user-api/db/sqlc"
)

type UserRepository interface {
	Create(ctx context.Context, name string, dob time.Time) (db.User, error)
	GetByID(ctx context.Context, id int64) (db.User, error)
	Update(ctx context.Context, id int64, name string, dob time.Time) (db.User, error)
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context, limit, offset int32) ([]db.User, error)
}
