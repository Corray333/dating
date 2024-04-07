package app

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/Corray333/dating/internal/config"
	"github.com/Corray333/dating/internal/domains/user"
	"github.com/Corray333/dating/internal/storage"
	"github.com/Corray333/dating/pkg/server/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	db     *sqlx.DB
	server *http.Server
}

//	@title			Dating API
//	@version		1.0
//	@description	This is a dating API
//	@host			localhost:3001
func NewApp() *App {
	db, err := storage.Connect()
	if err != nil {
		slog.Error("Failed to connect to the database: " + err.Error())
		panic(err)
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: viper.GetStringSlice("http_server.allowed_origins"),
		AllowedMethods: viper.GetStringSlice("http_server.allowed_methods"),
		AllowedHeaders: []string{"Authorization"},
		MaxAge:         300,
	}))

	if viper.GetString("env") == config.EnvDev {
		router.Use(middleware.RequestID)
		router.Use(logger.New(slog.Default()))
	}

	if err := user.Init(db, router); err != nil {
		slog.Error(err.Error())
		panic(err)

	}

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3001/swagger/doc.json"), //The url pointing to API definition
	))

	router.Get("/test", testFunc)

	return &App{
		db: db,
		server: &http.Server{
			Addr:    os.Getenv("APP_IP") + ":" + os.Getenv("APP_PORT"),
			Handler: router,
		},
	}
}

func (app *App) Run() {
	slog.Info("Server started on port " + os.Getenv("APP_PORT"))
	if err := app.server.ListenAndServe(); err != nil {
		slog.Error("Server failed to start: " + err.Error())
		panic(err)
	}
}

//	@Summary		Test function
//	@Description	Test function
//	@ID				test
//	@Accept			json
//	@Produce		json
func testFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
