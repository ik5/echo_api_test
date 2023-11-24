now:=$(shell date +'%Y-%m-%d_%T')
gitver:=$(shell git rev-parse HEAD)


linters:
	aligo check ./...
	golangci-lint run ./...

deps:
	go mod tidy
