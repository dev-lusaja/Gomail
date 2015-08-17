package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mailgun/mailgun-go"
	"io/ioutil"
	"log"
	"net/http"
)

// configuration variables Mailgun
var secret_api_key string
var public_api_key string
var domain_name string

type ErrorResponse struct {
	Error error `json:"error"`
}

type SenderResponse struct {
	Responses interface{} `json:"responses"`
}

type ItemResponse struct {
	Id  string `json:"id"`
	To  string `json:"to"`
	Msg string `json:"msg"`
}

type PayLoad struct {
	Data []Mail
}
type Mail struct {
	From    string
	Subject string
	Body    string
	To      string
}

func init() {
	secret_api_key = "xxxxxxxxxxxxx"
	public_api_key = "xxxxxxxxxxxxx"
	domain_name = "xxxxxxxxxxxxx"
}

func Log(msg interface{}) {
	log.Println(msg)
}

func Sender(w http.ResponseWriter, r *http.Request) {
	// response variables for shipping
	var Response []ItemResponse
	var Json_response []byte
	// open connection with Mailgun
	gun := mailgun.NewMailgun(domain_name, secret_api_key, public_api_key)
	// we set the content-type header to JSON for answers
	w.Header().Set("Content-Type", "application/json")
	// reading the body content
	b, _ := ioutil.ReadAll(r.Body)
	// creating a expresion for Payload
	p := &PayLoad{}
	err := json.Unmarshal(b, p)
	if err != nil {
		msg, _ := json.Marshal("{error: 'the format JSON in the request is invalid'}")
		Log(string(msg))
		w.Write(msg)
		return
	}
	for i := 0; i < len(p.Data); i++ {
		// we get the shipment data
		sender := p.Data[i].From
		subject := p.Data[i].Subject
		body := p.Data[i].Body
		recipient := p.Data[i].To
		// we validate the sender and recipient
		sender_valid, _ := gun.ValidateEmail(sender)
		recipient_valid, _ := gun.ValidateEmail(recipient)
		if sender_valid.IsValid != true || recipient_valid.IsValid != true {
			msg, _ := json.Marshal("{error: 'Invalid sender or recipient'}")
			Log(string(msg))
			w.Write(msg)
			return
		}
		// send
		m := mailgun.NewMessage(sender, subject, body, recipient)
		response, id, send_error := gun.Send(m)
		if send_error != nil {
			Log(send_error)
			msg, _ := json.Marshal(ErrorResponse{Error: send_error})
			Log(string(msg))
			w.Write(msg)
		} else {
			item := ItemResponse{id, recipient, response}
			Response = append(Response, item)
			sr := SenderResponse{Response}
			msg, json_error := json.Marshal(sr)
			if json_error != nil {
				// This error is only for parsing JSON
				Log(json_error)
				msg, _ := json.Marshal(ErrorResponse{Error: json_error})
				Log(string(msg))
				w.Write(msg)
				return
			} else {
				Log(string(msg))
				Json_response = msg
			}
		}
	}
	w.Write(Json_response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/Gomail", Sender).Methods("POST")
	Log("Listening on Port 5000")
	http.ListenAndServe(":5000", r)
}
