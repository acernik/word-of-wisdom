package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_New(t *testing.T) {
	err := os.Setenv("TEST_CONFIG", "config.test.yml")
	if err != nil {
		t.Fatal(err)
	}

	cfg, err := New()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "tcp", cfg.App.Network, "The network should be tcp.")
	assert.Equal(t, "0.0.0.0:9001", cfg.App.ServerAddress, "The server_address should be 0.0.0.0:9001.")
	assert.Equal(t, "192.168.1.210:9001", cfg.App.ClientAddress, "The client_address should be 192.168.1.210:9001.")
	assert.Equal(t, "INITIAL", cfg.App.RequestTypeInitial, "The request_type_initial should be INITIAL.")
	assert.Equal(t, "POW_REQUEST", cfg.App.ResponseTypePow, "The response_type_pow should be POW_REQUEST.")
	assert.Equal(t, "POW_INVALID", cfg.App.ResponseTypePowInvalid, "The response_type_pow_invalid should be POW_INVALID.")
	assert.Equal(t, "POW_VALID", cfg.App.ResponseTypePowValid, "The response_type_pow_valid should be POW_VALID.")
	assert.Equal(t, "POW_SOLUTION", cfg.App.RequestTypePowSolution, "The request_type_pow_solution should be POW_SOLUTION.")
	assert.Equal(t, "Hello, can I get a quote please?", cfg.App.InitialRequestMessage, "The initial_request_message should be Hello, can I get a quote please?.")
}
