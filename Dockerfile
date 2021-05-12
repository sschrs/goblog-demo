FROM golang:alpine3.12
WORKDIR /go/src/goblog
COPY . .
CMD ["/go/src/goblog/goblog"]

