

golang:
	docker build -t example-alpine -f Dockerfile.golang .

alpine:
	docker build -t example-alpine -f Dockerfile.alpine .

scratch:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main
	docker build -t example-scratch -f Dockerfile.scratch .
