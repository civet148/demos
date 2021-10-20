
.PHONY: linux

linux:
	go build -buildmode=c-shared -o lib/linux/libhello.so

arm64:arm64-gcc
	CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc GOARCH=arm64 go build -buildmode=c-shared -o lib/arm64/libhello.so

arm64-gcc:
	 sudo apt update && sudo apt install -y gcc-aarch64-linux-gnu
