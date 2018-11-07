FROM golang:latest as builder
WORKDIR /go/src/github.com/alekssaul/demoapp-initializr/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o github_create_repository ./cmd/github_create_repository && \
    mkdir /app && \ 
    mv ./github_create_repository /app/github_create_repository

FROM alpine:latest
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app .
RUN apk update ; apk add --no-cache ca-certificates && update-ca-certificates
CMD /bin/sh
