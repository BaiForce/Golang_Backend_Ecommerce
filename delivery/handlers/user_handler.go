package handlers

import (
	"github.com/Backend_Shoppe/models"
	"github.com/Backend_Shoppe/usecase"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: u}
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	req := new(models.User)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	// Melakukan login
	token, err := h.userUsecase.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	// Mengembalikan token JWT setelah login berhasil
	return c.JSON(fiber.Map{"token": token})
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	req := new(models.User)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	// Melakukan registrasi
	user, err := h.userUsecase.Register(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Mengembalikan data user setelah registrasi berhasil
	return c.JSON(fiber.Map{
		"message": "User registered successfully",
		"user":    user,
	})
}
