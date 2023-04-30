build-user-service:
		go build -v ./user_service/cmd/user
lint:
		golangci-lint run --fix --fast
fmt:
		gofmt -s -w .


.DEFAULT_GOAL := build-user-service
.PHONY: build-user-service lint fmt