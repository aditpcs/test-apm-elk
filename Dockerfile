FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/myapp .

FROM alpine:latest

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /home/appuser/app

COPY --chown=appuser:appgroup .env .

COPY --from=builder /go/bin/myapp /usr/local/bin/myapp

USER appuser

EXPOSE 8810

ENTRYPOINT ["myapp"]
