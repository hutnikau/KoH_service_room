.PHONY: build

build:
	sam build

build-RoomFunction:
	GOOS=linux GOARCH=amd64 go build -o main cmd/main.go
	cp ./main $(ARTIFACTS_DIR)/