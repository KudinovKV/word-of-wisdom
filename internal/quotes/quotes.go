package quotes

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

type Quotes struct {
	Quotes []Quote
}

func NewQuotes() (*Quotes, error) {
	r, err := http.Get("https://type.fit/api/quotes")
	if err != nil {
		return nil, err
	}
	var quotes Quotes
	if err := json.NewDecoder(r.Body).Decode(&quotes.Quotes); err != nil {
		return nil, err
	}
	return &quotes, nil
}

func (q Quotes) GetRandomQuote() Quote {
	return q.Quotes[rand.Intn(len(q.Quotes))]
}
