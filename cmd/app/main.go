package main

import (
	"fmt"
	"net"

	"github.com/acernik/word-of-wisdom/internal/client"
	"github.com/acernik/word-of-wisdom/internal/constants"
	"github.com/acernik/word-of-wisdom/internal/pow"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr(constants.Network, constants.Address)
	if err != nil {
		panic(err)
	}

	// Create a new value of type client Requester.
	cr := client.New()

	// Make initial request to the server which will request to do the PoW.
	_, err = cr.MakeInitialRequest(tcpAddr)

	// Perform PoW.
	pg := pow.New()
	if err != nil {
		panic(err)
	}

	solution, err := pg.GetPowSolution()
	if err != nil {
		panic(err)
	}

	// Send the PoW solution to the server. This will return a quote if PoW solution is correct.
	result, err := cr.MakePowSolutionRequest(solution, tcpAddr)
	if err != nil {
		panic(err)
	}

	if result.Type == constants.ResponseTypePowInvalid {
		fmt.Println("The PoW solution was invalid, please try again.")
		return
	}

	// Print the result.
	result.ToString()
}
