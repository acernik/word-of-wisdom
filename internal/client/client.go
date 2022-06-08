package client

import (
	"bytes"
	"encoding/json"
	"net"

	"github.com/acernik/word-of-wisdom/internal/constants"
	"github.com/acernik/word-of-wisdom/internal/models"
)

// Requester interface defines all the methods that are used by the client to make request to the server.
type Requester interface {
	MakeInitialRequest(tcpAddr *net.TCPAddr) (models.Response, error)
	MakePowSolutionRequest(solution string, tcpAddr *net.TCPAddr) (models.Response, error)
}

// requester is the type that implements the Requester interface.
type requester struct{}

// New returns the value of type that implements the Requester interface.
func New() Requester {
	return &requester{}
}

// MakeInitialRequest makes the initial request to the server requesting a quote. The server will respond with
// request to find a Proof of Work solution.
func (r *requester) MakeInitialRequest(tcpAddr *net.TCPAddr) (models.Response, error) {
	var initialResponse models.Response

	conn, err := net.DialTCP(constants.Network, nil, tcpAddr)
	if err != nil {
		return initialResponse, err
	}

	initialRequest := models.Request{
		Type: constants.RequestTypeInitial,
		Data: constants.InitialRequestMessage,
	}

	initialRequestBytes, err := json.Marshal(initialRequest)
	if err != nil {
		return initialResponse, err
	}

	// Make initial initialRequest. The response should be initialRequest to perform PoW.
	_, err = conn.Write(initialRequestBytes)
	if err != nil {
		return initialResponse, err
	}

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)
	if err != nil {
		return initialResponse, err
	}

	err = conn.Close()
	if err != nil {
		return initialResponse, err
	}

	// Read the response from the server.
	b := bytes.NewBuffer(reply)
	d := json.NewDecoder(b)

	err = d.Decode(&initialResponse)
	if err != nil {
		return initialResponse, err
	}

	return initialResponse, nil
}

// MakePowSolutionRequest makes request to the server after finding Proof of Work solution. If Proof of Work
// solution is correct this request will receive the response with a quote.
func (r *requester) MakePowSolutionRequest(solution string, tcpAddr *net.TCPAddr) (models.Response, error) {
	var quoteResponse models.Response

	powSolutionRequest := models.Request{
		Type: constants.RequestTypePowSolution,
		Data: solution,
	}

	powSolutionRequestBytes, err := json.Marshal(powSolutionRequest)
	if err != nil {
		return quoteResponse, err
	}

	conn, err := net.DialTCP(constants.Network, nil, tcpAddr)
	if err != nil {
		return quoteResponse, err
	}

	// Make powSolutionRequest. The response should be one of the quotes.
	_, err = conn.Write(powSolutionRequestBytes)
	if err != nil {
		return quoteResponse, err
	}

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)
	if err != nil {
		return quoteResponse, err
	}

	// Read the response from the server.
	b := bytes.NewBuffer(reply)
	d := json.NewDecoder(b)

	err = d.Decode(&quoteResponse)
	if err != nil {
		return quoteResponse, err
	}

	err = conn.Close()
	if err != nil {
		return quoteResponse, err
	}

	return quoteResponse, nil
}
