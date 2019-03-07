#!/usr/bin/env bash

# mkdir -p /opt/traefik
# touch /opt/traefik/acme.json && chmod 600 /opt/traefik/acme.json
# touch /opt/traefik/traefik.toml

sysctl -w \
net.ipv4.tcp_keepalive_time=600 \
net.ipv4.tcp_keepalive_intvl=60 \
net.ipv4.tcp_keepalive_probes=3

docker-compose pull && docker swarm init --advertise-addr 159.89.5.189 && docker stack deploy -c docker-compose.yml oa-stack
