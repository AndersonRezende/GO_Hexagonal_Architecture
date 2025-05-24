package port

import "gohexarc/internal/domain"

type UserRepository interface {
	Create(user domain.User) error
	GetByID(id string) (domain.User, error)
	Update(user domain.User) error
	Delete(id string) error
	List() ([]domain.User, error)
}

type UserService interface {
	CreateUser(name, email string) (domain.User, error)
	GetUser(id string) (domain.User, error)
	UpdateUser(id, name, email string) error
	DeleteUser(id string) error
	ListUsers() ([]domain.User, error)
}
