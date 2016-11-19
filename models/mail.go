package models

type PayLoad struct {
	Data []Mail
}

type Mail struct {
	From    string
	Subject string
	Body    string
	To      string
}