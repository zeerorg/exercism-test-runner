FROM golang:1.11

RUN go get github.com/exercism/cli/exercism \
    && cd $GOPATH/src/github.com/exercism/cli \
    && go get -u github.com/golang/dep/cmd/dep \
    && dep ensure \
    && cd $GOPATH/src/github.com/exercism/cli/exercism \
    && go build -o /exercism main.go
