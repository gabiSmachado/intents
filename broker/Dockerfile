FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./

RUN go build -o /intent_broker

EXPOSE 9555

CMD [ "/intent_broker"]