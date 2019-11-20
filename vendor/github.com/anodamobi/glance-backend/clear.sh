#!/usr/bin/env bash
# Delete all containers
docker rm -f $(docker ps -a -q) --force

# Delete all images
docker rmi -f $(docker images -q) --force

# https://github.com/chadoe/docker-cleanup-volumes

docker run --rm -v /var/run/docker.sock:/var/run/docker.sock:ro -v /var/lib/docker:/var/lib/docker martin/docker-cleanup-volumes

docker rmi martin/docker-cleanup-volumes --force

docker network prune
