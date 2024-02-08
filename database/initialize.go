package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func ObtainMySQLConnection() (*sql.DB, error) {
	mysqlConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_ACCESS_HOST"),
		os.Getenv("MYSQL_ACCESS_PORT"))

	db, err := sql.Open("mysql", mysqlConnectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// CreateMySQLDatabase creates a database in the MySQL server
func CreateMySQLDatabase() {
	db, err := ObtainMySQLConnection()

	if err != nil {
		log.Panic("Failed to connect to MySQL server:", err)
		return
	}
	defer db.Close()

	// Create the database
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + os.Getenv("MYSQL_DATABASE_NAME"))
}

// CreateMySQLTable creates a table in the MySQL database
func CreateMySQLTable(tableSchemaSQLFilePath string, tableName string) {
	db, err := ObtainMySQLConnection()
	if err != nil {
		log.Panic("Failed to connect to MySQL server:", err)
		return
	}
	defer db.Close()

	// Create the table
	schema, err := os.ReadFile(tableSchemaSQLFilePath)
	if err != nil {
		fmt.Println("Failed to read schema file:", err)
		return
	}

	// Substitute the database and table name in the schema
	schemaString := string(schema)
	schemaString = strings.Replace(schemaString, "$mysql_database_name", os.Getenv("MYSQL_DATABASE_NAME"), 1) // A single MySQL server can have multiple databases, so we need to specify the database name
	schemaString = strings.Replace(schemaString, "$mysql_table_name", tableName, 1)                           // A single database can have multiple tables, so we need to specify the table name

	_, err = db.Exec(schemaString)
	if err != nil {
		fmt.Println("Failed to create table:", err)
		return
	}
}
