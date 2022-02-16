FROM golang:1.17-alpine

ADD . /workdir

WORKDIR /workdir

RUN go mod download

RUN go build -o /server cmd/server/main.go

CMD [ "/server" ]
