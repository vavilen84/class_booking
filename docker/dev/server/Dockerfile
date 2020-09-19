FROM golang:1.15

WORKDIR /usr/server/
ENV GO111MODULE=on
ENV GOPATH=''
COPY ./server/ ./

CMD go mod download && go run .
