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
	department := controllers.NewDepartmentController()
	report := controllers.NewReportController()
	employee_request := controllers.NewEmployeeRequestController()

	r.Route("/admin", func(r chi.Router) {
		r.Post("/user/login", user.UserLogin)

		r.Group(func(r chi.Router) {
			//r.Use(middlewares.LoginMiddleware)
			// for testing
			r.Get("/hello", helloWorld.GetHelloWorld)

			// User
			r.Post("/user", user.UserRegister)
			r.Get("/user", user.GetUserData)
			r.Get("/user/userInfo", user.GetUserInfo)
			r.Delete("/user/{userId}", user.Delete)
			r.Patch("/user/{userId}", user.ResetPW)

			// Clock
			r.Get("/clock", clock.List)

			//report
			r.Get("/report", report.List)

			// Employe
			r.Post("/employee", employee.Add)
			r.Patch("/employee/{employeeId}", employee.Edit)
			r.Get("/employee", employee.List)
			r.Get("/employee/all", employee.All)
			r.Delete("/employee/{employeeId}", employee.Delete)

			// Employee Request
			r.Get("/employee_request", employee_request.List)
			r.Get("/employee_request/confirmation", employee_request.Confirmation)

			// Schedule
			r.Post("/schedule", schedule.Add)
			r.Get("/schedule", schedule.GetAllWithFormat)
			r.Get("/schedule/{employeeId}", schedule.GetByEmployeeId)
			r.Patch("/schedule", schedule.Update)

			// Department
			r.Get("/department/all", department.All)
			r.Get("/department", department.List)
			r.Post("/department", department.Add)
			r.Patch("/department/{departmentId}", department.Edit)
		})
	})
}
