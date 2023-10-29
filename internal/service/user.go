package service

import (
	"context"
	"github.com/Waelson/go-demo-user/internal/model"
	"github.com/Waelson/go-demo-user/internal/repository"
)

type UserService interface {
	Save(ctx context.Context, user model.User) (model.User, error)
	Update(ctx context.Context, user model.User) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func (s *userService) Save(ctx context.Context, user model.User) (model.User, error) {
	return model.User{}, nil
}
func (s *userService) Update(ctx context.Context, user model.User) error {
	return nil
}

func (s *userService) Delete(ctx context.Context, id string) error {
	return nil
}

func (s *userService) Get(ctx context.Context, id string) (model.User, error) {
	return model.User{}, nil
}

func NewUserService(userRepository repository.UserRepository) UserService {
	result := &userService{
		userRepository: userRepository,
	}
	return result
}
