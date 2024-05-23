#!/bin/bash
sleep 5
until mongosh --host mongo_gptv:27017 --file "$(dirname "$0")/init.js"
do
  sleep 1
done
