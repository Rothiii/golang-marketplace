package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// You'll need to add this function to your db package
func NewPostgreSQLDatabase(connStr string) (*sql.DB, error) {
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }
    
    if err := db.Ping(); err != nil {
        return nil, err
    }
    
    return db, nil
}