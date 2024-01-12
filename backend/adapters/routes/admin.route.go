package routes

import (
	"backend/adapters/controllers"

	"github.com/go-chi/chi"
)

func initAdminRoutes(r chi.Router) {
	helloWorld := controllers.NewHelloWorldController()
	user := controllers.NewUserController()
	employee := controllers.NewEmployeeController()
	clock := controllers.NewClockController()
	schedule := controllers.NewScheduleController()

	r.Route("/admin", func(r chi.Router) {
		// private route
		r.Group(func(r chi.Router) {
			// for testing
			r.Get("/", helloWorld.GetHelloWorld)

			// User
			r.Post("/user/register", user.UserRegister)
			r.Post("/user/login", user.UserLogin)

			// Clock
			r.Get("/clock", clock.List)

			// Employe
			r.Post("/employee", employee.Add)
			r.Get("/employee", employee.List)

			// Schedule
			r.Post("/schedule", schedule.Add)
			r.Get("/schedule", schedule.List)
		})
	})
}
