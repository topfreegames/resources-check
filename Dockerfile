FROM golang:1.10-alpine

WORKDIR $GOPATH/src/github.com/topfreegames/resources-check
COPY . .
RUN apk update && \
    apk add git && \
    go get -u github.com/golang/dep/cmd/dep && \
    dep ensure
RUN go build -o resources-check main.go && \
    cp resources-check /bin/resources-check && \
    cp config/local.yaml /etc/resources-check.yaml && \
    rm -r $GOPATH/src/github.com/topfreegames/resources-check

CMD ["/bin/resources-check", "start", "--json", "--incluster", "-c", "/etc/resources-check.yaml"]
