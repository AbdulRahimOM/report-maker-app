.PHONY: \
    run \
    build \
    build-and-run \
    open-app

run:
	go run main.go

build:
	go build -o ./reporter-app main.go

build-and-run:
	go build -o ./reporter-app main.go; ./reporter-app

open-app:
	./reporter-app