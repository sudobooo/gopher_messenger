user-service:
		go build -v -o user-service ./user_service/cmd/user_service
lint:
		golangci-lint run --fix --fast
fmt:
		gofmt -s -w .
test:
		go test -v -race -timeout 30s ./...
check: fmt lint test


.DEFAULT_GOAL := user-service