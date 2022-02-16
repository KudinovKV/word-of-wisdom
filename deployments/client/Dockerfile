FROM golang:1.17-alpine

ADD . /workdir

WORKDIR /workdir

RUN go mod download

RUN go build -o /client cmd/client/main.go

CMD [ "/client" ]
