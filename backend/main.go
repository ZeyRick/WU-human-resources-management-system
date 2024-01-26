package main

import (
	"backend/adapters/routes"
	"backend/pkg/bot"
	"backend/pkg/chi"
	"backend/pkg/db"
	"backend/pkg/logger"

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
	bot.NewBot().TelegramBot()

	r := routes.InitRoutes()
	// start the api server
	chi.StartServerWithGracefulShutdown(r)
}
