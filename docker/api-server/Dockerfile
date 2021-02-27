# Build
FROM golang:1.16.0-alpine3.12 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build ./cmd/api-server

# Release
FROM alpine:3.12.0

WORKDIR /app

COPY --from=builder /app/api-server .
COPY docker/api-server/entrypoint.sh .
COPY scripts/wait-for.sh .
RUN chmod +x wait-for.sh entrypoint.sh

ENV PATH="/app:$PATH"
ENTRYPOINT [ "entrypoint.sh" ]
CMD ["api-server"]
