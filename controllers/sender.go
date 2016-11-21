package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"os"
	
	"github.com/mailgun/mailgun-go"

	"github.com/dev-lusaja/gomail/models"
)

var (
	config models.Config
)

func init() {
	config_file, e := ioutil.ReadFile("configs/sender.json")
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}

	json.Unmarshal(config_file, &config)
}

func Sender(w http.ResponseWriter, r *http.Request) {
	// response variables for shipping
	var Response []models.ItemResponse
	var Json_response []byte
	// open connection with Mailgun
	gun := mailgun.NewMailgun(config.Domain, config.Secret_api_key, config.Public_api_key)
	// we set the content-type header to JSON for answers
	w.Header().Set("Content-Type", "application/json")
	// reading the body content
	b, _ := ioutil.ReadAll(r.Body)
	// creating a expresion for Payload
	p := &models.PayLoad{}
	err := json.Unmarshal(b, p)
	if err != nil {
		msg, _ := json.Marshal("{error: 'the format JSON in the request is invalid'}")
		log.Println(string(msg))
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
			log.Println(string(msg))
			w.Write(msg)
			return
		}
		// send
		m := mailgun.NewMessage(sender, subject, body, recipient)
		response, id, send_error := gun.Send(m)
		if send_error != nil {
			log.Println(send_error)
			msg, _ := json.Marshal(models.ErrorResponse{Error: send_error})
			log.Println(string(msg))
			w.Write(msg)
		} else {
			item := models.ItemResponse{id, recipient, response}
			Response = append(Response, item)
			sr := models.SenderResponse{Response}
			msg, json_error := json.Marshal(sr)
			if json_error != nil {
				// This error is only for parsing JSON
				log.Println(json_error)
				msg, _ := json.Marshal(models.ErrorResponse{Error: json_error})
				log.Println(string(msg))
				w.Write(msg)
				return
			} else {
				log.Println(string(msg))
				Json_response = msg
			}
		}
	}
	w.Write(Json_response)
}