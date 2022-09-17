FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd/main/*.go /app
COPY ./pkg/controllers/*.go /app/pkg/controllers/
COPY ./pkg/models/*.go /app/pkg/models/

RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD ["/docker-gs-ping"]