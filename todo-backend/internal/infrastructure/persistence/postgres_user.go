package persistence

import (
	"todo-app/internal/domain/user"

	"gorm.io/gorm"
)

// implementation of user.Repository based on GORM/Postgres
type PostgresUserRepo struct {
	db *gorm.DB
}

// creates a new user repository using the transferred GORM connection
func NewPostgresUserRepo(db *gorm.DB) user.Repository {
	return &PostgresUserRepo{db: db}
}

// saves the new user u to the users table
func (r *PostgresUserRepo) Insert(u *user.User) error {
	return r.db.Create(u).Error
}

// is looking for a user by email
func (r *PostgresUserRepo) GetByEmail(email string) (user.User, error) {
	var u user.User
	err := r.db.Where("email = ?", email).First(&u).Error
	return u, err
}

// searches for a user by their ID
func (r *PostgresUserRepo) GetByID(id int64) (user.User, error) {
	var u user.User
	err := r.db.First(&u, id).Error
	return u, err
}
