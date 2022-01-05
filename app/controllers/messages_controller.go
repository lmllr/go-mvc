package controllers

import (
	"net/http"
	"text/template"

	"github.com/lmllr/go-mvc/app/models"
)

// Init template
var Tpl *template.Template

// Func init() will get called on main initialization
func init() {
	// Create template
	Tpl = template.Must(template.ParseGlob("app/views/*.gohtml"))
}

func Index(w http.ResponseWriter, r *http.Request) {
	msgs, err := models.Messages()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	Tpl.ExecuteTemplate(w, "index.gohtml", msgs)
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	msg := models.Message{
		Name:    r.FormValue("name"),
		Message: r.FormValue("message"),
	}
	if err := msg.Create(); err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	Tpl.ExecuteTemplate(w, "messages.gohtml", msg)
}
