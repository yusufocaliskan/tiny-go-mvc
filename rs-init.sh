#!/bin/bash

# Start MongoDB
mongod --replSet my-replica-set --bind_ip_all --port 27017 &

# Wait for MongoDB to start
sleep 10

# Initialize the replica set
mongo --eval "rs.initiate()"

# Keep the container running
tail -f /dev/null
