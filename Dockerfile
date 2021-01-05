# Multistage Socker file for Build
FROM golang:1.15-alpine AS builder
RUN apk --no-cache add git
WORKDIR /go/github.com/torre76/gochat  
RUN go get -d -v github.com/torre76/gochat
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o chatServer .

FROM alpine:latest
LABEL maintainer="gianluca.dallatorre@gmail.com"
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/github.com/torre76/gochat .
CMD ["./chatServer"]  
