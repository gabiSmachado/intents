FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY /broker/intent_broker.go ./

RUN go get github.com/gabiSmachado/intents/datamodel
RUN go build -o /intent_broker

EXPOSE 8585

CMD [ "/intent_broker"]