#!/usr/bin/env bash

docker stack rm oa-stack && docker swarm leave --force
