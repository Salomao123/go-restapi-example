package service

import (
	"context"

	"github.com/Salomao123/go-restapi-example/models"
	"github.com/Salomao123/go-restapi-example/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo}
}

func (s *UserService) Create(ctx context.Context, user models.User) error {
	err := s.userRepo.Create(ctx, user)

	if err != nil {
		panic(err.Error())
	}

	return nil
}

func (s *UserService) GetAll(ctx context.Context) ([]models.User, error) {
	users, err := s.userRepo.FindAll(ctx)

	if err != nil {
		panic(err.Error())
	}

	return users, nil
}
