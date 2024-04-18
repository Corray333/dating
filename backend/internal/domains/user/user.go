package user

import (
	user_storage "github.com/Corray333/dating/internal/domains/user/storage"
	"github.com/Corray333/dating/internal/domains/user/transport"
	"github.com/Corray333/dating/internal/domains/user/types"
	"github.com/Corray333/dating/internal/storage"
	"github.com/Corray333/dating/pkg/server/auth"
	"github.com/go-chi/chi/v5"
)

type Storage interface {
	InsertUser(user types.User, agent string) (int, string, error)
	LoginUser(user types.User, agent string) (int, string, error)
	CheckAndUpdateRefresh(id int, refresh string) (string, error)
	SelectUser(id string) (types.User, error)
	UpdateUser(user types.User) error
}

func Init(storage *storage.Storage, router *chi.Mux) error {
	// TODO: fix this
	temp := user_storage.UserStorage(*storage)
	store := &temp

	router.Post("/users/login", transport.LogIn(store))
	router.Post("/users/signup", transport.SignUp(store))
	router.Get("/users/refresh", transport.RefreshAccessToken(store))
	router.Get("/users/verify/email", transport.VerifyEmail(store))
	router.With(auth.NewMiddleware()).Put("/users/{id}", transport.UpdateUser(store))
	router.With(auth.NewMiddleware()).Get("/users/{id}", transport.GetUser(store))
	return nil
}
