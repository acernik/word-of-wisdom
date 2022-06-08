client:
	go run cmd/app/main.go

server:
	go run cmd/server/main.go

tests:
	go test ./... -v -coverpkg=./... -coverprofile=profile.cov ./...
	go tool cover -func profile.cov

PHONY: client server tests