#!/bin/bash

APP_NAME="geth"

docker build -t "$APP_NAME":latest .
docker run -it \
--net host \
"$APP_NAME":latest \
attach http://localhost:8545 