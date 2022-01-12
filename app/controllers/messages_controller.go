package controllers

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/lmllr/go-mvc/app/models"
)

// Init templage
var Tpl *template.Template

func init() {
	// Create template
	Tpl = template.Must(template.ParseGlob("app/views/*.gohtml"))
}

// GET
// Simple index page
func Index(w http.ResponseWriter, r *http.Request) {
	// Check http request method
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		// http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	// PageData
	page_data := models.RawData{
		Title: "Index",
	}
	pd := models.PageData{
		Data: page_data,
	}

	// Render template
	Tpl.ExecuteTemplate(w, "index.gohtml", pd)
}

// GET
// Show all messages
func MessagesShowAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		return
	}
	msgs, err := models.Messages()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	anonymousStruct := struct {
		Msgs []models.Message
		Data models.RawData
	}{
		msgs,
		models.RawData{Title: "Show all messages"},
	}

	Tpl.ExecuteTemplate(w, "show_all_messages.gohtml", anonymousStruct)
}

// GET
// Show form to create a new message
func CreateMessageForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		return
	}
	page_data := models.RawData{
		Title: "Create message",
	}
	pd := models.PageData{
		Data: page_data,
	}

	Tpl.ExecuteTemplate(w, "create_message.gohtml", pd)
}

// POST
// Create a new message from form
func CreateMessageProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		return
	}
	msg := models.Message{
		Name:    r.FormValue("name"),
		Message: r.FormValue("message"),
	}
	page_data := models.RawData{
		Title: "Show message",
	}
	if err := msg.Create(); err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	pd := models.PageData{
		Msg:  msg,
		Data: page_data,
	}

	Tpl.ExecuteTemplate(w, "show_message.gohtml", pd)
}

// GET
// Show a single message
func ShowMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		return
	}
	id, err := strconv.ParseInt(r.FormValue("id"), 0, 0)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	msg := models.Message{
		Id: id,
	}
	page_data := models.RawData{
		Title: "Show message",
	}
	if err := msg.Show(); err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	pd := models.PageData{
		Msg:  msg,
		Data: page_data,
	}

	Tpl.ExecuteTemplate(w, "show_message.gohtml", pd)
}

// GET
// Show form to update a message
func UpdateMessageForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		return
	}
	id, err := strconv.ParseInt(r.FormValue("id"), 0, 0)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	msg := models.Message{
		Id: id,
	}
	if err = msg.Show(); err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	page_data := models.RawData{
		Title: "Update message",
	}
	pd := models.PageData{
		Msg:  msg,
		Data: page_data,
	}

	Tpl.ExecuteTemplate(w, "update_message.gohtml", pd)
}

// PUT
// Update a message
func UpdateMessageProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(405)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(r.FormValue("id"), 0, 0)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	msg := models.Message{
		Id:      id,
		Name:    r.FormValue("name"),
		Message: r.FormValue("message"),
	}
	page_data := models.RawData{
		Title: "Show message",
	}
	pd := models.PageData{
		Msg:  msg,
		Data: page_data,
	}
	if err := msg.Update(); err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	Tpl.ExecuteTemplate(w, "show_message.gohtml", pd)
}

// DELETE
// Delete a message
func DeleteMessageProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(405)
		return
	}
	id, err := strconv.ParseInt(r.FormValue("id"), 0, 0)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	msg := models.Message{
		Id: id,
	}
	if err := msg.Delete(); err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	// back to all messages
	http.Redirect(w, r, "/messages", http.StatusSeeOther)
}

// DELETE
// Delete all messages
func DeleteAllMessagesProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(405)
		return
	}
	models.DeleteAll()

	// back to all messages
	http.Redirect(w, r, "/messages", http.StatusSeeOther)
}
