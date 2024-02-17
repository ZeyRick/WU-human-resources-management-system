package main

import (
	"backend/adapters/routes"
	"backend/pkg/chi"
	"backend/pkg/db"
	"backend/pkg/logger"
	"backend/pkg/telegrambot"

	"github.com/joho/godotenv"
)

func main() {
	//loading env file
	godotenv.Load()

	//Init logger config
	logger.InitLogger()

	//init database
	db.InitDatabase()

	//Start Teleggram bot api
	telegrambot.NewBot().TelegramBot()

	r := routes.InitRoutes()
	// start the api server
	chi.StartServerWithGracefulShutdown(r)
}
