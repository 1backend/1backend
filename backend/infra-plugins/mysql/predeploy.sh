# $1 author
# $2 app name
# $3 infrastructure password
# $4 path
# $5 mysql ip

NAME="\`$1_$2\`" # quoted db and database name
echo "CREATE DATABASE IF NOT EXISTS $NAME; CREATE USER IF NOT EXISTS $NAME@'%'; GRANT ALL PRIVILEGES ON $NAME.* TO $NAME@'%' identified by '$3'; GRANT ALL PRIVILEGES ON $NAME.* TO $NAME@'localhost' identified by '$3'; FLUSH PRIVILEGES;" | mysql -h $5 -P 3306 --protocol=tcp -u root -proot || exit 1
echo "SET PASSWORD FOR $NAME@'%' = PASSWORD('$3');" | mysql -h $5 -P 3306 --protocol=tcp -u root -proot || exit 1

