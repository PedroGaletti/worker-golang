install:
	@echo "Downloading dependecies..."
	@go get
	@echo "Validating dependecies..."
	@go mod tidy
	@echo "Creating vendor..."
	@go mod vendor
	@echo "Installation completed successfully."

build:
	@echo "Building project..."
	@go build
	@echo "Build completed successfully."

run:
	@echo "Running application..."
	@go run main.go

clean:
	@echo "Cleaning up project..."
	@rm -rf ./vendor
	@rm -rf ./go.sum
	@rm -rf ./worker-golang
	@echo "Project cleaned successfully."
