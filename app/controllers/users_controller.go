package controllers

import (
	"fmt"
	"net/http"

	"github.com/lmllr/go-mvc/app/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		return
	}

	usr := models.User{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("pw"),
	}

	if err := usr.Create(); err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, usr)
}
