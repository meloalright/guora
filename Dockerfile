FROM golang:1.13-alpine AS builder
WORKDIR /go/src/github.com/meloalright/guora
ENV CC=gcc
COPY . .
RUN apk add --no-cache gcc musl-dev \
    && go build ./cmd/guora && mv guora /go/bin
###############
FROM alpine:3.6
COPY --from=builder /go/bin/guora /usr/local/bin
COPY --from=builder /go/src/github.com/meloalright/guora /guora
COPY configuration.example.yaml /etc/guora/configuration.yaml
WORKDIR /guora
CMD "guora" "-init"
