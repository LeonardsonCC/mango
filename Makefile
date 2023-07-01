.PHONY: build

build:
	@echo Building as ./build/cli
	@go build -o ./build/cli ./cmd/cli

build_win:
	@echo Building as ./build/cli.exe
	@go build -o ./build/cli.exe ./cmd/cli
