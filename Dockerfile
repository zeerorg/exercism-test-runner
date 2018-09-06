FROM golang:1.8

WORKDIR /go/src/app

RUN go get -d -v github.com/gorilla/mux
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

#CMD ["./app"]

FROM docker:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=0 /go/src/app/ .

CMD ["/go/src/app/dockerfiles/dind/build_test_images.sh", "&&", "./app"]
