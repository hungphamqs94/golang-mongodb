package routers

import (
	controller "mongodb/controller"

	"github.com/gorilla/mux"
)

func RouteEmployee(myRouter *mux.Router) {
	myRouter.HandleFunc("/employees", controller.GetEmployees)
	myRouter.HandleFunc("/employee/{id}", controller.SingleEmployee)
	myRouter.HandleFunc("/employee", controller.CreateEmployee).Methods("POST")
	myRouter.HandleFunc("/employee", controller.UpdateEmployee).Methods("PUT")
	myRouter.HandleFunc("/employee", controller.DeleteEmployee).Methods("DELETE")
}
