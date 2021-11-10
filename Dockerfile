FROM golang:1.17
WORKDIR /go/src/app
COPY . .
RUN go mod init
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["app"]