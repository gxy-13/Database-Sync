.PHONY: all build run gotool clean help

BINARY="DBS"

all: gotool build

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./build/${BINARY}

build-Windows32:
	GOOS=windows GOARCH=386 go build -o ./build/${BINARY}

build-Windows64:
	GOOS=windows GOARCH=amd64 go build -o ./build/${BINARY}

run:
	@go run ./main.go

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ./build/${BINARY} ]; then rm ./build/${BINARY}; fi

help:
	@echo "make - 格式化 Go 代码， 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码， 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"