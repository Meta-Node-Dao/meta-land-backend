#!/bin/bash

# start redis in docker 
docker run --name comunion-redis -d -p 6379:6379 redis

# start mysql in docker 
docker run --name comunion-mysql -e MYSQL_ROOT_PASSWORD=Comunion2021 -d -p 3306:3306 mysql:5.7 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci 