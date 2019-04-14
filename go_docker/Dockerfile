# Using multistage build.
FROM golang:1.11.5 as builder

# in go docker, default GOPATH is /go
WORKDIR /go/src/
RUN go get -d -v github.com/go-redis/redis
COPY ./src/  .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o redisconsumer .

FROM alpine:3.9.2
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/redisconsumer .

CMD ["./redisconsumer"]
