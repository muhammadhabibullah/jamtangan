
.PHONY: config mock

config:
	@cp -n config/config.example.json config/config.json

dependency:
	@docker-compose up -d

migrate:
	@go run main.go -command=migrate

serve:
	@go run main.go

swag:
	@echo "> Generate Swagger Docs"
	@if ! command -v swag &> /dev/null; then go install github.com/swaggo/swag/cmd/swag@v1.8.4; fi
	@swag init -o handler/http/docs --ot json

mock:
	@if ! command -v mockery &> /dev/null; then go install github.com/vektra/mockery/v2@v2.10.0; fi
	cd domain; echo "GENERATE definition mock with MOCKERY"; \
	mockery --output ../mock --all

test:
	@go test -v -race -timeout 10s -parallel 8 -shuffle=on -coverprofile=coverage.out -covermode=atomic ./...
	@go tool cover -func=coverage.out
