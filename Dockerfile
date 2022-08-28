FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o hello cmd/main.go

FROM alpine

WORKDIR /build
COPY configs/config.yml ./configs/config.yml
COPY .env .
COPY --from=builder /build/hello /build/hello

CMD ["./hello"]