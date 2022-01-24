package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lmllr/go-mvc/app/helpers"
	"github.com/lmllr/go-mvc/app/models"
)

// POST
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
	helpers.ParseBody(r, &msgs)

	for i := range msgs {
		if err := msgs[i].CreateFromJSON(); err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
	}

	// pd := models.PageDataJSON{
	// 	msgs,
	// 	models.RawData{Title: "Show json message"},
	// }

	test, err := json.MarshalIndent(&msgs, "", "   ")
	if err != nil {
		return
	}

	fmt.Fprintln(w, string(test))
}
