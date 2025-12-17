package service

import (
	"context"
	"backend_dev_task/internal/models"
	"backend_dev_task/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.User, error) {
	return s.repo.CreateUser(ctx, req.Name, req.DOB)
}

func (s *UserService) GetUser(ctx context.Context, id int) (*models.UserWithAge, error) {
	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	userWithAge := &models.UserWithAge{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.DOB.Format("2006-01-02"),
	}
	userWithAge.CalculateAge()
	return userWithAge, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int, req *models.UpdateUserRequest) (*models.User, error) {
	return s.repo.UpdateUser(ctx, id, req.Name, req.DOB)
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	return s.repo.DeleteUser(ctx, id)
}

func (s *UserService) ListUsers(ctx context.Context, limit, offset int) ([]*models.UserWithAge, error) {
	users, err := s.repo.ListUsers(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	var result []*models.UserWithAge
	for _, u := range users {
		userWithAge := &models.UserWithAge{
			ID:   u.ID,
			Name: u.Name,
			DOB:  u.DOB.Format("2006-01-02"),
		}
		userWithAge.CalculateAge()
		result = append(result, userWithAge)
	}
	return result, nil
}