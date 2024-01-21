FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./

RUN go build -o /smoapp

EXPOSE 3000

CMD [ "/smoapp"]