vet:
	go vet ./...

build:
	go build ./...

test:
	go test ./...

all: vet build test

tidy:
	go mod tidy --go=1.20
