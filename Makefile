user-service:
		go build -v -o user-service ./user_service/cmd/user_service
lint:
		golangci-lint run --fix --fast
fmt:
		gofmt -s -w .
check: fmt lint


.DEFAULT_GOAL := user-service