BINARY_NAME=start

build:
	go build -o ${BINARY_NAME} ./cmd/

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}