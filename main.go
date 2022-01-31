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
	fmt.Printf("Visit http://localhost:%s\n", port)

	//> Content-Type: application/x-www-form-urlencoded
	mux.HandleFunc("/", controllers.Index)
	mux.HandleFunc("/messages", controllers.MessagesShowAll)
	mux.HandleFunc("/messages/show", controllers.ShowMessage)
	mux.HandleFunc("/messages/create", controllers.CreateMessageForm)
	mux.HandleFunc("/messages/create/process", controllers.CreateMessageProcess)
	mux.HandleFunc("/messages/update", controllers.UpdateMessageForm)
	mux.HandleFunc("/messages/update/process", controllers.UpdateMessageProcess)
	mux.HandleFunc("/messages/delete/process", controllers.DeleteMessageProcess)
	mux.HandleFunc("/messages/deleteall/process", controllers.DeleteAllMessagesProcess)

	//> Content-Type: application/json
	mux.HandleFunc("/json/messages", controllers.ShowMessagesJSON)
	mux.HandleFunc("/json/messages/showjson", controllers.ShowMessageJSON)
	mux.HandleFunc("/json/messages/showjsonarr", controllers.ShowMessageJSONARR)
	mux.HandleFunc("/json/messages/createjson", controllers.CreateMessageJSON)
	mux.HandleFunc("/json/messages/createjsons", controllers.CreateMessagesJSON)
	mux.HandleFunc("/json/messages/updatejson", controllers.UpdateMessageJSON)
	mux.HandleFunc("/json/messages/updatejsons", controllers.UpdateMessagesJSON)
	mux.HandleFunc("/json/messages/delete", controllers.DeleteMessageJSON)

	// starting up the server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	server.ListenAndServe()
}
