package database

import (
	"log"
)

func Migrate() {
	// Optionally, drop the tables first to ensure a fresh start
	dropUserTable := "DROP TABLE IF EXISTS users;"
	_, err := DB.Exec(dropUserTable)
	if err != nil {
		log.Fatalf("Error dropping users table: %v", err)
	}

	dropProductTable := "DROP TABLE IF EXISTS products;"
	_, err = DB.Exec(dropProductTable)
	if err != nil {
		log.Fatalf("Error dropping products table: %v", err)
	}

	dropCategoryTable := "DROP TABLE IF EXISTS categories;"
	_, err = DB.Exec(dropCategoryTable)
	if err != nil {
		log.Fatalf("Error dropping categories table: %v", err)
	}

	// Create the users table
	createUserTable := `
		CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		jwt_token TEXT,
		role VARCHAR(50) NOT NULL 
	);
	`
	_, err = DB.Exec(createUserTable)
	if err != nil {
		log.Fatalf("Error creating users table: %v", err)
	}

	// Create the categories table
	createCategoryTable := `
		CREATE TABLE IF NOT EXISTS categories (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT
	);
	`
	_, err = DB.Exec(createCategoryTable)
	if err != nil {
		log.Fatalf("Error creating categories table: %v", err)
	}

	// Create the products table with category_id as a foreign key
	createProductTable := `
		CREATE TABLE IF NOT EXISTS products (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		price FLOAT NOT NULL,
		category_id INT,
		FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL
	);
	`
	_, err = DB.Exec(createProductTable)
	if err != nil {
		log.Fatalf("Error creating products table: %v", err)
	}
}
