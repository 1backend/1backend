# this script only works for the angular client yet - no ts or node supported

REPOSPATH=$1
USERNAME=$2 # author nick
PROJECTNAME=$3

echo "Committing generated API..."

cd $REPOSPATH

if [ ! -d "./$USERNAME/ng/$PROJECTNAME" ]; then
    echo "Ng project folder should exist by now, exiting"
    exit 1
fi

cd "./$USERNAME/ng/$PROJECTNAME"

npm version patch

npm publish --access public