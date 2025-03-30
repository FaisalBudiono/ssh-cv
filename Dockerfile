FROM golang:1.23.3 AS base

FROM base AS deps
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download -x

FROM deps AS builder-app
COPY . .
RUN CGO_ENABLED=0 go build -o app ./main.go

FROM alpine:3.10 AS api
USER 1000
WORKDIR /app
RUN mkdir logs
COPY --from=base /usr/local/go/lib/time/zoneinfo.zip /
ENV ZONEINFO=/zoneinfo.zip
COPY --from=builder-app /app/app /app/app
CMD ["/app/app"]
