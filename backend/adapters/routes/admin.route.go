package routes

import (
	"backend/adapters/controllers"
	"backend/pkg/https"
	"backend/pkg/jwttoken"
	"fmt"
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
			r.Get("/", helloWorld.GetHelloWorld)

			// User
			r.Post("/user/register", user.UserRegister)
			r.Post("/user/login", user.UserLogin)
			r.Post("/user/getusers", user.GetUserData)

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
			next.ServeHTTP(w, r)
			https.ResponseMsg(w, r, http.StatusOK, fmt.Sprintf("Logged in. User ID = ", userId))
			return
		}
		https.ResponseMsg(w, r, http.StatusUnauthorized, "Cookie not found")
		if r.URL.Path == "/user/login" || r.URL.Path == "/user/register" {
			next.ServeHTTP(w, r)
		}
		return
	})
}
