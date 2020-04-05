package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"time"

	"github.com/mohdfidzrin/webservice/models"
)

type UserController struct {
	userIDPattern *regexp.Regexp
}

func (uc UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		ID:   1,
		Name: "Rein",
		DOB:  time.Now(),
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		w.Write([]byte("error"))
		return
	}
	w.Write([]byte(string(userJson)))
}

func newUserController() *UserController {
	return &UserController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}