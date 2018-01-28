# $1 build path
# $2 author
# $3 app name
# $4 infrastructure password
# $5 recipe path (eg: "go", "nodejs-whatever")
# $6 absolute path to project
# $7 caller id (to support namespacing)

NAME="\`$2_$3\`" # quoted db and database name
echo "CREATE DATABASE IF NOT EXISTS $NAME; CREATE USER IF NOT EXISTS $NAME@'%'; GRANT ALL PRIVILEGES ON $NAME.* TO $NAME@'%' identified by '$4'; GRANT ALL PRIVILEGES ON $NAME.* TO $NAME@'localhost' identified by '$4'; FLUSH PRIVILEGES;" | mysql -h localhost -P 3306 --protocol=tcp -u root -proot || exit 1
echo "SET PASSWORD FOR $NAME@'%' = PASSWORD('$4');" | mysql -h localhost -P 3306 --protocol=tcp -u root -proot || exit 1

