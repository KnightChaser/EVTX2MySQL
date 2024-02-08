package main

import (
	"EVTX2MySQL/database"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	// Load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Replace the connection parameters with your MySQL server details
	db, err := database.ObtainMySQLConnection()
	if err != nil {
		log.Panicf("Failed to connect to MySQL server: %v", err)
	}
	defer db.Close()

	tableName := "evtx"

	// Initialize the MySQL database
	database.CreateMySQLDatabase()                                        // Create the database
	database.CreateMySQLTable("./database/evtxTableModel.sql", tableName) // Create the table named as var tableName

	// Migrate the EVTX file to MySQL to the table named as var tableName
	database.MigrateEVTX2MySQL("D:\\sampleEVTX.evtx", db, tableName)

	fmt.Println("::: Migration finished")
}
