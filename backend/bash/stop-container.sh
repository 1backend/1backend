# $1 author
# $2 project name

# returns success if a container is stopped. idempotent
docker stop -t=0 $1_$2
