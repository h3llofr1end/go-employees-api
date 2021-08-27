package controllers

import "github.com/h3llofr1end/go-employees-api/api/middlewares"

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	s.Router.HandleFunc("/departments", middlewares.SetMiddlewareJSON(s.CreateDepartment)).Methods("POST")
	s.Router.HandleFunc("/departments", middlewares.SetMiddlewareJSON(s.GetDepartments)).Methods("GET")
	s.Router.HandleFunc("/departments/{id}", middlewares.SetMiddlewareJSON(s.GetDepartment)).Methods("GET")
	s.Router.HandleFunc("/departments/{id}", middlewares.SetMiddlewareJSON(s.UpdateDepartment)).Methods("PUT")
	s.Router.HandleFunc("/departments/{id}", middlewares.SetMiddlewareJSON(s.DeleteDepartment)).Methods("DELETE")

	s.Router.HandleFunc("/employees", middlewares.SetMiddlewareJSON(s.CreateEmployee)).Methods("POST")
	s.Router.HandleFunc("/employees", middlewares.SetMiddlewareJSON(s.GetEmployees)).Methods("GET")
	s.Router.HandleFunc("/employees/{id}", middlewares.SetMiddlewareJSON(s.GetEmployee)).Methods("GET")
	s.Router.HandleFunc("/employees/{id}", middlewares.SetMiddlewareJSON(s.UpdateEmployee)).Methods("PUT")
	s.Router.HandleFunc("/employees/{id}", middlewares.SetMiddlewareJSON(s.DeleteEmployee)).Methods("DELETE")
}
