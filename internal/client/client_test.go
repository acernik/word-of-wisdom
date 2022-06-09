package client

import (
	"net"
	"testing"

	"github.com/acernik/word-of-wisdom/internal/config"
	"github.com/acernik/word-of-wisdom/internal/pow"
	"github.com/acernik/word-of-wisdom/internal/quotes"
	"github.com/acernik/word-of-wisdom/internal/srv"
)

func TestClient_New(t *testing.T) {
	cfg, err := config.New()
	if err != nil {
		t.Error(err)
	}

	mockRequester, err := New(cfg)
	if err != nil {
		t.Error(err)
	}

	if mockRequester == nil {
		t.Errorf("expected: value of type requester to be not nil")
	}
}

func TestClient_MakeInitialRequest(t *testing.T) {
	cfg, err := config.New()
	if err != nil {
		t.Error(err)
	}

	go func() {
		tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:8000")
		if err != nil {
			t.Error(err)
		}

		mockRequester, err := New(cfg)
		if err != nil {
			t.Error(err)
		}

		_, err = mockRequester.MakeInitialRequest(tcpAddr)
		if err != nil {
			t.Error(err)
		}
	}()

	listen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		t.Fatal(err)
	}

	defer listen.Close()

	qp := quotes.New(quotes.Quotes)

	mockHandler, err := srv.New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			t.Fatal(err)
		}

		mockHandler.HandleRequest(conn, qp)

		break
	}
}

func TestClient_MakePowSolutionRequest(t *testing.T) {
	cfg, err := config.New()
	if err != nil {
		t.Error(err)
	}

	go func() {
		tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:8001")
		if err != nil {
			t.Error(err)
		}

		mockRequester, err := New(cfg)
		if err != nil {
			t.Error(err)
		}

		pg, err := pow.New()
		if err != nil {
			t.Error(err)
		}

		solution, err := pg.GetPowSolution()
		if err != nil {
			t.Error(err)
		}

		_, err = mockRequester.MakePowSolutionRequest(solution, tcpAddr)
		if err != nil {
			t.Error(err)
		}
	}()

	listen, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		t.Fatal(err)
	}

	defer listen.Close()

	qp := quotes.New(quotes.Quotes)

	mockHandler, err := srv.New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			t.Fatal(err)
		}

		mockHandler.HandleRequest(conn, qp)

		break
	}
}
