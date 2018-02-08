BUILD_PATH=$1
AUTHOR=$2
PROJECT_NAME=$3
RECIPE_FOLDER=$4  # (eg: "go", "nodejs-whatever")
PROJECT_PATH=$5 # absolute
# envar is assumed to be in build folder under name "envars"

cp $PROJECT_PATH/tech-plugins/$RECIPE_FOLDER/Dockerfile $BUILD_PATH/Dockerfile
INTERNALIP=$(ip route get 8.8.8.8 | head -1 | cut -d' ' -f8)
cd $BUILD_PATH

# image and container name is the same
CONTAINER=$AUTHOR"_"$PROJECT_NAME

docker build -t $CONTAINER . || exit 1
docker stop $CONTAINER
docker rm $CONTAINER

docker run --name $CONTAINER --env-file=$BUILD_PATH/env -p=8883 -d $CONTAINER
