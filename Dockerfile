FROM  golang:1.23.2-alpine3.18 as builder

RUN apk add --no-cache make ca-certificates gcc musl-dev linux-headers git jq bash

COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum

WORKDIR /app

RUN go mod download

# build per_apollo with the shared go.mod & go.sum files
COPY . /app/per

WORKDIR /app/per

RUN make build

FROM alpine:3.18

COPY --from=builder /app/per/per_apollo /usr/local/bin
COPY --from=builder /app/per/config-demo.yaml /app/per/apollo.yaml
COPY --from=builder /app/per/migrations /app/per/migrations

ENV INDEXER_MIGRATIONS_DIR="/app/per/migrations"
WORKDIR /app/per
