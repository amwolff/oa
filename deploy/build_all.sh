#!/usr/bin/env bash

set -e
set -u
set -x

REPO=${GOPATH}/src/github.com/amwolff/oa
DEPLOYMENT=${GOPATH}/src/github.com/amwolff/oa/deploy

VERSION=$1

function buildBinary {
    COMMIT=$(git rev-list -1 HEAD)
    TIME=$(date +"%H:%M-%d-%m-%Y")
    DEV='false'
    
    cd ${REPO}/cmd/$1
    env GOARCH=amd64 GOOS=linux go build -a \
    -ldflags="-X main.BuildTimeCommitMD5=$COMMIT -X main.BuildTimeTime=$TIME -X main.BuildTimeIsDev=$DEV" \
    -o ${DEPLOYMENT}/cache/$1/$1
}

function buildContainer {
    buildBinary $1
    
    NOCACHE=${DEPLOYMENT}/services/$1
    CACHE=${DEPLOYMENT}/cache/$1
    
    cp ${NOCACHE}/config.yml ${CACHE}/
    cp ${NOCACHE}/container/Dockerfile ${CACHE}/
    cp ${NOCACHE}/container/.dockerignore ${CACHE}/
    
    cd ${CACHE}
    
    TAG=amwolff/oa:$1_${VERSION}
    docker build --tag ${TAG} .
    docker push ${TAG}
}

cp ${REPO}/db/*.sql ${DEPLOYMENT}/cache/db/
cp ${REPO}/db/Dockerfile ${DEPLOYMENT}/cache/db/
cd ${DEPLOYMENT}/cache/db
docker build --tag amwolff/oa:oadb_${VERSION} .
docker push amwolff/oa:oadb_${VERSION}

buildContainer dataharvester

buildContainer api

cp ${REPO}/pkg/frontend/* ${DEPLOYMENT}/cache/dirserver/
buildContainer dirserver
