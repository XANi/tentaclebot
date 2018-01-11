# generate version number
version=$(shell git describe --tags --long --always --dirty|sed 's/^v//')
binfile=main
goflags=-ldflags "-X main.version=$(version)"
DEPS=vendor
BINARIES=tentacle tentaclecli
all: $(BINARIES)  | glide.lock
	-@go fmt

%: %.go $(DEPS)
	go build $(goflags) -o bin/$@.arm64 $@.go
	go build $(goflags) -o bin/$@ $@.go

.PHONY: *.go

static: glide.lock vendor
	go build -ldflags "-X main.version=$(version) -extldflags \"-static\"" -o $(binfile).static $(binfile).go

arm:
	GOARCH=arm go build  -ldflags "-X main.version=$(version) -extldflags \"-static\"" -o $(binfile).arm $(binfile).go
	GOARCH=arm64 go build  -ldflags "-X main.version=$(version) -extldflags \"-static\"" -o $(binfile).arm64 $(binfile).go
clean:
	rm -rf vendor
	rm -rf _vendor
vendor: glide.lock
	glide install && touch vendor
glide.lock: glide.yaml
	glide update && touch glide.lock
glide.yaml:
version:
	@echo $(version)
