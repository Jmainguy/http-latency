FROM golang:1.12.9
WORKDIR /go/src/app
ADD main.go .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["app"]
