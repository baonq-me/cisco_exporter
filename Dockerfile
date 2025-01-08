FROM golang:1.23-alpine3.20 AS builder
ADD . /go/cisco_exporter/
WORKDIR /go/cisco_exporter
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/cisco_exporter


FROM golang:1.23-alpine3.20

RUN apk --no-cache add ca-certificates

WORKDIR /app

ENV CMD_FLAGS ""

COPY --from=builder /go/bin/cisco_exporter .
CMD ./cisco_exporter -ssh.keyfile=/app/ssh.key -config.file=/app/config.yml $CMD_FLAGS

EXPOSE 9362
