# $1 author
# $2 project name

# returns true if the container is running. idempotent

sudo docker start $1_$2
