package services

import (
	"github.com/br4tech/go-user-create/domain"
)

type UserService struct {
	repository domain.UserRepository[domain.User]
}

func NewUserService(repository domain.UserRepository[domain.User]) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) CreateUser(user *domain.User) error {
	return s.repository.Create(user)
}

func (s *UserService) GetUserByID(id uint) (*domain.User, error) {
	return s.repository.FindByID(id)
}

func (s *UserService) UpdateUser(user *domain.User) error {
	return s.repository.Update(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repository.Delete(id)
}
