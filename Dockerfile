FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/app ./cmd/app

FROM scratch

COPY --from=builder /bin/app /bin/app

COPY --from=builder /app/config.yml /config.yml

EXPOSE 8080

ENTRYPOINT ["/bin/app"]

