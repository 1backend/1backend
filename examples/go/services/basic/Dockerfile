FROM golang:1.23.2 as build

WORKDIR /app/services/basic

COPY services/basic/go.mod services/basic/go.sum ./
COPY sdk/go /app/sdk/go

RUN go mod download

COPY services/basic/ .

RUN go build -o /app/main .

FROM ubuntu:latest
RUN apt-get update && apt-get install -y ca-certificates

WORKDIR /app
COPY --from=build /app/main .

EXPOSE 9111

CMD ["./main"]
