#!/bin/bash

projectsBaseURL='/home/ubuntu'
s3BucketName='sharankonda'

function getPath() {
    echo "Entered getPath function"
    newPath="$projectsBaseURL/$2"
    eval "$1=$newPath"
}

function setupProject() {
    echo "Entered setupProject function"
    repoName=$1
    repoPath=''
    getPath repoPath $repoName

    repoURL=$2
    branch=$3
    forceBuild=$4

    cd $projectsBaseURL
    if [ -d "${repoName}" ]; then
        cd ${repoName}
        echo "repository: ${repoName} already existsing, pulling latest code."
        git pull origin ${branch}
    else
        git clone ${repoURL} -b ${branch} ${repoPath}
    fi

    aws s3 cp s3://${s3BucketName}/projects/${repoName}/.env ${repoPath}/aaaatestabcd
    buildProject $repoName $forceBuild
    upProject $repoName

}

function buildProject() {
    echo "Entered buildProject function"
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
    downAllProjects
    echo "Entered upProject function"
    repoName=''
    getPath repoName $1

    cd ${repoName} &&
        docker compose up -d
}

function downAllProjects() {
    echo "Entered downAllProject function"
    docker container stop $(docker ps -q)
}

function clearAllProjects() {
    downAllProjects
    docker volume prune -y
    docker container prune -y
}

"$@"
