BUILD=build
TEST=test

.setup:
	export GOPATH=${PWD} 

test: .setup
	go test -o ${TEST}/giks -json

all: .setup
	go build -o ${BUILD}/giks main.go
