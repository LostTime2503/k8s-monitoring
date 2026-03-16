FROM golang:1.26 as builder

WORKDIR /app
COPY go.mod ./
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o traffic-gen main.go

FROM scratch
COPY --from=builder /app/traffic-gen /traffic-gen
# Needed for HTTPS requests certificates if any, but scratch is empty. 
# For http internal cluster traffic, scratch is usually fine.
# If DNS is needed, we might need /etc/resolv.conf or a Distroless base.
# Let's use alpine/distroless for better compatibility if available, or just stick to scratch for simplicity as request by user "small binary".
# Actually, net/http needs /etc/ssl/certs/ca-certificates.crt for https. 
# But we are likely calling http://myapp-service.
ENTRYPOINT ["/traffic-gen"]