package repository

import (
	"context"
	"time"
	"backend_dev_task/db/sqlc"
	"backend_dev_task/internal/models"
)

type UserRepository struct {
	queries *sqlc.Queries
}

func NewUserRepository(queries *sqlc.Queries) *UserRepository {
	return &UserRepository{queries: queries}
}

func (r *UserRepository) CreateUser(ctx context.Context, name, dob string) (*models.User, error) {
	parsedDOB, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return nil, err
	}
	user, err := r.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Name: name,
		Dob:  parsedDOB,
	})
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:   int(user.ID),
		Name: user.Name,
		DOB:  user.Dob,
	}, nil
}

func (r *UserRepository) GetUser(ctx context.Context, id int) (*models.User, error) {
	user, err := r.queries.GetUser(ctx, int32(id))
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:   int(user.ID),
		Name: user.Name,
		DOB:  user.Dob,
	}, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, id int, name, dob string) (*models.User, error) {
	parsedDOB, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return nil, err
	}
	user, err := r.queries.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   int32(id),
		Name: name,
		Dob:  parsedDOB,
	})
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:   int(user.ID),
		Name: user.Name,
		DOB:  user.Dob,
	}, nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int) error {
	return r.queries.DeleteUser(ctx, int32(id))
}

func (r *UserRepository) ListUsers(ctx context.Context, limit, offset int) ([]*models.User, error) {
	users, err := r.queries.ListUsers(ctx, sqlc.ListUsersParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	var result []*models.User
	for _, u := range users {
		result = append(result, &models.User{
			ID:   int(u.ID),
			Name: u.Name,
			DOB:  u.Dob,
		})
	}
	return result, nil
}