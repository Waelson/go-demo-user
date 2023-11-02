package service

import (
	"context"
	"github.com/Waelson/go-demo-user/internal/model"
	"github.com/Waelson/go-demo-user/internal/repository"
	apierrors "github.com/Waelson/go-demo-user/pkg/errors"
	"github.com/google/uuid"
)

type UserService interface {
	Save(ctx context.Context, user model.User) (*model.User, error)
	Update(ctx context.Context, user model.User) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func (s *userService) Save(ctx context.Context, user model.User) (*model.User, error) {
	user.Id = uuid.New().String()
	newUser, err := s.userRepository.Save(user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
func (s *userService) Update(ctx context.Context, user model.User) error {
	exists, err := s.userRepository.Exists(user.Id)
	if err != nil {
		return err
	}

	if !exists {
		return apierrors.NewServiceError("User not found")
	}

	err = s.userRepository.Update(user)
	return err
}

func (s *userService) Delete(ctx context.Context, id string) error {
	exists, err := s.userRepository.Exists(id)
	if err != nil {
		return err
	}

	if !exists {
		return apierrors.NewServiceError("User not found")
	}
	return nil

	err = s.userRepository.Delete(id)
	return err
}

func (s *userService) Get(ctx context.Context, id string) (*model.User, error) {
	user, err := s.userRepository.Find(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserService(userRepository repository.UserRepository) UserService {
	result := &userService{
		userRepository: userRepository,
	}
	return result
}
