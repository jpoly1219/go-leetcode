FROM golang:1.17

WORKDIR /coderunner

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /binCoderunner

EXPOSE 8091

CMD ["/binCoderunner"]