REPOSPATH=$1
USERNAME=$2 # author nick
PROJECTNAME=$3
ORGNAME=$4 # the github org used to store the repos eg. 1backend
BOTNAME=$5
BOTTOKEN=$6
BUILDID=$7

echo "Committing generated API..."

cd $REPOSPATH

if [ ! -d "./$USERNAME" ]; then
    echo "Git repo should exists by now. Exiting"
    exit 1
fi

cd "./$USERNAME"

if git diff-index --quiet HEAD --; then
    # no changes
else
    git commit -am "Commit for build $"
    $? || echo "Failed to git commit"; exit 1
    git push https://$BOTNAME:$BOTTOKEN@github.com/$ORGNAME/$USERNAME.git
    $? || echo "Failed to push"; exit 1
fi