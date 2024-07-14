.PHONY: \
    run \
    build \
    build-and-run \
    windows-build-64 \
    windows-build-32 \
    linux-build-64 \
    linux-build-32 \
    freebsd-build-64 \
    freebsd-build-32 \
    mac-build \
    open-app

run:
	go run main.go

build: #mac-build
	go build -o ./reporter-app main.go

build-downloads:
	go build -o /Users/abdulrahimom/Downloads/reporter-app main.go
build-and-run: #mac-build
	go build -o ./reporter-app main.go; ./reporter-app

windows-build-64:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./builds/reporter-app-64bit.exe main.go

windows-build-32:
	GOOS=windows GOARCH=386 go build -o ./builds/reporter-app-32bit.exe main.go

linux-build-64:
	GOOS=linux GOARCH=amd64 go build -o ./builds/reporter-app-linux-64bit main.go

linux-build-32:
	GOOS=linux GOARCH=386 go build -o ./builds/reporter-app-linux-32bit main.go

freebsd-build-64:
	GOOS=freebsd GOARCH=amd64 go build -o ./builds/reporter-app-freebsd-64bit main.go

freebsd-build-32:
	GOOS=freebsd GOARCH=386 go build -o ./builds/reporter-app-freebsd-32bit main.go

mac-build:
	GOOS=darwin GOARCH=amd64 go build -o ./builds/reporter-app-mac main.go





open-app:
	./reporter-app