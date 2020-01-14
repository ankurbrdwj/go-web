package controllers

import (
	"net/http"

	"github.com/ankurbrdwj/go-web/go-login/api/responses"
)
//Home file that welcomes us to the API
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
