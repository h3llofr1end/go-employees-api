package controllers

import (
	"net/http"

	"github.com/h3llofr1end/go-employees-api/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to Employee API")
}
