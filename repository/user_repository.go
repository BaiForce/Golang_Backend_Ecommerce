package repository

import (
	"database/sql"
	"log"

	"github.com/Backend_Shoppe/models"
)

type UserRepository interface {
    GetUserByEmail(email string) (*models.User, error)
    GetUserByID(id int) (*models.User, error) // Tambahkan ini
    CreateUser(user *models.User) error
    UpdateJwtToken(user *models.User) error
}

type userRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
    var user models.User
    query := "SELECT id, name, email, password, jwt_token FROM users WHERE email = ?"
    err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.JwtToken)
    if err != nil {
        log.Println("Error fetching user by email:", err)
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) GetUserByID(id int) (*models.User, error) {
    var user models.User
    query := "SELECT id, name, email, password, jwt_token FROM users WHERE id = ?"
    err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.JwtToken)
    if err != nil {
        log.Println("Error fetching user by ID:", err)
        return nil, err
    }
    return &user, nil
}

func (r *userRepository) CreateUser(user *models.User) error {
    query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
    result, err := r.db.Exec(query, user.Name, user.Email, user.Password)
    if err != nil {
        log.Println("Error creating user:", err)
        return err
    }

    id, err := result.LastInsertId()
    if err != nil {
        log.Println("Error fetching last insert ID:", err)
        return err
    }
    user.ID = int(id)
    return nil
}

func (r *userRepository) UpdateJwtToken(user *models.User) error {
    query := "UPDATE users SET jwt_token = ? WHERE id = ?"
    _, err := r.db.Exec(query, user.JwtToken, user.ID)
    if err != nil {
        log.Println("Error updating JWT token:", err)
    }
    return err
}
