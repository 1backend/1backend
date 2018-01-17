<p align="center">
  <img width="200px" src="https://cdn.rawgit.com/1backend/1backend/master/src/assets/logos/trex.svg" />
<p>

# 1Backend [![circleci status](https://circleci.com/gh/1backend/1backend.svg?style=shield)](https://circleci.com/gh/1backend/1backend) [![go report](https://goreportcard.com/badge/github.com/1backend/1backend)](https://goreportcard.com/report/github.com/1backend/1backend) [![go coverage](https://codecov.io/gh/1backend/1backend/branch/master/graph/badge.svg)](https://codecov.io/gh/1backend/1backend/branch/master)

1Backend is a platform designed to make deploying, running and maintaining lambda functions/microservices easy.

![1backend service screenshot](https://raw.githubusercontent.com/1backend/1backend/master/1b.png)

It enables you to launch a new live app in seconds - after choosing your tech stack (e.g. Go with access to an SQL database) you get an empty [app](docs/services.md) which is already live and callable from the outside (through HTTP).

You just have to plug in your own code (no, you don't have to write code in the browser). It's even pre-connected to your database and other infrastructure elements of your choosing.

With the help of a simple [DSL](docs/types.md), you can even define language-agnostic types and APIs for your services, which are used to generate type-safe clients. These clients are then published on each language's package manager (e.g. NPM).

## Documentation

You can find the documentation [here](docs).

## What's supported?

### Languages

* Go
* JavaScript (Node.js)
* TypeScript

### Infrastructure

* MySQL

### Client library generation

* Go
* Angular: an NPM installable Angular service is generated and published on NPM
  if your service has type definitions.

## What's the state of the project?

It's very early.

## What's the tech used?

Angular 2, Go, MySQL, Redis.

## How can I install it?

### Frontend

In the project root:

```sh
npm install
npm start
```

### Backend

```sh
cd backend

sudo docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=root -p=3306:3306 -d mysql
# in case you don't have the mysql client installed:
sudo apt-get -y install mysql-client
mysql -h localhost -P 3306 --protocol=tcp -u root -proot < all.sql

sudo docker run --name some-redis -p=6379:6379 -v /var/redis:/data -d redis redis-server --appendonly yes

# this assumes you have go installed
rm main; go build main.go; sudo ./main
```

The server loads configuration from the location `/var/1backend-config.json`.
Details of the config parameters are
[here](https://github.com/1backend/1backend/blob/master/backend/config/config.go).

#### API generation

Given services and endpoints have type information saved with them (can be set
in the UI of each service), we generate client APIs for each service in a number
of languages.

To do this we need integration with GitHub.

#### GitHub

To get the GitHub integration working, you need two things:

* a user who can create repositories in your organisation (only possible through
  GitHub's HTTP API)
* an SSH key on the machine and added to a users' GitHub account who can commit
  into said repositories (`ssh-keygen -t rsa -b 4096 -C "your_email@example.com"; eval "$(ssh-agent -s)"; ssh-add ~/.ssh/id_rsa`).
  This user must be the one who runs the server - likely root, since root is
  needed to access docker, unless you configure your machine otherwise.
