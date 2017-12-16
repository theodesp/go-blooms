# Project name
PROJECT=go_blooms

# Check for syntax errors
.PHONY: vet
vet:
	GOPATH=$(GOPATH) go vet .

.PHONY: format
format:
	@find . -type f -name "*.go*" -print0 | xargs -0 gofmt -s -w

.PHONY: debs
debs:
	GOPATH=$(GOPATH) go get ./...
	GOPATH=$(GOPATH) go get -u github.com/spaolacci/murmur3
	GOPATH=$(GOPATH) go get -u gopkg.in/check.v1

.PHONY: test
test:
	GOPATH=$(GOPATH) go test

# Clean junk
.PHONY: clean
clean:
	GOPATH=$(GOPATH) go clean ./...