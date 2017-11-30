# this script only works for the angular client yet - no ts or node supported

REPOSPATH=$1
USERNAME=$2 # author nick
PROJECTNAME=$3
NPMTOKEN=$4

echo "Publishing to NPM..."

cd $REPOSPATH

if [ ! -d "./$USERNAME/ng/$PROJECTNAME" ]; then
    echo "Ng project folder should exist by now, exiting"
    exit 1
fi

cd "./$USERNAME/ng/$PROJECTNAME"

docker run --rm -v "$PWD:/usr/src/app" -e NPM_TOKEN=$NPMTOKEN crufter/npm-publish