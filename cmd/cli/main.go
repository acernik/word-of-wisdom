package main

import (
	"fmt"
	"net"

	"github.com/acernik/word-of-wisdom/internal/client"
	"github.com/acernik/word-of-wisdom/internal/config"
	"github.com/acernik/word-of-wisdom/internal/pow"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	tcpAddr, err := net.ResolveTCPAddr(cfg.App.Network, cfg.App.ClientAddress)
	if err != nil {
		panic(err)
	}

	// Create a new value of type client Requester.
	cr, err := client.New(cfg)
	if err != nil {
		panic(err)
	}

	// Make initial request to the server which will request to do the PoW.
	_, err = cr.MakeInitialRequest(tcpAddr)
	if err != nil {
		panic(err)
	}

	// Perform PoW.
	pg, err := pow.New()
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

	if result.Type == cfg.App.ResponseTypePowInvalid {
		fmt.Println("The PoW solution was invalid, please try again.")
		return
	}

	// Print the result.
	result.ToString()
}
