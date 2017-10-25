#!/usr/bin/env bash
docker-compose stop
docker stop $(docker ps -a -q)
docker rm -f $(docker ps -a -q)
docker rmi $(docker images -q)
docker-compose up -d --build
