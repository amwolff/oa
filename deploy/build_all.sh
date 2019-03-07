#!/usr/bin/env bash

set -e
set -u
set -x

DEV=false
VER='<unset>'

while getopts ":dv:" opt; do
    case ${opt} in
        'd')
            DEV=true
        ;;
        'v')
            VER=$OPTARG
        ;;
        \? )
            echo "Invalid option: $OPTARG" 1>&2
        ;;
        ':')
            echo "Invalid option: $OPTARG requires an argument" 1>&2
        ;;
    esac
done
# shift $((OPTIND -1))

REPOSITORY=${GOPATH}/src/github.com/amwolff/oa
DEPLOYMENT=${GOPATH}/src/github.com/amwolff/oa/deploy

function buildBinary {
    COMMIT=$(git rev-list -1 HEAD)
    TIME=$(date +"%H:%M-%d-%m-%Y")
    
    cd ${REPOSITORY}/cmd/$1
    env GOARCH=amd64 GOOS=linux go build -v -a \
    -ldflags="-X main.BuildTimeCommitMD5=$COMMIT -X main.BuildTimeTime=$TIME -X main.BuildTimeIsDev=$DEV" \
    -o ${DEPLOYMENT}/cache/$1/$1
}

function buildContainer {
    buildBinary $1
    
    NOCACHE=${DEPLOYMENT}/services/$1
    CACHE=${DEPLOYMENT}/cache/$1
    
    if [[ "$DEV" = true ]]; then
        cp ${NOCACHE}/example_config.yml ${CACHE}/config.yml
    else
        cp ${NOCACHE}/config.yml ${CACHE}/
    fi
    
    cp ${NOCACHE}/container/Dockerfile ${CACHE}/
    cp ${NOCACHE}/container/.dockerignore ${CACHE}/
    
    cd ${CACHE}
    
    TAG=amwolff/oa:$1_${VER}
    docker build --tag ${TAG} .
    if [[ "$DEV" = false ]]; then
        docker push ${TAG}
    fi
}

cp ${REPOSITORY}/db/Dockerfile ${DEPLOYMENT}/cache/db/
cp ${REPOSITORY}/db/*.sql ${DEPLOYMENT}/cache/db/
cd ${DEPLOYMENT}/cache/db
docker build --tag amwolff/oa:oadb_${VER} .
if [[ "$DEV" = false ]]; then
    docker push amwolff/oa:oadb_${VER}
fi

buildContainer dataharvester

buildContainer api

cp ${REPOSITORY}/pkg/frontend/icos/* ${DEPLOYMENT}/cache/dirserver/
cp ${REPOSITORY}/pkg/frontend/dist/* ${DEPLOYMENT}/cache/dirserver/
if [[ "$DEV" = true ]]; then
    mv ${DEPLOYMENT}/cache/dirserver/app_dev.js ${DEPLOYMENT}/cache/dirserver/app.js
else
    rm ${DEPLOYMENT}/cache/dirserver/app_dev.js
fi
buildContainer dirserver

cp ${REPOSITORY}/deploy/services/pinger/container/* ${DEPLOYMENT}/cache/pinger/
cp ${REPOSITORY}/cmd/pinger/pinger.sh ${DEPLOYMENT}/cache/pinger/
cd ${DEPLOYMENT}/cache/pinger
docker build --tag amwolff/oa:pinger_${VER} .
if [[ "$DEV" = false ]]; then
    docker push amwolff/oa:pinger_${VER}
fi
