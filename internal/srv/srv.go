package srv

import (
	"bytes"
	"encoding/json"
	"log"
	"net"

	"github.com/acernik/word-of-wisdom/internal/constants"
	"github.com/acernik/word-of-wisdom/internal/models"
	"github.com/acernik/word-of-wisdom/internal/pow"
	"github.com/acernik/word-of-wisdom/internal/quotes"
)

// Handler interface defines all the methods that are used to handle incoming requests from clients.
type Handler interface {
	HandleRequest(conn net.Conn, qp quotes.Picker)
}

// handler is the type that implements the Handler interface.
type handler struct {
}

// New returns a new value of type that implements the Handler interface.
func New() Handler {
	return &handler{}
}

// HandleRequest handles incoming requests from clients.
func (h *handler) HandleRequest(conn net.Conn, qp quotes.Picker) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	b := bytes.NewBuffer(buffer)
	d := json.NewDecoder(b)

	var request models.Request
	err = d.Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	switch request.Type {
	case constants.RequestTypeInitial:
		initialResponse := models.Response{
			Type:  constants.ResponseTypePow,
			Quote: quotes.Quote{},
		}
		initialResponseBytes, err := json.Marshal(initialResponse)
		if err != nil {
			log.Fatal(err)
		}

		_, err = conn.Write(initialResponseBytes)
		if err != nil {
			log.Fatal(err)
		}
	case constants.RequestTypePowSolution:
		gp := pow.New()
		valid, err := gp.Verify(request.Data)
		if err != nil {
			log.Fatal(err)
		}

		if !valid {
			initialResponse := models.Response{
				Type:  constants.ResponseTypePowInvalid,
				Quote: quotes.Quote{},
			}
			initialResponseBytes, err := json.Marshal(initialResponse)
			if err != nil {
				log.Fatal(err)
			}

			_, err = conn.Write(initialResponseBytes)
			if err != nil {
				log.Fatal(err)
			}

			break
		}

		idx, err := qp.PickQuoteIndex()
		if err != nil {
			log.Fatal(err)
		}

		quote, err := qp.GetQuote(idx)
		if err != nil {
			log.Fatal(err)
		}

		result := models.Response{
			Type:  constants.ResponseTypePowValid,
			Quote: quote,
		}

		resultBytes, err := json.Marshal(result)
		if err != nil {
			log.Fatal(err)
		}

		_, err = conn.Write(resultBytes)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}
}
