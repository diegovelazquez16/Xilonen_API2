package core

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq" // Driver para PostgreSQL
)

// GetDB retorna una conexión a PostgreSQL
func GetDB() (*sql.DB, error) {
	dbURL := os.Getenv("POSTGRES_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("❌ Error al conectar a PostgreSQL: %v", err)
		return nil, err
	}
	return db, nil
}
