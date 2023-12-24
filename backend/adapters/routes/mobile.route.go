package routes

import (
	"backend/adapters/controllers"

	"github.com/go-chi/chi"
)

func initAppRoutes(r chi.Router) {

	clock := controllers.NewClockController()

	r.Route("/app", func(r chi.Router) {
		// private route
		r.Group(func(r chi.Router) {
			// clock
			r.Post("/clock", clock.Clock)
		})
	})
}
