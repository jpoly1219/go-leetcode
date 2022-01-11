FROM golang:1.17

WORKDIR /coderunner

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /binCoderunner

RUN curl -fsSL https://deb.nodesource.com/setup_16.x | bash

RUN apt install nodejs

EXPOSE 8091

CMD ["/binCoderunner"]