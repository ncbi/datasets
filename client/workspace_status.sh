#!/usr/bin/env bash

function normalizeTag()
{
    local TAG=$1
    if [[ -z $TAG ]]
    then
        printf ""
        return
    fi
    TAG=${TAG//+/-}
    TAG=${TAG//\//-}
    printf "$TAG"
}

function getVersion()
{
    if [ "$1" == "--stable" ]; then
        params="--no-abbr"
    else
        params=""
    fi
    version=$(git describe --tags $params 2>/dev/null)
    echo ${version} | awk '''BEGIN {FS="-"}
        $1=="" {print "[git-pep440] The \"git describe\" output is empty. Are you sure you have tags?" > "/dev/stderr"; print "0.0"; exit 0}
        $2=="" {print $1; exit 0}
        {printf("%s.post%s+%s", $1, $2, $3)}
    '''

}

GIT_BRANCH="$(git rev-parse --abbrev-ref HEAD)"

version=$(getVersion --stable)
version_without_prefix=$(echo $version | sed 's/^v//')

cat <<EOF
SOFTWARE_VERSION $(getVersion)
STABLE_SOFTWARE_VERSION $version
STABLE_SOFTWARE_VERSION_WITHOUT_PREFIX $version_without_prefix
STABLE_GIT_COMMIT $(git rev-parse HEAD)
STABLE_GIT_COMMIT_SHORT $(git rev-parse --short HEAD)
GIT_BRANCH $GIT_BRANCH
GIT_BRANCH_TAG $(normalizeTag $GIT_BRANCH)
EOF
