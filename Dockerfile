FROM golang:alpine

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.io,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go env && go build -o app .

WORKDIR /dist

RUN cp /build/app .

EXPOSE 8888