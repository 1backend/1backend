# 1Backend

1Backend is a PaaS designed to make deploying, running and maintaining services cheap and fun.

![1backend service screenshot](https://raw.githubusercontent.com/1backend/1backend/master/1b.png)

#### Why?

The idea of building this came when we realised we have a bunch of services running all over the place and it costs a lot of money and time to maintain them, and we couldn't find an open source solution to host them efficiently.

We've also spent more than 5 years building and using microservices at various companies.
We have seen very many ways of it going wrong.

This project is an excercise in daydreaming about how could it possibly work well.

Perhaps, most importantly, it is also a lot of fun to build this.

#### What's the state of the project?

It's extremely early. It's already usable as it is if you want to host it yourself, but don't trust our version yet.

#### What's the tech used?

Angular 2, Go, MySQL, Redis.

#### How can I install it?

##### Frontend

In project root

```sh
npm install
ng serve
```
##### Backend

```sh
cd backend

sudo docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=root -p=3306:3306 -d mysql
# if you don't have the mysql client installed then:
# sudo apt-get install mysql-client
mysql -h localhost -P 3306 --protocol=tcp -u root -proot < all.sql

docker run --name some-redis -p=6379:6379 -v /var/redis:/data -d redis redis-server --appendonly yes

# this assumes you have go installed
rm main; go build main.go; sudo ./main
```