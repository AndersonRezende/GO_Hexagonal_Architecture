package memory

import (
	"errors"
	"gohexarc/internal/domain"
)

type InMemoryUserRepository struct {
	users map[string]domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{users: make(map[string]domain.User)}
}

func (r *InMemoryUserRepository) Create(user domain.User) error {
	if _, exists := r.users[user.ID]; exists {
		return errors.New("user already exists")
	}
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) GetByID(id string) (domain.User, error) {
	user, ok := r.users[id]
	if !ok {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func (r *InMemoryUserRepository) Update(user domain.User) error {
	if _, exists := r.users[user.ID]; !exists {
		return errors.New("user not found")
	}
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) Delete(id string) error {
	if _, exists := r.users[id]; !exists {
		return errors.New("user not found")
	}
	delete(r.users, id)
	return nil
}

func (r *InMemoryUserRepository) List() ([]domain.User, error) {
	users := make([]domain.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}
