.PHONY: default run build test docs clean

# Variables
APP_NAME=gopportunities

#Tasks
default: run-with-docs

run:
	@go run cmd/main.go
run-with-docs:
	@swag init --parseDependency --parseInternal --generalInfo cmd/main.go --output ./docs
	@go run cmd/main.go
build:
	@go build -o $(APP_NAME) cmd/main.go
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs