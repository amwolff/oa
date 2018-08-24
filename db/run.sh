#!/usr/bin/env bash

docker run -d --name oa-postgres -p 5432:5432 oa/oadb:latest
