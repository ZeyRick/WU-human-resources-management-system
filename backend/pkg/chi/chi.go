package chi

import (
	"backend/adapters/controllers"
	"backend/pkg/logger"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// No need to understand this function
// graceful shutdown mean shutdown smartly (sometime system have work to be done before shutting down)
func StartServerWithGracefulShutdown() {
	baseUrl := os.Getenv("BASE_URL")
	if baseUrl == "" {
		baseUrl = "0.0.0.0:3333"
	}
	// The HTTP Server
	server := &http.Server{Addr: baseUrl, Handler: service()}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for system signals for process to interrupt/quit (ctrl + c on terminal to shutdown the server)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with  period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger  shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	// Run the server
	logger.Success(fmt.Sprintf("We Have Started The Server On \033[0;34m http://%s \033[0m\n", baseUrl))
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}

// this function is called in the upper function to set up paths and controller for paths
func service() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	helloWorld := controllers.NewHelloWorldController()
	user := controllers.NewUserController()
	clock := controllers.NewClockController()
	r.Get("/", helloWorld.GetHelloWorld) // setting the path '/' handler or we can call controller to control the request sent into this path by front end
	r.Post("/register", user.UserRegister)
	r.Post("/login", user.UserLogin)
	r.Post("/clock", clock.Clock)
	return r
}
