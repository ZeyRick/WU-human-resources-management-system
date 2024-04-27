package routes

import (
	"backend/adapters/controllers"
	"backend/adapters/middlewares"

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
	employeeRequest := controllers.NewEmployeeRequestController()
	clockSetting := controllers.NewClockSettingController()

	r.Route("/admin", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(middlewares.LoginMiddleware)
			r.Post("/employee/uploadFiles", employee.UploadFiles)
		})

		r.Group(func(r chi.Router) {
			r.Use(middlewares.DecryptMiddleware)
			r.Post("/user/login", user.UserLogin)

			// export
			r.Get("/clock/attendence/export", clock.AttendenceExport)
			r.Get("/clock/report/export", report.ExportList)

			r.Group(func(r chi.Router) {
				r.Use(middlewares.LoginMiddleware)

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
				r.Get("/clock/attendence", clock.Attendence)

				// Clock Setting
				r.Get("/clock-setting", clockSetting.Get)
				r.Post("/clock-setting", clockSetting.Save)

				//report
				r.Get("/report", report.List)

				// Employe
				r.Post("/employee", employee.Add)
				r.Patch("/employee/{employeeId}", employee.Edit)
				r.Get("/employee", employee.List)
				r.Get("/employee/all", employee.All)
				r.Delete("/employee/{employeeId}", employee.Delete)

				// Employee Request
				r.Get("/employee_request", employeeRequest.List)
				r.Post("/employee_request/confirmation", employeeRequest.Confirmation)

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

	})
}
