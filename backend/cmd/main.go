package main

import (
	"os"

	_ "github.com/Corray333/dating/docs"
	"github.com/Corray333/dating/internal/app"
	"github.com/Corray333/dating/internal/config"
)

func main() {
	config.Configure(os.Args[1])
	app := app.NewApp()
	app.Run()
}
