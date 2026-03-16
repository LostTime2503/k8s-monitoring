FROM golang:1.26 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp main.go

FROM scratch
COPY --from=builder /app/myapp /myapp
EXPOSE 8080
ENTRYPOINT ["/myapp"]