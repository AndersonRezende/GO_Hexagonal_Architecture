package service

import (
	"gohexarc/internal/domain"
	"gohexarc/internal/port"
)

type UserServiceImpl struct {
	repository port.UserRepository
}

func NewUserService(repository port.UserRepository) port.UserService {
	return &UserServiceImpl{repository}
}

func (u *UserServiceImpl) CreateUser(name, email string) (domain.User, error) {
	user := domain.User{Name: name, Email: email}
	return user, u.repository.Create(user)
}

func (u *UserServiceImpl) GetUser(id string) (domain.User, error) {
	return u.repository.GetByID(id)
}

func (u *UserServiceImpl) UpdateUser(id, name, email string) error {
	user := domain.User{ID: id, Name: name, Email: email}
	return u.repository.Update(user)
}

func (u *UserServiceImpl) DeleteUser(id string) error {
	return u.repository.Delete(id)
}

func (u *UserServiceImpl) ListUsers() ([]domain.User, error) {
	return u.repository.List()
}
