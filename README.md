# 1Backend

1Backend is a platform designed to make deploying, running and maintaining
lambda functions/microservices easy and quick.

![1backend service screenshot](https://raw.githubusercontent.com/1backend/1backend/master/1b.png)

After choosing your tech stack (eg. go with access to an SQL database) you get
an empty [app](docs/services.md) which is already live and callable from the outside (through HTTP).

You just have to plug in your own code (no, you don't have to write code in the browser, although you can). I't even preconnected to your database and other
infratructure elements of your choosing.

With the help of a small [DSL](docs/types.md) you can even define languages agnostic types and APIs for your services which is used to generate type safe clients. These clients are then published on each language's package manager (eg. NPM).

You can find the documentation [here](docs).

## What's supported?

### Languages

* Go
* Javascript (Node.js)
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

In project root:

```sh
npm install
ng serve
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

#### Api generation

Given services and endpoints have type informations saved with them (can be set
on the UI of each service), we generate client APIs for each service in a number
of languages.

To do this we need integration with github.

#### GitHub

To get the github integration working you need two things:

* a user who can create repositories in your organisation (only possible through
  Github's HTTP API)
* an SSH key on the machine and added to a users' Github account who can commit
  into said repositories (`ssh-keygen -t rsa -b 4096 -C "your_email@example.com"; eval "$(ssh-agent -s)"; ssh-add ~/.ssh/id_rsa`).
  This user must be the one to who runs the server - likely root, since root is
  needed to access docker, unless you configure your machine otherwise.
