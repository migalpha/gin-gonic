FROM golang:1.16-alpine AS builder

WORKDIR $GOPATH/src/gin-gonic

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo main.go

FROM alpine 
RUN apk add ca-certificates

COPY --from=builder /go/src/gin-gonic ./
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /zoneinfo.zip

ENV ZONEINFO=/zoneinfo.zip
