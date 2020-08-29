#!/usr/bin/env bash
echo Sending SIGHUP to:
docker ps --format '{{.Names}}' | grep 'telegraf' | while read line; do (docker kill -s HUP $line); done