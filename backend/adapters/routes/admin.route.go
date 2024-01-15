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
	r.Route("/admin", func(r chi.Router) {
		// private route
		r.Use(LoginMiddleware)

		r.Group(func(r chi.Router) {

			// for testing
			r.Get("/hello", helloWorld.GetHelloWorld)

			// User
			r.Post("/user/register", user.UserRegister)
			r.Post("/user/login", user.UserLogin)
			r.Get("/user/userdata", user.GetUserData)

			// Employee
			r.Post("/employee/add", employee.Add)
			r.Get("/employee/list", employee.List)
		})
	})
}

func LoginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ok, userId := jwttoken.CheckCookie(w, r, "LoginCookie", 24*30)
		if ok {
			https.ResponseJSON(w, r, http.StatusOK, userId)
			next.ServeHTTP(w, r)
			return
		}
		if r.URL.Path == "/admin/user/login" || r.URL.Path == "/admin/user/register" || r.URL.Path == "/admin/test" {
			next.ServeHTTP(w, r)
			return
		}
		https.ResponseMsg(w, r, http.StatusUnauthorized, "Missing Cookie")
		return
	})
}
