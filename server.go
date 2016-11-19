package main

import (
	"net/http"
	"fmt"
	"log"

	"github.com/gorilla/mux"

	"./routes"
)

var port int

func init() {
	port = 5000
}

func main() {
	r := mux.NewRouter()

	// Load routes
	routes.Load(r)
	log.Println(fmt.Sprintf("Listening on Port %d", port))
	listen_port := fmt.Sprintf(":%d", port)

	// Start server
	http.ListenAndServe(listen_port, r)
}