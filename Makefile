WEB = ${PWD}/frontend
PKG = npm # pnpm, yarn, etc...
CLI = wails # https://wails.io/docs/reference/cli

format:
	@printf "Formatting backend...\n"
	@go fmt ./...
	@go mod tidy
	@printf "Formatting frontend...\n"
	@cd $(WEB) && $(PKG) run format

update:
	@printf "Updating backend...\n"
	@go mod tidy
	@printf "Updating frontend...\n"
	@cd $(WEB) && $(PKG) install

## Production & Development
.PHONY: build
build:
	@printf "Building for production...\n\n"
	@$(CLI) build

.PHONY: dev
dev:
	@printf "Running in development mode...\n\n"
	@$(CLI) dev

.PHONY: build-linux
build-linux:
	@printf "Building for without the Wails CLI...\n\n"
	@./scripts/build
