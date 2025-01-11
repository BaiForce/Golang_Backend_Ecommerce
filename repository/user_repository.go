package repository

import (
	"database/sql"

	"github.com/Backend_Shoppe/models"
)

// UserRepository defines the interface for user-related operations
type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateJwtToken(user *models.User) error
}

// userRepository is the concrete implementation of UserRepository
type userRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new instance of userRepository
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}


// GetUserByEmail retrieves a user by email
func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := "SELECT id, name, email, password, jwt_token FROM users WHERE email = ?"
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.JwtToken)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user in the database
func (r *userRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	// Ambil ID yang baru saja dibuat
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)
	return nil
}


// UpdateJwtToken updates the JWT token for the user in the database
func (r *userRepository) UpdateJwtToken(user *models.User) error {
	query := "UPDATE users SET jwt_token = ? WHERE id = ?"
	_, err := r.db.Exec(query, user.JwtToken, user.ID)
	return err
}


