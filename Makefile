now:=$(shell date +'%Y-%m-%d_%T')
git_ver:=$(shell git rev-parse HEAD)
git_branch:=$(shell git branch | grep -E '^\*' | cut -f2 -d' ')

linters:
	aligo check ./...
	golangci-lint run ./...

deps:
	go mod tidy

clean-api:
	rm -f bin/api

build-api: clean-api deps linters
	cd cmd/api; go build -v \
		-ldflags="-linkmode external -extldflags -static -X main.gitVersion=${git_ver} -X main.buildTime=${now} -X main.gitBranch=${git_branch}"\
		-race -trimpath \
		-o ../../bin/api

build-released-api: clean-api deps linters
	cd cmd/api; go build -v \
		-ldflags="-s -w -linkmode external -extldflags -static -X main.gitVersion=${git_ver} -X main.buildTime=${now} -X main.gitBranch=${git_branch}"\
		-trimpath \
		-o ../../bin/api


