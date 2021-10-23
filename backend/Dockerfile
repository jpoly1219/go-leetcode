FROM golang:1.17

WORKDIR /backend

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /binBackend

EXPOSE 8090

CMD ["/binBackend"]