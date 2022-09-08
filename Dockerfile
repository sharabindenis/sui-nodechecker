CMD ["/docker-gs-ping"]

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd/main/*.go /app

RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD ["/docker-gs-ping"]