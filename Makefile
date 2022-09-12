BINARY_NAME=main.out

build:
	go build -o ${BINARY_NAME} server.go

run:
	go build -o ${BINARY_NAME} server.go
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

deps:
	go mod download
