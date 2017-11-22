# $1 build path
# $2 author
# $3 app name
# $4 infrastructure password
# $5 recipe path (eg: "go", "nodejs-whatever")
# $6 absolute path to project
cp $6/tech-pack/$5/Dockerfile $1/Dockerfile
INTERNALIP=$(ip route get 8.8.8.8 | head -1 | cut -d' ' -f8)
NAME="\`$2_$3\`"
cd $1
sudo docker build -t $2_$3 . || exit 1
sudo docker stop $2_$3
sudo docker rm $2_$3
echo "CREATE DATABASE IF NOT EXISTS $NAME; CREATE USER IF NOT EXISTS $NAME@'%'; GRANT ALL PRIVILEGES ON $NAME.* TO $NAME@'%' identified by '$4'; GRANT ALL PRIVILEGES ON $NAME.* TO $NAME@'localhost' identified by '$4'; FLUSH PRIVILEGES;" | mysql -h localhost -P 3306 --protocol=tcp -u root -proot || exit 1
echo "SET PASSWORD FOR $NAME@'%' = PASSWORD('$4');" | mysql -h localhost -P 3306 --protocol=tcp -u root -proot || exit 1
# echo "CREATE DATABASE IF NOT EXISTS $2_$3; CREATE USER IF NOT EXISTS $2_$3@'%' IDENTIFIED BY 't8ecNpCf5u0d'; GRANT ALL PRIVILEGES ON $2_$3.* TO $2_$3@'localhost';" | mysql -h localhost -P 3306 --protocol=tcp -u root -proot
sudo docker run --name $2_$3 -e MYSQLIP=$INTERNALIP -e INFRAPASS=$4 -p=8883 -d $2_$3
