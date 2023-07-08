#!/bin/bash

APP_NAME="blockchain"

docker build -t "$APP_NAME":latest .
docker run -it \
--net host \
"$APP_NAME":latest \
--host 0.0.0.0 \
--miner.blockTime=10 \
--logging.verbose=true