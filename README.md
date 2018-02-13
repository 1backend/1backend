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

## What's happening?

For announcements please follow [our twitter account](https://twitter.com/1backend).

## How can I install it?

A simple one click/single command installation is coming soon, but in the mean time:

### Frontend

If you have docker:

```sh
sudo docker run -p 4222:80 1backend/frontend
```

The above means on http://127.0.0.1:4222 you will have the 1backend app available, once the server is also running (see below).

Or if you want to hack on the Angular app:

```sh
npm install
npm start
```

### Backend

You need two containers running on your box: a MySQL one and a Redis one. The 1backen flavoured MySQL one has the table schemas loaded into it (available in [this](backend/schema.sql) file) already.

You can launch the complete sytem with 3 commands:

```sh
# Start mysql container. Comes with the database schema preloaded.
sudo docker run -e MYSQL_ROOT_PASSWORD=root -p=3306:3306 -d 1backend/mysql

# Start redis container.
sudo docker run -p=6379:6379 -d redis redis-server --appendonly yes

# Launch the 1backend server with the following command:
sudo docker run -e INTERNAL_IP=$(ip route get 8.8.8.8 | head -1 | cut -d' ' -f8) \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -p 8883:8883 1backend/server
```

The above does 3 things:

* passes the host internal network ip as an envar to the container
* mounts the docker socket

You can also mount a config file into the container under the path `/var/1backend-config.json`.
We haven't talked about configuration, so let's do it now:

### Configuration

The server container loads configuration from the location `/var/1backend-config.json`.
Details of the config parameters are
[here](https://github.com/1backend/1backend/blob/master/backend/config/config.go).

A very basic and working example of such file would be:

```
{
   "SiteUrl": "http://127.0.0.1",
   "Path": "/go/src/github.com/1backend/1backend/backend",
}
```

This is the default config file if you don't mount a config file to run the 1backend server docker container.

Of course, there are more in depth things to consider...

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
