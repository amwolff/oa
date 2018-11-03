#!/usr/bin/env bash

set -u
set -x

CACHEDIR=${GOPATH}/src/github.com/amwolff/oa/deploy/cache

rm -i ${CACHEDIR}/api/*
rm -i ${CACHEDIR}/dataharvester/*
rm -i ${CACHEDIR}/db/*
rm -i ${CACHEDIR}/dirserver/*