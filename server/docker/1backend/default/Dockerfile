FROM golang:1.23.2 as build

WORKDIR /app/server

COPY server/go.mod server/go.sum ./
COPY clients/go /app/clients/go
COPY sdk/go /app/sdk/go

RUN go mod download

COPY server/ .

RUN go build -o /app/main .

FROM ubuntu:latest
RUN apt-get update && apt-get install -y ca-certificates

WORKDIR /app
COPY --from=build /app/main .

EXPOSE 58231

CMD ["./main"]
