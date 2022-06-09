package srv

import (
	"encoding/json"
	"net"
	"testing"

	"github.com/acernik/word-of-wisdom/internal/config"
	"github.com/acernik/word-of-wisdom/internal/models"
	"github.com/acernik/word-of-wisdom/internal/pow"
	"github.com/acernik/word-of-wisdom/internal/quotes"
)

func TestSrv_New(t *testing.T) {
	cfg, err := config.New()
	if err != nil {
		t.Error(err)
	}

	mockHandler, err := New(cfg)
	if err != nil {
		t.Error(err)
	}

	if mockHandler == nil {
		t.Errorf("expected: value of type handler to be not nil")
	}
}

func TestSrv_HandleRequest_InitialRequest(t *testing.T) {
	cfg, err := config.New()
	if err != nil {
		t.Error(err)
	}

	mockHandler := &handler{
		cfg: cfg,
	}

	go func() {
		tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:8002")
		if err != nil {
			t.Error(err)
		}

		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			t.Error(err)
		}

		initialRequest := models.Request{
			Type: mockHandler.cfg.App.RequestTypeInitial,
			Data: mockHandler.cfg.App.InitialRequestMessage,
		}

		initialRequestBytes, err := json.Marshal(initialRequest)
		if err != nil {
			t.Error(err)
		}

		_, err = conn.Write(initialRequestBytes)
		if err != nil {
			t.Error(err)
		}
	}()

	listen, err := net.Listen("tcp", "localhost:8002")
	if err != nil {
		t.Fatal(err)
	}

	defer listen.Close()

	qp := quotes.New(quotes.Quotes)

	for {
		conn, err := listen.Accept()
		if err != nil {
			t.Fatal(err)
		}

		mockHandler.HandleRequest(conn, qp)

		break
	}
}

func TestSrv_HandleRequest_PowRequest(t *testing.T) {
	cfg, err := config.New()
	if err != nil {
		t.Error(err)
	}

	mockHandler := &handler{
		cfg: cfg,
	}

	go func() {
		tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:8003")
		if err != nil {
			t.Error(err)
		}

		conn, err := net.DialTCP("tcp", nil, tcpAddr)
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

		powSolutionRequest := models.Request{
			Type: mockHandler.cfg.App.RequestTypePowSolution,
			Data: solution,
		}

		powSolutionRequestBytes, err := json.Marshal(powSolutionRequest)
		if err != nil {
			t.Error(err)
		}

		_, err = conn.Write(powSolutionRequestBytes)
		if err != nil {
			t.Error(err)
		}
	}()

	listen, err := net.Listen("tcp", "localhost:8003")
	if err != nil {
		t.Fatal(err)
	}

	defer listen.Close()

	qp := quotes.New(quotes.Quotes)

	for {
		conn, err := listen.Accept()
		if err != nil {
			t.Fatal(err)
		}

		mockHandler.HandleRequest(conn, qp)

		break
	}
}
