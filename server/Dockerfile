FROM golang:1.20

ENV GIN_MODE=release
ENV HOST=http://localhost
ENV PORT=8080
ENV CGO_ENABLED 0
ENV GOOS linux
ENV REDIS_HOST=redis:6379
ENV REDIS_DB=0

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./src/main.go

EXPOSE 8080
CMD ["app"]