tidy:
	@go mod tidy
	@go fmt

build:
	@go build -o garbage

package: build
	@zip garbage.zip garbage garbage.csv

test:
	@go test -v ./...
