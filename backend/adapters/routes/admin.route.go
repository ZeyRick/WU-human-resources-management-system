package routes

import (
	"backend/adapters/controllers"
	"backend/pkg/https"
	"backend/pkg/jwttoken"
	"net/http"

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

		r.Post("/user/login", user.UserLogin)

		r.Group(func(r chi.Router) {

			r.Use(LoginMiddleware)
			// for testing
			r.Get("/hello", helloWorld.GetHelloWorld)

			// User
			r.Post("/user/register", user.UserRegister)
			r.Get("/user/userdata", user.GetUserData)

			// Clock
			r.Get("/clock", clock.List)

			// Employe
			r.Post("/employee", employee.Add)
			r.Get("/employee", employee.List)
			r.Get("/employee/all", employee.All)

			// Schedule
			r.Post("/schedule", schedule.Add)
			r.Get("/schedule", schedule.List)
		})
	})
}

func LoginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ok, _ := jwttoken.CheckCookie(w, r, "LoginCookie", 24*30)
		if ok {
			next.ServeHTTP(w, r)
			return
		}
		https.ResponseMsg(w, r, http.StatusUnauthorized, "Missing Cookie")
		return
	})
}
