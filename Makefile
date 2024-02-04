WEB = ${PWD}/frontend
PKG = pnpm # npm, pnpm, yarn, etc...
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

## Platforms
.PHONY: build-linux
build-linux:
	@printf "Building without the Wails CLI...\n\n"
	@./scripts/build

.PHONY: build-windows
build-windows:
	@printf "Building for windows with the Wails CLI...\n\n"
	@CGO_ENABLED=1 GOOS=windows wails build -platform windows
