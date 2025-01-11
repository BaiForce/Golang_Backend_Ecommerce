package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Import driver MySQL
)

var DB *sql.DB

// Connect menghubungkan ke database MySQL
func Connect() error {
	// Ganti dengan konfigurasi koneksi database Anda
	dsn := "root@tcp(127.0.0.1:3306)/shoppe_db?parseTime=true"
	var err error

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %v", err)
	}

	// Cek koneksi
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	log.Println("Database connected successfully!")
	return nil
}
