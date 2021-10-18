
# Go编译C共享动态库

## 1. 编译并安装go标准库

```bash

# 注意go安装目录的读写权限
sudo chmod -R 0777 /usr/local/go

# 编译安装golang标准库为so
go install -buildmode=shared -linkshared std
```

## 2. 编译插件

```bash
go build -buildmode=c-shared -o test/libhello.so 
```