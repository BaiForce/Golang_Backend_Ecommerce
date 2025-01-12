package middleware

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var JwtSecret = []byte("secret") // Gantilah dengan secret key Anda

// JwtAuth adalah middleware untuk memvalidasi token JWT
func JwtAuth(c *fiber.Ctx) error {
	// Periksa apakah route membutuhkan otentikasi
	if !isProtectedRoute(c.Path()) {
		return c.Next()
	}

	// Ambil token dari header Authorization
	token := c.Get("Authorization")
	if len(token) <= 7 || token[:7] != "Bearer " {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
	}
	token = token[7:]

	// Validasi token
	claims, err := parseToken(token, string(JwtSecret))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	// Simpan klaim di context untuk digunakan oleh handler
	c.Locals("claims", claims)

	return c.Next()
}

// parseToken memvalidasi token JWT dan mengembalikan klaimnya
func parseToken(tokenStr, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

// isProtectedRoute menentukan apakah rute membutuhkan otentikasi
func isProtectedRoute(path string) bool {
	protectedRoutes := []string{
		"/products",
		"/products/",
		"/categories",
		"/categories/",
		"/logout",
	}

	for _, route := range protectedRoutes {
		if strings.HasPrefix(path, route) {
			return true
		}
	}
	return false
}
