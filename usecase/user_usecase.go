package usecase

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"log"
	"time"

	"github.com/Backend_Shoppe/models"
	"github.com/Backend_Shoppe/repository"
	"github.com/golang-jwt/jwt/v4"
)

func hashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

type UserUsecase interface {
    Login(email, password string) (string, error)
    Register(user *models.User) (*models.User, error)
    GetUserByEmail(email string) (*models.User, error)
    Logout(id int, tokenStr string) error // Gunakan id
}


type userUsecase struct {
	userRepo repository.UserRepository
	secret   string
}

func NewUserUsecase(repo repository.UserRepository, secret string) UserUsecase {
	return &userUsecase{userRepo: repo, secret: secret}
}


func (u *userUsecase) Login(email, password string) (string, error) {
	// Check if the user exists in the database
	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil || user == nil {
		return "", errors.New("invalid credentials")
	}

	// Compare the password with the one in the database
	if password != user.Password {
		return "", errors.New("invalid credentials")
	}

	// Generate a new JWT token (no `exp`)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "id":    user.ID,
    "email": user.Email, // Tambahkan email ke dalam klaim
    "exp":   time.Now().Add(time.Hour * 24).Unix(), // Contoh menambahkan `exp`
})



	// Sign the token
	signedToken, err := token.SignedString([]byte(u.secret))
	if err != nil {
		return "", err
	}

	// Update the JWT token in the database
	user.JwtToken = signedToken
	err = u.userRepo.UpdateJwtToken(user)
	if err != nil {
		return "", err
	}


	// Return the signed token
	return signedToken, nil
}


func (u *userUsecase) GetUserByEmail(email string) (*models.User, error) {
	return u.userRepo.GetUserByEmail(email)
}

func (u *userUsecase) Register(user *models.User) (*models.User, error) {
	// Hash the password before saving it to the database
	hashedPassword := hashPassword(user.Password)
	user.Password = hashedPassword

	// Check if the user already exists
	existingUser, err := u.userRepo.GetUserByEmail(user.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) { // Jika error terjadi (kecuali karena tidak ada data)
		return nil, err
	}

	if existingUser != nil { // Jika pengguna sudah ada
		return nil, errors.New("user already exists")
	}

	// Save the user to the database
	err = u.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	// Generate a new JWT token (no `exp`)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "id":    user.ID,
    "email": user.Email, // Tambahkan email ke dalam klaim
    "exp":   time.Now().Add(time.Hour * 24).Unix(), // Contoh menambahkan `exp`
})


	signedToken, err := token.SignedString([]byte(u.secret))
	if err != nil {
		return nil, err
	}

	// Save the JWT token in the user record
	user.JwtToken = signedToken
	err = u.userRepo.UpdateJwtToken(user)
	if err != nil {
		return nil, err
	}

	// Return the user with the JWT token
	return user, nil
}

func (u *userUsecase) Logout(id int, tokenStr string) error {
    user, err := u.userRepo.GetUserByID(id)
    if err != nil {
        log.Println("User not found for ID:", id)
        return errors.New("user not found")
    }

    if user.JwtToken != tokenStr {
        log.Println("Token mismatch for user ID:", id)
        return errors.New("invalid token or mismatched session")
    }

    user.JwtToken = ""
    if err := u.userRepo.UpdateJwtToken(user); err != nil {
        log.Println("Failed to invalidate token for user ID:", id)
        return err
    }

    log.Println("User logged out successfully:", id)
    return nil
}






