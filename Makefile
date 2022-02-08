all: test

test:
	go test ./...

petromino: driver/main.go color/*.go state/*.go mino/*.go solver/*.go
	go build -o petromino driver/main.go

run: petromino
	./petromino

.PHONY: test run
