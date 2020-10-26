#!/bin/sh
set -e
APP_NAME="gobuilder"
APP_VERSION="v1.0.1"
BUILD_TIME=$(date "+%F %T")
GIT_REVISION=$(git rev-parse --short HEAD)
GIT_BRANCH=$(git name-rev --name-only HEAD)
GO_VERSION=$(go version)

go build -ldflags "-s -X 'main.AppName=${APP_NAME}' \
			-X 'main.AppVersion=${APP_VERSION}' \
			-X 'main.BuildTime=${BUILD_TIME}' \
			-X 'main.GitRevision=${GIT_REVISION}' \
			-X 'main.GitBranch=${GIT_BRANCH}' \
			-X 'main.GoVersion=${GO_VERSION}'"
