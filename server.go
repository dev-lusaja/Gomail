package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mailgun/mailgun-go"
	"log"
	"net/http"
)

var secret_api_key string
var public_api_key string
var domain_name string

type SenderResponse struct {
	ResponseID string `json:"id"`
	ServerMsg  string `json:"message"`
}
type ErrorResponse struct {
	Error error `json:"error"`
}

func init() {
	secret_api_key = "secret_key"
	public_api_key = "public_key"
	domain_name = "domain"
}

func Log(msg interface{}) {
	log.Println(msg)
}

func Sender(w http.ResponseWriter, r *http.Request) {
	// Seteamos el contenido de la cabecera a tipo JSON para las respuestas
	w.Header().Set("Content-Type", "application/json")

	gun := mailgun.NewMailgun(domain_name, secret_api_key, public_api_key)

	// recibimos los datos para enviar el email
	sender := r.FormValue("from")
	subject := r.FormValue("subject")
	body := r.FormValue("body")
	recipient := r.FormValue("to")

	// Validamos el sender y el recipient
	sender_valid, _ := gun.ValidateEmail(sender)
	recipient_valid, _ := gun.ValidateEmail(recipient)
	if sender_valid.IsValid != true || recipient_valid.IsValid != true {
		msg, _ := json.Marshal("{error: 'Invalid sender or recipient'}")
		w.Write(msg)
		return
	}

	// enviamos el mensaje
	m := mailgun.NewMessage(sender, subject, body, recipient)
	response, id, send_error := gun.Send(m)
	if send_error != nil {
		// mostramos el error de envio
		Log(send_error)
		msg, _ := json.Marshal(ErrorResponse{Error: send_error})
		w.Write(msg)
	} else {
		msg, json_error := json.Marshal(SenderResponse{ResponseID: id, ServerMsg: response})
		if json_error != nil {
			// mostramos el error de conversion a JSON
			Log(json_error)
			msg, _ := json.Marshal(ErrorResponse{Error: json_error})
			w.Write(msg)
		} else {
			// mostramos la respuesta de envio del mensaje
			w.Write(msg)
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/Gomail", Sender).Methods("POST")
	Log("Listening on Port 5000")
	http.ListenAndServe(":5000", r)
}
