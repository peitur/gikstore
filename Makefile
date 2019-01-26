BUILD=build
TEST=test

.setup:
	export GOPATH=${PWD} 

main: all

test: .setup
	go test -v -o ${TEST}/giks -json

clean :
	go clean

all: .setup
	go build -v -o ${BUILD}/giks main.go


