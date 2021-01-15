package models

type freelance struct {
	Posts []post `json:"posts"`
}

type post struct {
	Profile string `json:"profile"`
	TJM     int    `json:"tjm"`
}
