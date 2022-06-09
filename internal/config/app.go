package config

// App holds the configuration values for the application.
type App struct {
	Network                string `env:"NETWORK" yaml:"network" default:"tcp"`
	ServerAddress          string `env:"SERVER_ADDRESS" yaml:"server_address" default:"0.0.0.0:9001"`
	ClientAddress          string `env:"CLIENT_ADDRESS" yaml:"client_address" default:"192.168.1.210:9001"`
	RequestTypeInitial     string `env:"REQUEST_TYPE_INITIAL" yaml:"request_type_initial" default:"INITIAL"`
	ResponseTypePow        string `env:"RESPONSE_TYPE_POW" yaml:"response_type_pow" default:"POW_REQUEST"`
	ResponseTypePowInvalid string `env:"RESPONSE_TYPE_POW_INVALID" yaml:"response_type_pow_invalid" default:"POW_INVALID"`
	ResponseTypePowValid   string `env:"RESPONSE_TYPE_POW_VALID" yaml:"response_type_pow_valid" default:"POW_VALID"`
	RequestTypePowSolution string `env:"REQUEST_TYPE_POW_SOLUTION" yaml:"request_type_pow_solution" default:"POW_SOLUTION"`
	InitialRequestMessage  string `env:"INITIAL_REQUEST_MESSAGE" yaml:"initial_request_message" default:"Hello, can I get a quote please?"`
}
