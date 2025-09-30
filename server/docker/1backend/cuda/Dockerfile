ARG BACKEND_VERSION="12.5.0"
ARG UBUNTU_VERSION="22.04"

FROM golang:1.24.5 as build

WORKDIR /app/server

COPY server/go.mod server/go.sum ./
COPY clients/go /app/clients/go
COPY sdk/go /app/sdk/go

RUN go mod download

COPY server/ .

RUN go build -o /app/main .

FROM nvidia/cuda:${BACKEND_VERSION}-base-ubuntu${UBUNTU_VERSION}

RUN apt-get update && apt-get install -y \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=build /app/main .

EXPOSE 11337

CMD ["./main"]
