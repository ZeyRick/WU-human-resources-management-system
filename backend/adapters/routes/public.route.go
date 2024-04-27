package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

func initPublicRoutes(r chi.Router) {
	r.Route("/public", func(r chi.Router) {
		// private route
		fs := http.FileServer(http.Dir("./public/"))
		r.Handle("/images/employee/*", http.StripPrefix("/public", fs))

	})
}
