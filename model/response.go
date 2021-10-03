package model

type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data Data `json:"data"`
}

type Data struct {
	From string `json:"from"`
	To string `json:"to"`
}