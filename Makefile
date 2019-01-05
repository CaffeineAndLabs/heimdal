BINARY=heimdal
GOARCH=amd64

all: linux

linux:
	GOOS=linux GOARCH=${GOARCH} go build -o ${BINARY}-linux
