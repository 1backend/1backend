# $1 sql query
# $2 author
# $3 project name
echo "USE $2_$3; $1" | mysql -h localhost -P 3306 --protocol=tcp -u root -proot
