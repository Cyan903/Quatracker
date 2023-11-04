APP = "QuaverBuddy" # App name
WEB = ${PWD}/frontend
PKG = npm # pnpm, yarn, etc...

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
	@wails build

.PHONY: dev
dev:
	@printf "Running in development mode...\n\n"
	@wails dev
