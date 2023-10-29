package main

import (
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

	// start the api server
	chi.StartServerWithGracefulShutdown()

}
