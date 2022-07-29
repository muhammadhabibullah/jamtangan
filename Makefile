
.PHONY: env mock

env:
	@cp -n config/config.example.json config/config.json

mock:
	@if ! command -v mockery &> /dev/null; then go install github.com/vektra/mockery/v2@v2.10.0; fi
	cd domain; echo "GENERATE definition mock with MOCKERY"; \
	mockery --output ../mock --all

test:
	@go test -v -race -timeout 10s -parallel 8 -shuffle=on -coverprofile=coverage.out -covermode=atomic ./...
	@go tool cover -func=coverage.out
