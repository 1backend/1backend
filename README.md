# 1Backend

1Backend is a PaaS designed to make deploying, running and maintaining services cheap and fun.

![1backend service screenshot](https://raw.githubusercontent.com/1backend/1backend/master/1b.png)

#### Why?

We've spent more than 5 years building and using microservices at various companies.
We have seen very many ways of it going wrong.

This project is an excercise in daydreaming about how could it possibly work well.

#### What's the state of the project?

It's extremely early. It's already usable as it is if you want to host it yourself, but don't trust our version yet.

#### What's the tech used?

Angular 2, Go, Material Design, MySQL, Redis.

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