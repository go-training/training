

golang:
	docker build -t example-golang -f Dockerfile.golang .

alpine:
	docker build -t example-$@ -f Dockerfile.$@ .

scratch:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main
	docker build -t example-$@ -f Dockerfile.$@ .

certs:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main
	docker build -t example-$@ -f Dockerfile.$@ .
