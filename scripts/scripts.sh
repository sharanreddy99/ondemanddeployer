#!/bin/bash

projectsBaseURL='/home/sharanreddy/Documents'
s3BucketName='sharankonda'

function getPath() {
    newPath="$projectsBaseURL/$2"
    eval "$1=$newPath"
}

function setupProject() {
    echo "entered function"
    repoName=$1
    repoPath=''
    getPath repoPath $repoName

    repoURL=$2
    branch=$3

    cd $projectsBaseURL
    if [ -d "${repoName}" ]; then
        cd ${repoName}
        echo "repository: ${repoName} already existsing, pulling latest code."
        git pull origin ${branch}
    else
        git clone ${repoURL} -b ${branch} ${repoPath}
    fi

    aws s3 cp s3://${s3BucketName}/projects/${repoName}/.env ${repoPath}/aaaatestabcd
}

function buildProject() {
    repoName=$1
    isForceBuild=$2
    repoPath=''
    getPath repoPath $repoName

    cd $repoPath
    if [[ "${isForceBuild}" == "true" ]]; then
        docker compose build --no-cache
    else
        docker compose build
    fi
}

function upProject() {
    repoName=''
    getPath repoName $1

    cd ${repoName} &&
        docker compose up -d
}

function downAllProjects() {
    docker container stop $(docker ps -aq)
}

function clearAllProjects() {
    downAllProjects
    docker volume prune -y
    docker container prune -y
}

"$@"
