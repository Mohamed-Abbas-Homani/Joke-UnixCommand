build:
	go build -o ./bin/joke

run: build
	./bin/joke