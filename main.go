package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/lmllr/go-mvc/app/controllers"
)

func main() {
	// Print something on startup
	fmt.Println("Starting HTTP server...")
	mux := http.NewServeMux()
	// Create file server for public/static files
	fs := http.FileServer(http.Dir("app/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	port := os.Getenv("PORT")
	fmt.Println("Running on PORT:", port)

	mux.HandleFunc("/", controllers.Index)
	mux.HandleFunc("/messages", controllers.MessagesShowAll)
	mux.HandleFunc("/messages/show", controllers.ShowMessage)
	mux.HandleFunc("/messages/create", controllers.CreateMessageForm)
	mux.HandleFunc("/messages/create/process", controllers.CreateMessageProcess)
	mux.HandleFunc("/messages/update", controllers.UpdateMessageForm)
	mux.HandleFunc("/messages/update/process", controllers.UpdateMessageProcess)

	// starting up the server
	server := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: mux,
	}
	server.ListenAndServe()
}
