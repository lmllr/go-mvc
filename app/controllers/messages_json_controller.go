package controllers

import (
	"fmt"
	"net/http"

	"github.com/lmllr/go-mvc/app/helpers"
	"github.com/lmllr/go-mvc/app/models"
)

func CreateMessageProcessJSON(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		return
	}
	if r.Header.Get("content-type") != "application/json" {
		w.WriteHeader(406)
		return
	}

	msgs := []models.MessageJSON{}
	helpers.ParseBody(r, &msgs)

	// DOES NOT WORK
	for _, msg := range msgs {
		if err := msg.CreateFromJSON(); err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
	}

	// WORKS
	// if err := msgs[0].CreateFromJSON(); err != nil {
	// 	http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	// 	return
	// }

	pd := models.PageDataJSON{
		msgs,
		models.RawData{Title: "Show json message"},
	}

	fmt.Fprintln(w, pd)
}
