
.PHONY: linux

linux:
	go build -buildmode=c-shared -o lib/linux/libhello.so

arm64:arm64-gcc
	CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc GOARCH=arm64 go build -buildmode=c-shared -o lib/arm64/libhello.so

arm64-gcc:
	 sudo apt update && sudo apt install -y gcc-aarch64-linux-gnu

android_arm64:
	GOARCH=arm64 GOOS=android CGO_ENABLED=1 CC=/usr/local/android-ndk-r20b/toolchains/llvm/prebuilt/linux-x86_64/bin/aarch64-linux-android21-clang go build -buildmode=c-shared -o lib/android_arm64/libhello.so

android_arm:
	GOARCH=arm GOOS=android CGO_ENABLED=1 CC=/usr/local/android-ndk-r20b/toolchains/llvm/prebuilt/linux-x86_64/bin/armv7a-linux-androideabi21-clang go build -buildmode=c-shared -o lib/android_arm/libhello.so