package usecase

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/Backend_Shoppe/models"
	"github.com/Backend_Shoppe/repository"
	"github.com/golang-jwt/jwt/v4"
)

// Fungsi untuk meng-hash password yang dimasukkan saat login
func hashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

type UserUsecase interface {
	Login(email, password string) (string, error)
	Register(user *models.User) (*models.User, error) // Mengubah untuk mengembalikan User
}

type userUsecase struct {
	userRepo repository.UserRepository
	secret   string
}

func NewUserUsecase(repo repository.UserRepository, secret string) UserUsecase {
	return &userUsecase{userRepo: repo, secret: secret}
}

func (u *userUsecase) Login(email, password string) (string, error) {
    // Cek apakah pengguna ada di database
    user, err := u.userRepo.GetUserByEmail(email)
    if err != nil || user == nil {
        return "", errors.New("invalid credentials")
    }

    // Bandingkan password langsung dengan yang ada di database
    if password != user.Password {
        return "", errors.New("invalid credentials")
    }

    // Generate token JWT yang baru
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id":  user.ID,
        "exp": time.Now().Add(time.Hour * 72).Unix(),
    })

    // Sign token JWT
    signedToken, err := token.SignedString([]byte(u.secret))
    if err != nil {
        return "", err
    }

    // Perbarui token JWT di database
    user.JwtToken = signedToken
    err = u.userRepo.UpdateJwtToken(user) // Pastikan ada method ini di repository
    if err != nil {
        return "", err
    }

    // Kembalikan token yang baru
    return signedToken, nil
}


func (u *userUsecase) Register(user *models.User) (*models.User, error) {
	// Hash password sebelum menyimpan ke database
	hashedPassword := hashPassword(user.Password)
	user.Password = hashedPassword

	// Simpan user ke database
	err := u.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	// Generate token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	signedToken, err := token.SignedString([]byte(u.secret))
	if err != nil {
		return nil, err
	}

	user.JwtToken = signedToken
	err = u.userRepo.UpdateJwtToken(user)
	if err != nil {
		return nil, err
	}

	// Kembalikan user dengan password yang di-hash
	return user, nil
}




