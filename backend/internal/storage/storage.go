package storage

import (
	"log"
	"log/slog"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", os.Getenv("DB_CONNSTR"))
	if err != nil {
		log.Fatal(err)
	}

	// Test the connection to the database
	if err := db.Ping(); err != nil {
		return nil, err
	} else {
		slog.Info("Successfully Connected")
	}

	return db, nil
}
