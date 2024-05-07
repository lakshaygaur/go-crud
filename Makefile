
build:
	echo "Building build"
	go build -o build/ ./cmd

build-alpine:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/ ./cmd

clean:
	rm -rf build
	rm -rf storage/gocrud.db
	rm -rf gocrud.db

all: build

test-db:
	go test -v ./storage/

run:
	go run ./cmd