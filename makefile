BUILD_ENV=GO111MODULE=on GOFLAGS=-mod=vendor CGO_ENABLED=0 GOOS=linux GOARCH=amd64
PROJECT_NAME=go-gomock
GO_FILES=$(shell go list ./... | grep -v /vendor/)
BUILD_TAG=$(shell git rev-parse --short=8 HEAD)

.SILENT:

build:
	$(BUILD_ENV) go build -o $(PROJECT_NAME)-$(BUILD_TAG).bin