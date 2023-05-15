FROM golang:latest

RUN go version

WORKDIR /server

COPY . .

EXPOSE 8080

ENV GIN_MODE=release \
    HOST="http://localhost" \
    PORT=8080

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -v -o server src/main.go

CMD ["server"]
