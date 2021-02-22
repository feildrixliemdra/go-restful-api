
start:
	@go run main.go

dep:
	@echo "Installing dependencies..."
	@rm -rf vendor
	@go mod tidy && go mod download && go mod vendor
