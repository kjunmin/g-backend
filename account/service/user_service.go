package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/kjunmin/g-backend/model"
)

type UserService struct {
	UserRepository model.UserRepository
}

func (s *UserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	u, err := s.UserRepository.FindByID(ctx, uid)
	return u, err
}

type USConfig struct {
	UserRepository model.UserRepository
}

func NewUserService(c *USConfig) model.UserService {
	return &UserService{
		UserRepository: c.UserRepository,
	}
}
