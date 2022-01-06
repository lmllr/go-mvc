package controllers

import (
	"net/http"
	"strconv"
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
	Tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func MessagesShowAll(w http.ResponseWriter, r *http.Request) {
	msgs, err := models.Messages()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	Tpl.ExecuteTemplate(w, "messages.gohtml", msgs)
}

func MessagesCreateForm(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "create.gohtml", nil)
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
	Tpl.ExecuteTemplate(w, "show_message.gohtml", msg)
}

func ShowMessage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.FormValue("id"), 0, 0)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	msg := models.Message{
		Id: id,
	}
	if err := msg.Show(); err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	Tpl.ExecuteTemplate(w, "show_message.gohtml", msg)
}
