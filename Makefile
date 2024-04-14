.PHONY: setup
setup:
	go install golang.org/x/tools/cmd/deadcode@latest

detect:
	deadcode ./...
