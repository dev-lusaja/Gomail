package routes

import(
	"github.com/gorilla/mux"

	"../controllers"
)

func Load(r *mux.Router) {

	// Sender Route
	r.HandleFunc("/api/v1/gomail", controllers.Sender).Methods("POST")
}