package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/h3llofr1end/go-employees-api/api/models"
	"github.com/h3llofr1end/go-employees-api/api/responses"
)

func (server *Server) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	employee := models.Employee{}
	err = json.Unmarshal(body, &employee)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	employee.Prepare()
	err = employee.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	employeeCreated, err := employee.SaveEmployee(server.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, employeeCreated.ID))
	responses.JSON(w, http.StatusCreated, employeeCreated)
}

func (server *Server) GetEmployees(w http.ResponseWriter, r *http.Request) {
	employee := models.Employee{}
	employees, err := employee.FindAllEmployees(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, employees)
}

func (server *Server) GetEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	employee := models.Employee{}
	findedEmployee, err := employee.FindEmployeeByID(server.DB, eid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, findedEmployee)
}

func (server *Server) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	employee := models.Employee{}
	err = server.DB.Debug().Model(models.Employee{}).Where("id = ?", eid).Take(&employee).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Employee not found"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	employeeUpdate := models.Employee{}
	err = json.Unmarshal(body, &employeeUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	employeeUpdate.Prepare()
	err = employeeUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	employeeUpdate.ID = employee.ID

	employeeUpdated, err := employeeUpdate.UpdateEmployee(server.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, employeeUpdated)

}

func (server *Server) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	employee := models.Employee{}
	eid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = employee.DeleteEmployee(server.DB, eid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", eid))
	responses.JSON(w, http.StatusNoContent, "")
}
