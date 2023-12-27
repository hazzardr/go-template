PROJECT_NAME := "go-template"
EXEC_NAME := template

.PHONY: help ## print this
help:
	@echo ""
	@echo "$(PROJECT_NAME) Development CLI"
	@echo ""
	@echo "Usage:"
	@echo "  make <command>"
	@echo ""
	@echo "Commands:"
	@grep '^.PHONY: ' Makefile | sed 's/.PHONY: //' | awk '{split($$0,a," ## "); printf "  \033[34m%0-10s\033[0m %s\n", a[1], a[2]}'

.PHONY: run ## Run the project
run:
	$(MAKE) generate
	$(MAKE) build
	@go run cmd/main.go

.PHONY: doctor ## checks if local environment is ready for development
doctor:
	@echo "Checking local environment..."
	@if ! command -v go &> /dev/null; then \
		echo "`go` is not installed. Please install it first."; \
		exit 1; \
	fi
	@if [[ ! ":$$PATH:" == *":$$HOME/go/bin:"* ]]; then \
		echo "GOPATH/bin is not in PATH. Please add it to your PATH variable."; \
		exit 1; \
	fi
	@if ! command -v cobra-cli &> /dev/null; then \
		echo "Cobra-cli is not installed. Please run 'make deps'."; \
		exit 1; \
	fi
	@if ! command -v sqlc &> /dev/null; then \
		echo "`sqlc` is not installed. Please run 'make deps'."; \
		exit 1; \
	fi

	@if ! command -v docker &> /dev/null; then \
		echo "`docker` is not installed. Please install it first."; \
		exit 1; \
	fi
	@echo "Local environment OK"


.PHONY: deps ## install dependencies used for development
deps:
	@echo "Installing dependencies..."
	@go install github.com/spf13/cobra-cli@latest
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

.PHONY: build ## build the project
build:
	@echo "Building..."
	@go build -o $(EXEC_NAME) ./cmd/main.go

.PHONY: generate ## generate server and database code
generate:
	@echo "Generating database models..."
	@sqlc generate

.PHONY: docker ## build docker image
docker:
	@echo "Building docker image..."
	@docker build -t $(PROJECT_NAME) .