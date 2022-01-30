package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lmllr/go-mvc/app/helpers"
	"github.com/lmllr/go-mvc/app/models"
)

// POST
// READ JSON ARRAY
// create messageS
func CreateMessagesProcessJSON(w http.ResponseWriter, r *http.Request) {
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

	w.Write(res)
	// fmt.Fprintln(w, string(res))
}

// POST
// READ JSON
// create a single message
func CreateMessageProcessJSON(w http.ResponseWriter, r *http.Request) {
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

	fmt.Fprintln(w, string(res))
}

// GET
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
	// w.Write(res) // <-- {"..."}%
	fmt.Fprintln(w, string(res))
}

// GET
// get message by id
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
	res = append(res, "\n"...)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
