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
	port := os.Getenv("PORT")
	fmt.Println("Running on PORT:", port)

	mux.HandleFunc("/", controllers.Index)
	mux.HandleFunc("/messages/create/process", controllers.CreateMessage)

	// starting up the server
	server := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: mux,
	}
	server.ListenAndServe()

}
