FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY /smoapp/intent_receiver.go ./

EXPOSE 8585

CMD sleep 100000000