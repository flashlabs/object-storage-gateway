# Spacelift Homework Object Storage Makefile

# lint: runs a golangci-lint with the same settings as in the CI.
lint:
	golangci-lint run ./...

# check: executes a static check.
check:
	staticcheck ./...

# executes a test suite.
test:
	go test ./...

# builds application.
build:
	env GOOS=linux GOARCH=amd64 go build -o homework-object-storage main.go
