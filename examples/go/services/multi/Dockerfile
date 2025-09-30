FROM golang:1.24.5 as build

WORKDIR /app/services/multi

COPY services/multi/go.mod services/multi/go.sum ./
COPY sdk/go /app/sdk/go

RUN go mod download

COPY services/multi/ .

RUN go build -o /app/main .

FROM ubuntu:latest
RUN apt-get update && apt-get install -y ca-certificates

WORKDIR /app
COPY --from=build /app/main .

EXPOSE 9211

CMD ["./main"]
