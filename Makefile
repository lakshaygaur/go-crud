
build:
	echo "Building build"
	go build -o build/

clean:
	rm -rf build
	rm -rf storage/gocrud.db
	rm -rf gocrud.db

all: build

test-db:
	go test -v ./storage/

run:
	go run ./cmd