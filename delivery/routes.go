package delivery

import (
	"github.com/Backend_Shoppe/database"
	"github.com/Backend_Shoppe/delivery/handlers"
	"github.com/Backend_Shoppe/middleware"
	"github.com/Backend_Shoppe/repository"
	"github.com/Backend_Shoppe/usecase"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Inisialisasi repository, usecase, dan handler untuk User
	userRepo := repository.NewUserRepository(database.DB)
	userUsecase := usecase.NewUserUsecase(userRepo, "secret")
	userHandler := handlers.NewUserHandler(userUsecase)

	// Inisialisasi repository, usecase, dan handler untuk Product
	productRepo := repository.NewProductRepository(database.DB)
	productUsecase := usecase.NewProductUsecase(productRepo)
	productHandler := handlers.NewProductHandler(productUsecase)

	// Inisialisasi repository, usecase, dan handler untuk Category
	categoryRepo := repository.NewCategoryRepository(database.DB)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryUsecase)

	// Route Login dan Register untuk User
	app.Post("/login", userHandler.Login)
	app.Post("/register", userHandler.Register)
	app.Post("/logout", middleware.JwtAuth, userHandler.Logout)

	// Group routes untuk produk
	productRoutes := app.Group("/products", middleware.JwtAuth)
	productRoutes.Post("/", productHandler.CreateProduct)
	productRoutes.Get("/", productHandler.GetProducts)
	productRoutes.Put("/:id", productHandler.UpdateProduct)
	productRoutes.Delete("/:id", productHandler.DeleteProduct)

	// Group routes untuk category
	categoryRoutes := app.Group("/categories", middleware.JwtAuth)
	categoryRoutes.Post("/", categoryHandler.CreateCategory)
	categoryRoutes.Get("/", categoryHandler.GetCategories)
	categoryRoutes.Put("/:id", categoryHandler.UpdateCategory)
	categoryRoutes.Delete("/:id", categoryHandler.DeleteCategory)

	// logoutRoutes := app.Group("/logout", middleware.JwtAuth)
	// logoutRoutes.Post("/logout", userHandler.Logout)

}
