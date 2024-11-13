FROM golang:latest

WORKDIR /app

COPY . .

RUN mkdir -p /data/kntrlnt

RUN go mod download

RUN go build -o /kntrlnt

EXPOSE 8001

CMD ["/kntrlnt"]