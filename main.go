package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Replace the connection parameters with your MySQL server details
	mysqlConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_ACCESS_HOST"),
		os.Getenv("MYSQL_ACCESS_PORT"),
		os.Getenv("MYSQL_DATABASE_NAME"))
	db, err := sql.Open("mysql", mysqlConnectionString)
	if err != nil {
		fmt.Println("Failed to connect to MySQL:", err)
		return
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to ping MySQL:", err)
		return
	}

	// Create table if not exists.
	// Schema is located at "./database/model.sql"
	schemaFile := "./database/model.sql"
	schema, err := os.ReadFile(schemaFile)
	if err != nil {
		fmt.Println("Failed to read schema file:", err)
		return
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		fmt.Println("Failed to create table:", err)
		return
	}

	fmt.Println("Connected to MySQL server!")
}
