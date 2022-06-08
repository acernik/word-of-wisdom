package main

import (
	"net"

	"github.com/acernik/word-of-wisdom/internal/constants"
	"github.com/acernik/word-of-wisdom/internal/quotes"
	"github.com/acernik/word-of-wisdom/internal/srv"
)

func main() {
	listen, err := net.Listen(constants.Network, constants.Address)
	if err != nil {
		panic(err)
	}

	defer listen.Close()

	qp := quotes.New(quotes.Quotes)

	h := srv.New()

	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}

		go h.HandleRequest(conn, qp)
	}
}
