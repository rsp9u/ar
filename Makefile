CC = gcc
BINARY = ar

GOOPTIONS=GOOS=$(GOOS) GOARCH=$(GOARCH) CC=${CC} CGO_ENABLED=0

bin/${BINARY}: *.go cmd/main.go
	${GOOPTIONS} go build -o bin/${BINARY} cmd/main.go
