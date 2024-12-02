FROM golang:1.23-alpine
RUN apk update && apk add git
RUN go install github.com/air-verse/air@latest
RUN mkdir -p /go/app
WORKDIR /go/app

CMD ["air", "-c", ".air.toml"]