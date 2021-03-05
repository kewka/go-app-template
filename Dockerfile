# Go builder
FROM golang:1.16.0-alpine3.12 as go-builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build ./cmd/app

# Web builder
FROM mhart/alpine-node:14 as web-builder
WORKDIR /app
COPY web/package.json web/yarn.lock ./
RUN yarn
COPY web/src src
COPY web/public public
RUN yarn build

FROM alpine:3.12.0
WORKDIR /app
COPY --from=go-builder /app/app app
COPY --from=web-builder /app/build web/build
COPY docker-entrypoint.sh .
COPY scripts/wait-for.sh .
RUN chmod +x wait-for.sh docker-entrypoint.sh
ENTRYPOINT [ "./docker-entrypoint.sh" ]
CMD ["./app"]
