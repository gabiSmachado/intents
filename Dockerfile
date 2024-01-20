FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY /smoapp/intent_receiver.go ./


RUN go get github.com/gabiSmachado/intents/database
RUN go get github.com/gabiSmachado/intents/datamodel
RUN go build -o /smoapp

EXPOSE 9090

CMD [ "/smoapp"]