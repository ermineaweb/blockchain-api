#!/bin/bash

APP_NAME="server-go"

docker build --build-arg SERVICE_NAME="server" -t "$APP_NAME":latest .