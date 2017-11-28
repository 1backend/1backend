# Although the repository creation needs to use the http api thus also a personal access token
# (can't create a repo through command line)
# the usual git commands need an ssh key set up and added to the proper github account.
#
# The ssh key is not in the config.go, it exists on the machine.

REPOSPATH=$1
USERNAME=$2 # author nick
PROJECTNAME=$3
ORGNAME=$4 # the github org used to store the repos eg. 1backend
BOTNAME=$5 # username used to access the git http api
BOTTOKEN=$6 # http api personal access token

echo "Github repo checker script running..."

cd $REPOSPATH

if [ ! -d "./$USERNAME" ]; then
    echo "Trying to clone github repo"
    timeout 5 git clone git@github.com:$ORGNAME/$USERNAME.git
fi

if [ ! -d "./$USERNAME" ]; then
    echo "Cloning did not work. Creating & cloning git repo"
    if curl --user "$BOTNAME:$BOTTOKEN" https://api.github.com/orgs/$ORGNAME/repos -d "{\"name\":\"$USERNAME\",\"description\":\"Generated API clients of user $USERNAME\"}" ; then
        echo "Successfully created repo";
    else
        echo "Create failed. Exiting"; exit 1
    fi
    git clone git@github.com:$ORGNAME/$USERNAME.git
fi

if [ ! -d "./$USERNAME" ]; then
    echo "Git repo still does not exists, exiting"
    exit 1
fi

cd "./$USERNAME"
git pull origin master || true # fresly created repos have no master branch