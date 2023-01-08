start:
	go run ./cmd/server/main.go

build_windows:
	CGO_ENABLED=0 GOOS=windows go build -a -ldflags '-extldflags "-static"' -o ./bin/server.exe ./src/main.go

build_linux:
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o ./bin/server ./src/main.go

swagger:
	swag init -g cmd/server/main.go -o ./docs