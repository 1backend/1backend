FROM golang:alpine

RUN apk add --no-cache bash curl git openssh-client
RUN apk add --no-cache mysql-client

ENV VERSION "17.12.0-ce"
RUN curl -L -o /tmp/docker-$VERSION.tgz https://download.docker.com/linux/static/stable/x86_64/docker-$VERSION.tgz \
    && tar -xz -C /tmp -f /tmp/docker-$VERSION.tgz \
    && mv /tmp/docker/docker /usr/bin \
    && rm -rf /tmp/docker-$VERSION /tmp/docker

WORKDIR /go/src/github.com/1backend/1backend/backend
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["backend"]
