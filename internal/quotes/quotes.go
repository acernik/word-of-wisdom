package quotes

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Picker interface defines all the methods used to get quotes.
type Picker interface {
	PickQuoteIndex() (int64, error)
	GetQuote(id int64) (Quote, error)
}

// Quote holds all the data about a quote.
type Quote struct {
	Author     string `json:"author"`
	Country    string `json:"country"`
	Occupation string `json:"occupation"`
	Quote      string `json:"quote"`
}

// quotesPicker is the type that implements Picker interface.
type quotesPicker struct {
	quotesList []Quote
}

// New returns new value of type that implements Picker interface.
func New(quotesList []Quote) Picker {
	return &quotesPicker{
		quotesList: quotesList,
	}
}

// All the quotes are coming from this book -> https://globalyoungacademy.net/wp-content/uploads/2016/06/Words_of_Wisdom_GYA_2016.pdf
var (
	Quotes = []Quote{
		{
			Author:     "Adewale Adewuyi",
			Country:    "Nigeria",
			Occupation: "Industrial Chemistry",
			Quote:      "Think globally but act locally.",
		},
		{
			Author:     "Gregory Weiss",
			Country:    "USA",
			Occupation: "Chemical Biology",
			Quote:      "Don’t worry about the competition.\nOnly worry if you’re working in an area\nthat’s not attracting any competition.",
		},
		{
			Author:     "Rapela Maphanga",
			Country:    "South Africa",
			Occupation: "Physics",
			Quote:      "Don’t be a Jack of all trades and master of none.",
		},
		{
			Author:     "Masanori Arita",
			Country:    "Japan",
			Occupation: "Bioinformatics",
			Quote:      "Be sticky to your research goal.",
		},
		{
			Author:     "Hong Ching Goh",
			Country:    "Malaysia",
			Occupation: "Geography and Urban Planning",
			Quote:      "Never lose the sight of a forest just for a tree.",
		},
		{
			Author:     "Heather Ferguson",
			Country:    "UK",
			Occupation: "Cognitive Psychology",
			Quote:      "Stick to the big questions, don’t get bogged down in the details.",
		},
		{
			Author:     "Annemieke Aartsma-Rus",
			Country:    "Netherlands",
			Occupation: "Genetics & Therapy",
			Quote: `Make sure you don\’t waste your potential on
			a mediocre project, and be sure not to pursue
			something indefinitely. Question whether
			you are on the right track. If the answer is
			No, then move on to something else.`,
		},
		{
			Author:     "Matthew Finkbeiner",
			Country:    "Australia",
			Occupation: "Cognitive Science",
			Quote:      "Ideas are free, it’s what you do with them that counts.",
		},
		{
			Author:     "Rawiwan Laocharoensuk",
			Country:    "Thailand",
			Occupation: "Nanomaterial Chemistry",
			Quote: `Do something that interests you and do it to the
			absolute best of your ability. With this recipe,
			whatever your limitations, you will almost
			certainly still do better than anyone else.`,
		},
		{
			Author:     "Aron Telegdi-Csetr",
			Country:    "Romania",
			Occupation: "Philosophy",
			Quote: `In the Humanities, we often busy ourselves
			with the past, or with dead people. We must
			remember though that we are alive, and
			must act accordingly. To be useful to life,
			be independent, assertive, and young.`,
		},
	}
	ErrInvalidQuoteIndex = fmt.Errorf("invalid quote index, must be 0 or greater and less than length of quotes array")
)

// PickQuoteIndex picks a random quote element index from the quotes array.
func (qp *quotesPicker) PickQuoteIndex() (int64, error) {
	if len(qp.quotesList) <= 0 {
		return 0, fmt.Errorf("quotes list must contain at least one quote")
	}

	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(qp.quotesList))))
	if err != nil {
		return 0, err
	}

	return nBig.Int64(), nil
}

// GetQuote returns a quote from the quotes array at specified index.
func (qp *quotesPicker) GetQuote(id int64) (Quote, error) {
	if id < 0 || int(id) >= len(qp.quotesList) {
		return Quote{}, ErrInvalidQuoteIndex
	}

	return qp.quotesList[id], nil
}
