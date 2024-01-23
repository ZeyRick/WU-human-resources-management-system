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
	department := controllers.NewDepartmentController()



	r.Route("/admin", func(r chi.Router) {
		r.Post("/user/login", user.UserLogin)
		r.Get("/user/logout", user.UserLogout)

		

		r.Group(func(r chi.Router) {

			r.Use(LoginMiddleware)
			// for testing
			r.Get("/hello", helloWorld.GetHelloWorld)

			// User
			r.Post("/user", user.UserRegister)
			r.Get("/user", user.GetUserData)

			// Clock
			r.Get("/clock", clock.List)

			// Employe
			r.Post("/employee", employee.Add)
			r.Get("/employee", employee.List)
			r.Get("/employee/all", employee.All)

			// Schedule
			r.Post("/schedule", schedule.Add)
			r.Get("/schedule", schedule.GetAllWithFormat)
			r.Get("/schedule/{employeeId}", schedule.GetByEmployeeId)
			r.Patch("/schedule", schedule.Update)

			//department
			r.Get("/department/all", department.All)
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
		https.ResponseMsg(w, r, http.StatusUnauthorized, "Unauthorized")
	})
}
