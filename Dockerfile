# https://iximiuz.com/en/posts/containers-distroless-images/

FROM golang:1.22.1-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/caddy-demo

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /app/bin/caddy-demo /app/bin/caddy-demo

ENTRYPOINT ["/app/bin/caddy-demo"]

