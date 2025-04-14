BINARY_NAME=start

build:
	go build -o ${BINARY_NAME} ./

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}
