package user

import (
	"github.com/Corray333/dating/internal/domains/user/storage"
	"github.com/Corray333/dating/internal/domains/user/transport"
	"github.com/Corray333/dating/internal/domains/user/types"
	"github.com/Corray333/dating/pkg/server/auth"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type Storage interface {
	InsertUser(user types.User) (int, string, error)
	LoginUser(user types.User) (int, string, error)
	CheckAndUpdateRefresh(id int, refresh string) (string, error)
	SelectUser(id int) (types.User, error)
}

func Init(db *sqlx.DB, router *chi.Mux) error {
	store := storage.NewStorage(db)

	router.Post("/users/login", transport.LogIn(store))
	router.Post("/users/signup", transport.SignUp(store))
	router.Get("/users/refresh", transport.RefreshAccessToken(store))
	router.With(auth.NewMiddleware()).Put("/users/{id}", transport.UpdateUser(store))
	router.With(auth.NewMiddleware()).Get("/users/{id}", transport.GetUser(store))
	return nil
}
