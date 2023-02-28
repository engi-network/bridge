# Rebuild the source code only when needed
FROM golang:1.18-stretch AS builder
WORKDIR /build
COPY . .

RUN go build -o /bridge .

# Production image, copy all the files and run next
FROM debian:stretch-slim AS runner

# Install trusted CA certificates
RUN apt-get -y update && apt-get -y upgrade && apt-get install ca-certificates wget dnsutils vim -y

COPY --from=builder /bridge /bridge

EXPOSE 8001

ENTRYPOINT ["/bridge"]

CMD ["run"]
