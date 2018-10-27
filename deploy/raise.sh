#!/usr/bin/env bash

docker-compose pull && docker swarm init --advertise-addr 68.183.64.110 && docker stack deploy -c docker-compose.yml oa-stack
