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

func (server *Server) CreateDepartment(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	department := models.Department{}
	err = json.Unmarshal(body, &department)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	department.Prepare()
	err = department.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	departmentCreated, err := department.SaveDepartment(server.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, departmentCreated.ID))
	responses.JSON(w, http.StatusCreated, departmentCreated)
}

func (server *Server) GetDepartments(w http.ResponseWriter, r *http.Request) {
	department := models.Department{}
	departments, err := department.FindAllDepartments(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, departments)
}

func (server *Server) GetDepartment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	did, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	department := models.Department{}
	findedDepartment, err := department.FindDepartmentByID(server.DB, did)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, findedDepartment)
}

func (server *Server) UpdateDepartment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	did, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	department := models.Department{}
	err = server.DB.Debug().Model(models.Department{}).Where("id = ?", did).Take(&department).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Department not found"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	departmentUpdate := models.Department{}
	err = json.Unmarshal(body, &departmentUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	departmentUpdate.Prepare()
	err = departmentUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	departmentUpdate.ID = department.ID

	departmentUpdated, err := departmentUpdate.UpdateDepartment(server.DB)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, departmentUpdated)

}

func (server *Server) DeleteDepartment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	department := models.Department{}
	did, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = department.DeleteDepartment(server.DB, did)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", did))
	responses.JSON(w, http.StatusNoContent, "")
}
