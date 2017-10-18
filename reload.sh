#!/bin/bash
docker-compose stop go-stat
docker rmi go-stat
docker-composer up -d --build