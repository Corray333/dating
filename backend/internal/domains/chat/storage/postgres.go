package storage

import "github.com/jmoiron/sqlx"

type ChatStorage struct {
	db *sqlx.DB
}

func NewChatStorage(db *sqlx.DB) *ChatStorage {
	return &ChatStorage{
		db: db,
	}
}
