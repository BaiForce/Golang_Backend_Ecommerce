package main

import (
	"fmt"
	"log"

	"github.com/Backend_Shoppe/database"
	"github.com/Backend_Shoppe/delivery"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Hubungkan ke database
	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// // Jalankan migrasi
	// database.Migrate()

	// Setup routes dari package delivery
	delivery.SetupRoutes(app)

	// Cetak semua route
	fmt.Println("Daftar Route:")
	for _, routes := range app.Stack() {
		for _, r := range routes {
			fmt.Printf("%s -> %s\n", r.Method, r.Path)
		}
	}

	// Jalankan server
	app.Listen(":3000")
}
