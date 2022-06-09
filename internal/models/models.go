package models

import (
	"fmt"

	"github.com/acernik/word-of-wisdom/internal/quotes"
)

// Request holds the data sent with the request to the server.
type Request struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// Response holds the data sent with the response to the cli.
type Response struct {
	Type  string       `json:"type"`
	Quote quotes.Quote `json:"quote"`
}

// ToString prints out all the data associated with a quote.
func (r Response) ToString() {
	fmt.Println("Author: " + r.Quote.Author)
	fmt.Println("Country: " + r.Quote.Country)
	fmt.Println("Occupation: " + r.Quote.Occupation)
	fmt.Println("Quote: " + r.Quote.Quote)
}
