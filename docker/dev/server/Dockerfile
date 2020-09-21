FROM golang:1.15
ARG WORKDIR
WORKDIR $WORKDIR
ENV GO111MODULE=on
ENV GOPATH=''
COPY ./server/ ./

CMD go mod download && go run .
