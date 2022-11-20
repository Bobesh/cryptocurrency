FROM golang:1.19-alpine AS builder

RUN apk add --no-cache git
RUN apk upgrade --update-cache --available && \
    apk add openssl && \
    rm -rf /var/cache/apk/*

WORKDIR /tmp/app

COPY server/go.mod .
COPY server/go.sum .
COPY generate_files.sh .

RUN chmod u+x generate_files.sh
RUN ./generate_files.sh
RUN ls

RUN go mod download

COPY server .

RUN go build -o ./out/server .

FROM alpine:latest
RUN apk add ca-certificates

COPY --from=builder /tmp/app/out/server /app/server
COPY --from=builder /tmp/app/files /files

CMD ["/app/server"]