client:
	go run cmd/cli/main.go

server:
	go run cmd/server/main.go

tests:
	go test ./... -v -coverpkg=./... -coverprofile=profile.cov ./...
	go tool cover -func profile.cov

dkb-srv:
	docker build -f Dockerfile.server -t server .

dkr-srv:
	docker run -p "9001:9001" server

launch-srv: dkb-srv dkr-srv

exp_addr:
	./run_cli.sh

dkb-cli:
	docker build -f Dockerfile.client -t client .

dkr-cli:
	docker run -p "9002:9001" client

launch-cli: exp_addr dkb-cli dkr-cli

PHONY: client server tests dkb-srv dkr-srv launch-srv exp_addr dkb-cli dkr-cli launch-cli