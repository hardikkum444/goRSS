build:
	@go build -o bin/goRSS

run: build
	@./bin/goRSS 

test:
	@go test -b ./...
