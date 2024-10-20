FROM golang:1.23-alpine AS builder
WORKDIR /usr/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o inventory ./cmd/main.go

FROM scratch
COPY --from=builder /usr/src/app/inventory ./inventory
CMD ["./inventory"]
