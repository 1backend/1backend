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

In the project root:

```sh
npm install
npm start
```

### Backend

You need two containers running on your box: a MySQL one and a Redis one. The MySQL one also needs the table schema loaded into it (available in [this](backend/all.sql) file).

Here is a short help:

```sh
cd backend

# start mysql container
sudo docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=root -p=3306:3306 -d mysql

# install mysql client so we can load the tables into the mysql container
sudo apt-get -y install mysql-client

# use the mysql client to load the schema
mysql -h localhost -P 3306 --protocol=tcp -u root -proot < all.sql

# start redis container
sudo docker run --name some-redis -p=6379:6379 -v /var/redis:/data -d redis redis-server --appendonly yes
```

After this is done, you can launch the 1backend server with the following command:

```sh
sudo docker run -e INTERNAL_IP=$(ip route get 8.8.8.8 | head -1 | cut -d' ' -f8) -v /var/run/docker.sock:/var/run/docker.sock -v /var/1backend-config.json:/var/1backend-config.json
-p 8883:8883 1backend/server
```

The above does 3 things:

* passes the host internal network ip as an envar to the container
* mounts the docker socket
* mounts the 1backend config file

We haven't talked about the latter, so let's do it now:

### Configuration

The server loads configuration from the location `/var/1backend-config.json`.
Details of the config parameters are
[here](https://github.com/1backend/1backend/blob/master/backend/config/config.go).

A very basic and working example of such file would be:

```
{
   "SiteUrl": "http://127.0.0.1",
   "Path": "/go/src/github.com/1backend/1backend/backend",
}
```

Such a minimal config file is enough to run the 1backend server docker container.

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
