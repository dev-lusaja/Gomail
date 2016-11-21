package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"encoding/json"
	"io/ioutil"

	"github.com/gorilla/mux"

	"github.com/dev-lusaja/gomail/routes"
	"github.com/dev-lusaja/gomail/models"
)

var server models.Server

func init() {
	config_file, e := ioutil.ReadFile("configs/server.json")
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}

	json.Unmarshal(config_file, &server)
}

func main() {
	r := mux.NewRouter()

	// Load routes
	routes.Load(r)
	log.Println(fmt.Sprintf("Listening on Port %d", server.Port))
	listen_port := fmt.Sprintf(":%d", server.Port)

	// Start server
	http.ListenAndServe(listen_port, r)
}