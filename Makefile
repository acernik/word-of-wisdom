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

dkb-cli:
	docker build -f Dockerfile.client -t client .

dkr-cli:
	docker run -p "9002:9001" client

launch-cli: dkb-cli dkr-cli

PHONY: client server tests dkb-srv dkr-srv launch-srv dkb-cli dkr-cli launch-cli