FROM golang:1.12-alpine
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY main.go .
RUN go build -o main

EXPOSE 8080

CMD ["./main"]
