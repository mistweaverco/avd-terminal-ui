build-linux-32:
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -ldflags "-s -w" -o dist/avd-linux-386 src/*.go
build-linux-64:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o dist/avd-linux src/*.go
build-windows-32:
	GOOS=windows GOARCH=386 CGO_ENABLED=0 go build -ldflags "-s -w" -o dist/avd-386.exe src/*.go
build-windows-64:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o dist/avd.exe src/*.go
build-macos-arm64:
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "-s -w" -o dist/avd-mac-arm64 src/*.go

build: build-linux-32 build-linux-64 build-windows-32 build-windows-64 build-macos-arm64

build-and-install: build-linux-64
	sudo cp dist/avd-linux /usr/local/bin/avd

run:
	go run ./src/*.go

