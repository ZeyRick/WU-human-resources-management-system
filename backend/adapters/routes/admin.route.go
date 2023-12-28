package routes

import (
	"backend/adapters/controllers"

	"github.com/go-chi/chi"
)

func initAdminRoutes(r chi.Router) {
	helloWorld := controllers.NewHelloWorldController()
	user := controllers.NewUserController()
    employee := controllers.NewEmployeeController()

	r.Route("/admin", func(r chi.Router) {
		// private route
		r.Group(func(r chi.Router) {
			// for testing
			r.Get("/", helloWorld.GetHelloWorld)

			// User
			r.Post("/user/register", user.UserRegister)
			r.Post("/user/login", user.UserLogin)

            // Employee
            r.Post("/employee/add", employee.Add)
            r.Get("/employee/list", employee.List)
		})
	})
}
