FROM golang:1.20

WORKDIR server_compare_labs

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

EXPOSE 8080

RUN go build main.go

CMD [ "./main" ]
