format:
    gofmt -s -w -l .

test: format
    go test ./internal/client


