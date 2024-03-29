ARG GO_VERSION="1.20"

FROM golang:${GO_VERSION}-alpine as builder

ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk --no-cache add ca-certificates
RUN update-ca-certificates

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies
# and only redownloading them in subsequent builds if they change
COPY . .
RUN go mod download && go mod verify
RUN go build -a -installsuffix -v -o cmp_lab_srv ./src/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /usr/src/app/cmp_lab_srv ./cmp_lab_srv
COPY --from=builder /usr/src/app/config_for_search ./config_for_search
COPY --from=builder /usr/src/app/static ./static
COPY --from=builder /usr/src/app/.env ./.env
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV GIN_MODE=release \
    HOST=http://localhost \
    PORT=8080 \
    REDIS_HOST=redis:6379 \
    REDIS_DB=0

EXPOSE 8080

CMD ["./cmp_lab_srv"]



