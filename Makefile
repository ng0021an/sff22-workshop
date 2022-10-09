deps:
	go get ./...
	npm install

build:
	go build -o bin/main cmd/main.go

run:
	go run cmd/main.go
