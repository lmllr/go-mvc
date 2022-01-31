package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lmllr/go-mvc/app/helpers"
	"github.com/lmllr/go-mvc/app/models"
)

// POST
// CREATE: JSON ARRAY
// create messageS
func CreateMessagesJSON(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// w.Header().Add("error", "error msg")
		// w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(405)
		fmt.Fprintf(w, `{"error":"%s"}`, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	if r.Header.Get("content-type") != "application/json" {
		w.WriteHeader(406)
		fmt.Fprintf(w, `{"error":"%s"}`, http.StatusText(http.StatusNotAcceptable))
		return
	}

	msgs := []models.MessageJSON{}
	helpers.ParseBody(r, &msgs) // <-- Parse r.Body to get json data

	for i := range msgs {
		if err := msgs[i].CreateFromJSON(); err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
	}

	res, err := json.MarshalIndent(&msgs, "", "   ")
	if err != nil {
		fmt.Fprintf(w, `{"error":"%s"}`, "error marshalling arguments")
		return
	}

	w.Header().Set("content-type", "application/json")
	// w.WriteHeader(http.StatusOK) // <-- Automatic?
	w.Write(res)
}

// POST
// CREATE: JSON
// create a single message
func CreateMessageJSON(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		fmt.Fprintf(w, `{"error":"%s"}`, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	if r.Header.Get("content-type") != "application/json" {
		w.WriteHeader(406)
		fmt.Fprintf(w, `{"error":"%s"}`, http.StatusText(http.StatusNotAcceptable))
		return
	}

	msg := models.MessageJSON{}
	helpers.ParseBody(r, &msg)

	if err := msg.CreateFromJSON(); err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res, err := json.MarshalIndent(&msg, "", "   ")
	if err != nil {
		fmt.Fprintf(w, `{"error":"%s"}`, "error marshalling arguments")
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

// GET
// READ: JSON ARRAY
// show all messages
func ShowMessagesJSON(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		fmt.Fprintf(w, `{"error":"%s"}`, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	msgs, err := models.MessagesJSON()
	if err != nil {
		http.Error(w, `{"error":"%s"}`, http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(msgs)
	if err != nil {
		fmt.Fprintf(w, `{"error":"%s"}`, "error marshalling arguments")
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(res) // <-- {"..."}%
}

// GET
// READ: JSON
// get a single message by id
func ShowMessageJSON(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// w.Header().Add("error", "error msg")
		w.WriteHeader(405)
		fmt.Fprintf(w, `{"error":"%s"}`, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	if r.Header.Get("content-type") != "application/json" {
		w.WriteHeader(406)
		fmt.Fprintf(w, `{"error":"%s"}`, http.StatusText(http.StatusNotAcceptable))
		return
	}

	msg := models.MessageJSON{}
	helpers.ParseBody(r, &msg)

	if err := msg.Show(); err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(msg)
	if err != nil {
		fmt.Fprintf(w, `{"error":"%s"}`, "error marshalling arguments")
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

// GET
// JSON ARRAY
// get mutliple messageS by id
func ShowMessageJSONARR(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// w.Header().Add("error", "error msg")
		w.WriteHeader(405)
		fmt.Fprintf(w, `{"error":"%s"}`, http.StatusText(http.StatusMethodNotAllowed))
		return
	}
	if r.Header.Get("content-type") != "application/json" {
		w.WriteHeader(406)
		fmt.Fprintf(w, `{"error":"%s"}`, http.StatusText(http.StatusNotAcceptable))
		return
	}

	msgs := []models.MessageJSON{}
	helpers.ParseBody(r, &msgs)

	for i := range msgs {
		if err := msgs[i].Show(); err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
	}

	res, err := json.Marshal(msgs)
	if err != nil {
		fmt.Fprintf(w, `{"error":"%s"}`, "error marshalling arguments")
		return
	}
	// res = append(res, "\n"...)

	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

// PUT
// UPDATE: JSON ARRAY
// update mutliple messages
func UpdateMessagesJSON(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(405)
		return
	}
	if r.Header.Get("content-type") != "application/json" {
		w.WriteHeader(406)
		fmt.Fprintf(w, `{"error":"%s"}`, http.StatusText(http.StatusNotAcceptable))
		return
	}

	msgs := []models.MessageJSON{}
	helpers.ParseBody(r, &msgs)

	for i := range msgs {
		if err := msgs[i].Update(); err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
	}

	res, err := json.Marshal(msgs)
	if err != nil {
		fmt.Fprintf(w, `{"error":"%s"}`, "error marshalling arguments")
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

// PUT
// UPDATE: JSON
// update a single message
func UpdateMessageJSON(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(405)
		return
	}
	if r.Header.Get("content-type") != "application/json" {
		w.WriteHeader(406)
		fmt.Fprintf(w, `{"error":"%s"}`, http.StatusText(http.StatusNotAcceptable))
		return
	}

	msg := models.MessageJSON{}
	helpers.ParseBody(r, &msg)

	if err := msg.Update(); err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(msg)
	if err != nil {
		fmt.Fprintf(w, `{"error":"%s"}`, "error marshalling arguments")
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

// DELETE
// DELETE: JSON
// update a single message
func DeleteMessageJSON(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(405)
		return
	}
	if r.Header.Get("content-type") != "application/json" {
		w.WriteHeader(406)
		fmt.Fprintf(w, `{"error":"%s"}`, http.StatusText(http.StatusNotAcceptable))
		return
	}

	msg := models.MessageJSON{}
	helpers.ParseBody(r, &msg)

	if err := msg.DeleteJSON(); err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	// w.Header().Set("content-type", "application/json")
	// w.Write(res)
	msgs, err := models.MessagesJSON()
	if err != nil {
		http.Error(w, `{"error":"%s"}`, http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(msgs)
	if err != nil {
		fmt.Fprintf(w, `{"error":"%s"}`, "error marshalling arguments")
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(res) // <-- {"..."}%
}
