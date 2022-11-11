FROM golang:1.19.3-alpine3.16

WORKDIR /app
COPY go.mod ./
COPY *.go ./

RUN go build -o /echo-server

EXPOSE 8080

CMD [ "/echo-server" ]

