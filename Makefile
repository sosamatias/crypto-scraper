install:
	go get ./...

test:
	go test ./... -cover

run:
	go run main.go
