// Handler with Logging
package handlers

import (
	"errors"
	"log"

	"github.com/Backend_Shoppe/models"
	"github.com/Backend_Shoppe/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

	type UserHandler struct {
		userUsecase usecase.UserUsecase
	}

	func NewUserHandler(u usecase.UserUsecase) *UserHandler {
		return &UserHandler{userUsecase: u}
	}

	func parseToken(tokenStr, secret string) (jwt.MapClaims, error) {
		log.Println("Parsing token")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Println("Unexpected signing method")
				return nil, errors.New("unexpected signing method")
			}
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			log.Printf("Invalid or expired token: %v\n", err)
			return nil, errors.New("invalid or expired token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Println("Invalid token claims")
			return nil, errors.New("invalid token claims")
		}

		log.Println("Token parsed successfully", claims)
		return claims, nil
	}

	func (h *UserHandler) Login(c *fiber.Ctx) error {
		log.Println("Processing login request")

		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&req); err != nil {
			log.Printf("Failed to parse request body: %v\n", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		token, err := h.userUsecase.Login(req.Email, req.Password)
		if err != nil {
			log.Printf("Login failed for email %s: %v\n", req.Email, err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
		}

		user, err := h.userUsecase.GetUserByEmail(req.Email)
		if err != nil {
			log.Printf("Failed to fetch user details for email %s: %v\n", req.Email, err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user details"})
		}

		user.JwtToken = token
		log.Printf("User %s logged in successfully\n", req.Email)
		return c.JSON(fiber.Map{
			"message": "User Login successfully",
			"user":    user,
		})
	}

	func (h *UserHandler) Register(c *fiber.Ctx) error {
		log.Println("Processing registration request")

		req := new(models.User)
		if err := c.BodyParser(req); err != nil {
			log.Printf("Failed to parse request body: %v\n", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		user, err := h.userUsecase.Register(req)
		if err != nil {
			log.Printf("Registration failed for email %s: %v\n", req.Email, err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		log.Printf("User %s registered successfully\n", req.Email)
		return c.JSON(fiber.Map{
			"message": "User registered successfully",
			"user":    user,
		})
	}

	func (h *UserHandler) Logout(c *fiber.Ctx) error {
    log.Println("Processing logout request")

    // Extract the token from the Authorization header
    token := c.Get("Authorization")
    if len(token) <= 7 || token[:7] != "Bearer " {
        log.Println("Invalid token format")
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
    }
    token = token[7:]

    // Parse the token
    claims, err := parseToken(token, "secret")
    if err != nil {
        log.Printf("Failed to parse token: %v\n", err)
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
    }

    // Extract the user ID from the token claims
    userID, ok := claims["id"].(float64) // JWT claims "id" will be of type float64
    if !ok {
        log.Println("User ID not found in token claims:", claims)
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
    }

    // Call the usecase to logout the user by ID
    err = h.userUsecase.Logout(int(userID), token)
    if err != nil {
        log.Printf("Failed to logout user with ID %d: %v\n", int(userID), err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    // Fetch user details after logout
    user, err := h.userUsecase.GetUserByEmail(claims["email"].(string))
    if err != nil {
        log.Printf("Failed to fetch user details for email %s: %v\n", claims["email"], err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user details"})
    }

    log.Printf("User %d logged out successfully\n", int(userID))
    return c.JSON(fiber.Map{
        "message": "User logged out successfully",
        "user": fiber.Map{
            "id":       user.ID,
            "name":     user.Name,
            "password": user.Password, // Include password here (be careful with exposing passwords)
        },
    })
}

