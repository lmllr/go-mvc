package controllers

import (
	"net/http"

	"github.com/lmllr/go-mvc/app/helpers"
	"github.com/lmllr/go-mvc/app/models"
)

func CreateMessageProcessJSON(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("content-type") != "application/json" ||
		r.Method != http.MethodPost {
		w.WriteHeader(406)
		return

	}
	msg := &models.MessageJSON{}
	helpers.ParseBody(r, msg)

	if err := msg.CreateFromJSON(); err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	pd := models.PageDataJSON{
		*msg,
		models.RawData{Title: "TEST"},
	}

	Tpl.ExecuteTemplate(w, "show_message.gohtml", pd)
}
