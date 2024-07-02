FROM golang:1.22 AS builder
WORKDIR /app
COPY src src
COPY go.mod go.mod
COPY go.sum go.sum
COPY init_dependencies.go init_dependencies.go
COPY main.go main.go

ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=LINUX GOARCH=amd64 \
GOOS=linux go build -o go-web-app

FROM golang:1.22-alpine3.19 AS runner
RUN adduser -D goapp
COPY --from=builder /app/go-web-app /app/go-web-app
RUN chown -R goapp:goapp /app
RUN chmod +x /app/go-web-app
EXPOSE 8080
USER goapp
CMD ["/app/go-web-app"]