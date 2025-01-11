package handlers

import (
	"fmt"
	"net/http"

	"github.com/Backend_Shoppe/models"
	"github.com/Backend_Shoppe/usecase"
	"github.com/gofiber/fiber/v2"
)

// ProductHandler - Struktur handler untuk produk
type ProductHandler struct {
	productUsecase usecase.ProductUsecase
}

// NewProductHandler - Fungsi untuk inisialisasi handler produk
func NewProductHandler(u usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{productUsecase: u}
}

// CreateProduct - Handler untuk membuat produk baru
func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	req := new(models.Product)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Validasi input
	if req.Name == "" || req.Price <= 0 || req.CategoryID <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input data"})
	}


	product, err := h.productUsecase.CreateProduct(req)
	if err != nil {
		// Tambahkan log error
		fmt.Printf("CreateProduct handler error: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create product"})
	}

	return c.Status(fiber.StatusCreated).JSON(product)
}


// GetProducts - Handler untuk mendapatkan semua produk
func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
	// Panggil usecase untuk mendapatkan daftar produk
	products, err := h.productUsecase.GetProducts()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch products",
		})
	}

	// Kembalikan response sukses dengan data produk
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": products,
	})
}

// UpdateProduct - Handler untuk memperbarui produk berdasarkan ID
func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Product ID is required",
		})
	}

	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to parse request body",
		})
	}

	// Validasi input
	if product.Name == "" || product.Price <= 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input, product name and price must be provided.",
		})
	}

	// Panggil usecase untuk memperbarui produk
	updatedProduct, err := h.productUsecase.UpdateProduct(id, &product)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update product",
		})
	}

	// Kembalikan response sukses
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Product updated successfully",
		"data":    updatedProduct,
	})
}

// DeleteProduct - Handler untuk menghapus produk berdasarkan ID
func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Product ID is required",
		})
	}

	// Panggil usecase untuk menghapus produk
	err := h.productUsecase.DeleteProduct(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete product",
		})
	}

	// Kembalikan response sukses
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Product deleted successfully",
	})
}
