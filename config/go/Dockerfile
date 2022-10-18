FROM golang:1.17.5-alpine
MAINTAINER kemper0530

ENV GOPATH /go
ENV PATH=$PATH:$GOPATH/src

# 以下、Docker run 用の設定
ENV PATH=$PATH:$GOPATH/src/go-es-testcode
WORKDIR $GOPATH/src/go-es-testcode
COPY /src $GOPATH/src/go-es-testcode

RUN go mod init go-es-testcode
RUN go mod tidy

RUN GOOS=linux go build -o go-es-testcode
EXPOSE 8090

ENTRYPOINT ["/go/src/github.com/kemper0530/go-es-testcode/go-es-testcode"]
