package service

import (
	"context"
	"errors"
	"time"

	db "user-api/db/sqlc"
)

// UserService handles business logic
type UserService struct {
	queries *db.Queries
}

// Constructor
func NewUserService(queries *db.Queries) *UserService {
	return &UserService{
		queries: queries,
	}
}

// Domain model returned to handlers
type User struct {
	ID   int64     `json:"id"`
	Name string    `json:"name"`
	Dob  time.Time `json:"dob"`
	Age  int       `json:"age"`
}

// validateDOB ensures DOB is not in the future
func validateDOB(dob time.Time) error {
	if dob.After(time.Now()) {
		return errors.New("date of birth cannot be in the future")
	}
	return nil
}

// CalculateAge calculates age from DOB
func CalculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	// Adjust if birthday hasn’t occurred yet this year
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}

// CreateUser creates a new user after validation
func (s *UserService) CreateUser(
	ctx context.Context,
	name string,
	dob time.Time,
) (*User, error) {

	if err := validateDOB(dob); err != nil {
		return nil, err
	}

	result, err := s.queries.CreateUser(ctx, db.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &User{
		ID:   id,
		Name: name,
		Dob:  dob,
		Age:  CalculateAge(dob),
	}, nil
}

// GetUserByID fetches user and calculates age
func (s *UserService) GetUserByID(ctx context.Context, id int64) (*User, error) {
	u, err := s.queries.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:   u.ID,
		Name: u.Name,
		Dob:  u.Dob,
		Age:  CalculateAge(u.Dob),
	}, nil
}

// ListUsers returns users with calculated age
func (s *UserService) ListUsers(ctx context.Context, limit, offset int32) ([]User, error) {
	users, err := s.queries.ListUsers(ctx, db.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	result := make([]User, 0, len(users))
	for _, u := range users {
		result = append(result, User{
			ID:   u.ID,
			Name: u.Name,
			Dob:  u.Dob,
			Age:  CalculateAge(u.Dob),
		})
	}

	return result, nil
}

// UpdateUser updates user after validation
func (s *UserService) UpdateUser(
	ctx context.Context,
	id int64,
	name string,
	dob time.Time,
) (*User, error) {

	if err := validateDOB(dob); err != nil {
		return nil, err
	}

	err := s.queries.UpdateUser(ctx, db.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob,
	})
	if err != nil {
		return nil, err
	}

	return &User{
		ID:   id,
		Name: name,
		Dob:  dob,
		Age:  CalculateAge(dob),
	}, nil
}

// DeleteUser deletes a user
func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	return s.queries.DeleteUser(ctx, id)
}
