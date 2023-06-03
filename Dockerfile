FROM golang:1.20.4-alpine

RUN apk update && apk add git

ENV GO111MODULE=on

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH/src/backend-shopbee

COPY . .

# RUN go mod init shopbee

WORKDIR $GOPATH/src/backend-shopbee/cmd/pro
RUN GOOS=linux go build -o app

ENTRYPOINT ["./app"]

EXPOSE 3000