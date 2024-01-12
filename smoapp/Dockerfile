FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./

RUN go build -o /intentreceiver

EXPOSE 8585

CMD [ "/intentreceiver"]