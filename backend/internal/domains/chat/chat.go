package chat

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func Init(db *sqlx.DB, router *chi.Mux) error {
	// store := storage.NewChatStorage(db)

	return nil
}
