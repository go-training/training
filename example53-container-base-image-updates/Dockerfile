# syntax=docker/dockerfile:1

##
## STEP 1 - BUILD
##

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.20-alpine AS build

# create a working directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY go.mod ./
COPY go.sum ./

# download Go modules and dependencies
RUN go mod download

# copy directory files i.e all files ending with .go
COPY *.go ./

# compile application
RUN CGO_ENABLED=0 go build -o /server -ldflags="-w -s"

##
## STEP 2 - DEPLOY
##
FROM scratch

WORKDIR /

COPY --from=build /server /server

EXPOSE 8080

ENTRYPOINT ["/server"]
