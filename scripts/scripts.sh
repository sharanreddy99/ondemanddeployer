#!/bin/bash

function setupRepository() {
    repoName=$1
    repoURL=$2
    branch=$3

    if [ -d "${repoName}" ]; then
        echo "repository: ${repoName} already existsing, pulling latest code."
        cd ${repoName} &&
            git pull origin ${branch}
    else
        git clone ${repoURL}
    fi
}

"$@"
