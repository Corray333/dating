package storage

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	DB    *sqlx.DB
	Redis *redis.Client
}

func Connect() (*Storage, error) {
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
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	if err := redisClient.Ping(ctx).Err(); err != nil {
		slog.Error("Failed to connect to Redis")
		return nil, err
	}

	return &Storage{db, redisClient}, err
}
