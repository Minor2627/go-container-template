FROM docker.io/library/golang:1.27.1-alpine3.24 AS builder

WORKDIR /app-src

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /binary-name

FROM scratch

COPY --from=builder /binary-name /binary-name

ENTRYPOINT ["/binary-name"]
