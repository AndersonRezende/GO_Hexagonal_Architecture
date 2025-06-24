package mock

import (
	"gohexarc/internal/domain"

	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (m *UserRepository) Create(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepository) GetByID(id string) (domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *UserRepository) Update(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *UserRepository) List() ([]domain.User, error) {
	args := m.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}
