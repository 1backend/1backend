# $1 author
# $2 project name
docker port $1_$2 | awk -F':' '{print $2}'
