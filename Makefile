TARGET_ALIENS=./dist/alien-invasion
TARGET_MAPGEN=./dist/alien-mapgen

.PHONY : test build-app build-mapgen
.DEFAULT_GOAL := _help

# Help: Each target starting with "_" will be ignored. Add description for target by adding "## <my help description>"
_help:
	@grep -E '^[0-9a-zA-Z][0-9a-zA-Z_-]+:.*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*## ?"}; {split($$1, t, ":"); printf "\033[1;34m%-30s\033[0m\t%s\n", t[1], $$2}'

build-app: ## Builds alien-invasion CLI app
	@go build -o ./$(TARGET_ALIENS) ./cmd/alien-invasion

build-mapgen: ## Builds alien-invasion map generation utility
	@go build -o ./$(TARGET_MAPGEN) ./cmd/mapgen

build-all: build-mapgen build-app

test:
	@go test -v ./...
