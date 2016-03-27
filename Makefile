.PHONY: all clean

all: format test build

test:
	go test -v . 

format:
	gofmt -w ./core ./handler ./logger ./middleware ./models ./util ./main.go

build:
	mkdir -p builds
	# 设置交叉编译参数:
	# GOOS为目标编译系统, mac os则为 "darwin", window系列则为 "windows"
	# 生成二进制执行文件 akbs , 如在windows下则为 akbs.exe
	GOOS="linux" GOARCH="amd64" go build -v -o builds/akbs ./main.go && cp -rf config builds/config

clean:
	go clean -i