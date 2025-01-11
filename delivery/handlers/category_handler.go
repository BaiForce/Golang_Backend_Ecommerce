package handlers

import (
	"fmt"
	"net/http"

	"github.com/Backend_Shoppe/models"
	"github.com/Backend_Shoppe/usecase"
	"github.com/gofiber/fiber/v2"
)

// CategoryHandler - Struktur handler untuk kategori
type CategoryHandler struct {
	categoryUsecase usecase.CategoryUsecase
}

// NewCategoryHandler - Fungsi untuk inisialisasi handler kategori
func NewCategoryHandler(u usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{categoryUsecase: u}
}

// CreateCategory - Handler untuk membuat kategori baru
func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	req := new(models.Category)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Validasi input
	if req.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Category name is required"})
	}

	category, err := h.categoryUsecase.CreateCategory(req)
	if err != nil {
		// Tambahkan log error
		fmt.Printf("CreateCategory handler error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create category"})
	}

	return c.Status(fiber.StatusCreated).JSON(category)
}

// GetCategories - Handler untuk mendapatkan semua kategori
func (h *CategoryHandler) GetCategories(c *fiber.Ctx) error {
	// Panggil usecase untuk mendapatkan daftar kategori
	categories, err := h.categoryUsecase.GetCategories()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch categories",
		})
	}

	// Kembalikan response sukses dengan data kategori
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": categories,
	})
}

// UpdateCategory - Handler untuk memperbarui kategori berdasarkan ID
func (h *CategoryHandler) UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Category ID is required",
		})
	}

	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to parse request body",
		})
	}

	// Validasi input
	if category.Name == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input, category name must be provided.",
		})
	}

	// Panggil usecase untuk memperbarui kategori
	updatedCategory, err := h.categoryUsecase.UpdateCategory(id, &category)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update category",
		})
	}

	// Kembalikan response sukses
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Category updated successfully",
		"data":    updatedCategory,
	})
}

// DeleteCategory - Handler untuk menghapus kategori berdasarkan ID
func (h *CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Category ID is required",
		})
	}

	// Panggil usecase untuk menghapus kategori
	err := h.categoryUsecase.DeleteCategory(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete category",
		})
	}

	// Kembalikan response sukses
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Category deleted successfully",
	})
}
