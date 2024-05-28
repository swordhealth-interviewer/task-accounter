# syntax=docker/dockerfile:1

# build stage
FROM golang:1.22-alpine

WORKDIR /app

ADD . /app/

RUN go build -o api cmd/main.go

# final stage
FROM golang:1.22-alpine

WORKDIR /app

COPY --from=0 /app/api ./

EXPOSE 8080

CMD ["./api"]