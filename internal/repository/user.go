package repository

import (
	"database/sql"
	"github.com/Waelson/go-demo-user/internal/model"
)

type UserRepository interface {
	Find(id string) (*model.User, error)
	Save(user model.User) (*model.User, error)
	Update(user model.User) error
	Delete(id string) error
	Exists(id string) (bool, error)
}

type userRepository struct {
	db *sql.DB
}

func (r *userRepository) Find(id string) (*model.User, error) {
	return nil, nil
}

func (r *userRepository) Exists(id string) (bool, error) {
	return true, nil
}

func (r *userRepository) Save(user model.User) (*model.User, error) {
	return nil, nil
}
func (r *userRepository) Update(user model.User) error {
	return nil
}
func (r *userRepository) Delete(id string) error {
	return nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	result := &userRepository{
		db: db,
	}
	return result
}
