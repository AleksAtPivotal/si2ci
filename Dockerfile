FROM golang:latest as builder
WORKDIR /go/src/github.com/alekssaul/si2ci
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o si2ci ./ && \
    mkdir /app && \ 
    mv ./si2ci /app/si2ci

FROM alpine:latest
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app .
RUN apk update ; apk add --no-cache ca-certificates && update-ca-certificates
CMD /bin/sh
