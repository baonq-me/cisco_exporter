FROM golang:1.23-alpine3.20 AS builder
ADD . /go/cisco_exporter/
WORKDIR /go/cisco_exporter

# Set CGO_ENABLED=0 to reduce binary file size
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/cisco_exporter


FROM golang:1.23-alpine3.20

# RUN apk --no-cache add ca-certificates

RUN addgroup -S app && adduser -S app -s /bin/bash -h /app -G app
USER app

ENV CMD_FLAGS=""

WORKDIR /app

COPY --from=builder /go/bin/cisco_exporter .
CMD ["/app/cisco_exporter", "-ssh.keyfile=/app/ssh.key", "-config.file=/app/config.yml", "$CMD_FLAGS"]

EXPOSE 9362
