package middleware

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var JwtSecret = []byte("secret") // Gantilah dengan secret key Anda

// Middleware untuk memeriksa JWT
func JwtAuth(c *fiber.Ctx) error {
	// Pastikan middleware hanya dipanggil untuk route yang membutuhkan otentikasi
	if !isProtectedRoute(c.Path()) {
		return c.Next() // Langsung lanjutkan jika route tidak dilindungi
	}

	// Ambil token dari header Authorization
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "No token provided"})
	}

	// Potong prefix "Bearer "
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Parse dan verifikasi token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return JwtSecret, nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	// Token valid, lanjutkan ke handler berikutnya
	return c.Next()
}

// Fungsi untuk menentukan apakah route memerlukan otentikasi
func isProtectedRoute(path string) bool {
	// Daftar route yang membutuhkan JWT auth
	protectedRoutes := []string{
		//product
		"/products",     
		"/products/*", 
		//category  
		"/categories",     
		"/categories/*",   
	}

	for _, route := range protectedRoutes {
		if strings.HasPrefix(path, route) {
			return true
		}
	}
	return false
}
