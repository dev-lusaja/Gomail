package models

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