FROM golang:1.17.5-alpine

ENV GOPATH /go
ENV PATH=$PATH:$GOPATH/src

# 以下、Docker run 用の設定
ENV PATH=$PATH:$GOPATH/src/github.com/go-es-testcode
WORKDIR $GOPATH/src/github.com/go-es-testcode
COPY /src $GOPATH/src/github.com/go-es-testcode/src

RUN go mod init go-es-testcode
RUN go mod tidy

RUN GOOS=linux go build -o go-es-testcode ./src
EXPOSE 8090

ENTRYPOINT ["/go/src/github.com/go-es-testcode/go-es-testcode"]
