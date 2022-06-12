.PHONY: build
build:
	go build -o bin/shmya-eda *.go || exit 1