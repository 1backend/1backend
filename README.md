# 1Backend

1Backend is a PaaS designed to make deploying, running and maintaining services
cheap and fun.

![1backend service screenshot](https://raw.githubusercontent.com/1backend/1backend/master/1b.png)

## Why?

The idea of building this came when we realised we have a bunch of services
running all over the place and it costs a lot of money and time to maintain
them, and we couldn't find an open source solution to host them efficiently.

We've also spent more than 5 years building and using microservices at various
companies. We have seen very many ways of it going wrong.

This project is an excercise in daydreaming about how could it possibly work
well.

Perhaps, most importantly, it is also a lot of fun to build this.

## What's the state of the project?

It's very early. It's already usable as it is if you want to host it
yourself, but our version available under https://1backend.com is only recommended for the adventurous.

It's fully functional however, and you can use it. All the quota you buy and use now will be given to you again once the testing ends. Beware of storing sensitive data there, since we haven't focus on plugging all security holes yet.

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
Details of the config parameters are [here](https://github.com/1backend/1backend/blob/master/backend/config/config.go).

#### Api generation

Given services and endpoints have type informations saved with them (can be set on the UI of each service), we generate client APIs for each service in a number of languages.

To do this we need integration with github.

#### GitHub

To get the github integration working you need two things:
- a user who can create repositories in your organisation (only possible through Github's HTTP API)
- an SSH key on the machine and added to a users' Github account who can commit into said repositories (`ssh-keygen -t rsa -b 4096 -C "your_email@example.com"; eval "$(ssh-agent -s)"; ssh-add ~/.ssh/id_rsa`). This user must be the one to who runs the server - likely root, since root is needed to access docker, unless you configure your machine otherwise.