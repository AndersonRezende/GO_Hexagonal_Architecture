package sqlite

import (
	"database/sql"
	"gohexarc/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewSqliteUserRepository(db *sql.DB) *UserRepository {
	sqliteUserRepository := &UserRepository{db: db}
	sqliteUserRepository.init()
	return sqliteUserRepository
}

func (r *UserRepository) init() error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE
	);`
	_, err := r.db.Exec(query)
	return err
}

func (r *UserRepository) Create(user domain.User) error {
	query := `INSERT INTO users (name, email) VALUES (?, ?)`
	_, err := r.db.Exec(query, user.Name, user.Email)
	return err
}

func (r *UserRepository) GetByID(id string) (domain.User, error) {
	query := `SELECT id, name, email FROM users WHERE id = ?`
	row := r.db.QueryRow(query, id)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, nil
		}
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Update(user domain.User) error {
	query := `UPDATE users SET name = ?, email = ? WHERE id = ?`
	_, err := r.db.Exec(query, user.Name, user.Email, user.ID)
	return err
}

func (r *UserRepository) Delete(id string) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *UserRepository) List() ([]domain.User, error) {
	query := `SELECT id, name, email FROM users`
	rows, err := r.db.Query(query)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
