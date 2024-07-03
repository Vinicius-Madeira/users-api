FROM golang:1.22 AS builder
WORKDIR /app
COPY src src
COPY docs docs
COPY go.mod go.mod
COPY go.sum go.sum
COPY init_dependencies.go init_dependencies.go
COPY main.go main.go

ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=LINUX GOARCH=amd64 \
GOOS=linux go build -o users-api

FROM golang:1.22-alpine3.19 AS runner
RUN adduser -D api-user
COPY --from=builder /app/users-api /app/users-api
RUN chown -R api-user:api-user /app
RUN chmod +x /app/users-api
EXPOSE 8080
USER api-user
CMD ["/app/users-api"]