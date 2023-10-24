package main

import (
	"backend/pkg/chi"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World")
	godotenv.Load()

	chi.StartServerWithGracefulShutdown()

}