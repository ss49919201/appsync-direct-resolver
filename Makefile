build:
	GOOS=linux GOARCH=amd64 go build -o main ./...

zip: build
	zip main.zip main
