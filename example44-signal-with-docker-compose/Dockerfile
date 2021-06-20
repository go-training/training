FROM golang:1.16-alpine

COPY main.go /app/
COPY go.mod /app/
RUN go build -o /app/main /app/main.go

WORKDIR "/app"

CMD ["/app/main"]
