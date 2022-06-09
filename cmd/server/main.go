package main

import (
	"net"

	"github.com/acernik/word-of-wisdom/internal/config"
	"github.com/acernik/word-of-wisdom/internal/quotes"
	"github.com/acernik/word-of-wisdom/internal/srv"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	listen, err := net.Listen(cfg.App.Network, cfg.App.ServerAddress)
	if err != nil {
		panic(err)
	}

	defer listen.Close()

	qp := quotes.New(quotes.Quotes)

	h, err := srv.New(cfg)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}

		go h.HandleRequest(conn, qp)
	}
}
