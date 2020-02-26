-include .env
build-linux:
	GOOS=linux go build -o build/main main.go
