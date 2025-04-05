format:
    gofmt -s -w -l .

test: format
    go test ./internal/client

build: format
    go build cmd/catacloud.go

