REPOSPATH=$1
USERNAME=$2 # author nick
PROJECTNAME=$3
ORGNAME=$4 # the github org used to store the repos eg. 1backend
BOTNAME=$5
BOTTOKEN=$6

echo "Github repo checker script running..."

cd $REPOSPATH

if [ ! -d "./$USERNAME" ]; then
    echo "Trying to clone github repo"
    git clone https://github.com/$ORGNAME/$USERNAME
fi

if [ ! -d "./$USERNAME" ]; then
    echo "Cloning did not work. Creating & cloning git repo"
    curl --user '$USERNAME:$ACCESSTOKEN' https://api.github.com/orgs/$ORGNAME/repos -d "{\"name\":\"$USERNAME\"}"
    $? || echo "Create failed. Exiting"; exit 1
    git clone https://github.com/$ORGNAME/$USERNAME
    $? || echo "Cloning failed. Exiting"; exit 1
fi

if [ ! -d "./$USERNAME" ]; then
    echo "Git repo still does not exists, exiting"
    exit 1
fi