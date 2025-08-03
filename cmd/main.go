package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Rothiii/golang-marketplace/cmd/api"
	"github.com/Rothiii/golang-marketplace/config"
	"github.com/Rothiii/golang-marketplace/db"
	_ "github.com/lib/pq"
)

func main() {
    // PostgreSQL connection string format
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        "localhost", // or use getEnv("DB_HOST", "localhost") 
        config.Envs.DBPort,
        config.Envs.DBUser,
        config.Envs.DBPassword,
        config.Envs.DBName,
    )
    
    db, err := db.NewPostgreSQLDatabase(connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    initStorage(db)
    
    server := api.NewAPIServer(":8080", nil)
    if err := server.Run(); err != nil {
        log.Fatal(err)
    }
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	log.Println("Database connection established successfully")
}

