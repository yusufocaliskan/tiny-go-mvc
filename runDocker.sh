#!/bin/bash

MONGODB1=db

echo "Waiting for MongoDB startup..."
until curl http://${MONGODB1}:27017/serverStatus\?text\=1 2>&1 | grep uptime | head -1; do
  printf '.'
  sleep 1
done

# check if replica set is already initiated
RS_STATUS=$( mongo --quiet --host ${MONGODB1}:27017 -u $MONGO_INITDB_ROOT_USERNAME -p $MONGO_INITDB_ROOT_PASSWORD --eval "rs.status().ok" )
if [[ $RS_STATUS != 1 ]]
then
  echo "[INFO] Replication set config invalid. Reconfiguring now."
  RS_CONFIG_STATUS=$( mongo --quiet --host ${MONGODB1}:27017 -u $MONGO_INITDB_ROOT_USERNAME -p $MONGO_INITDB_ROOT_PASSWORD --eval "rs.status().codeName" )
  if [[ $RS_CONFIG_STATUS == 'InvalidReplicaSetConfig' ]]
  then
    mongo --quiet --host ${MONGODB1}:27017 -u $MONGO_INITDB_ROOT_USERNAME -p $MONGO_INITDB_ROOT_PASSWORD <<EOF
config = rs.config()
config.members[0].host = db # Here is important to set the host name of the db instance
rs.reconfig(config, {force: true})
EOF
  else
    echo "[INFO] MongoDB setup finished. Initiating replicata set."
    mongo --quiet --host ${MONGODB1}:27017 -u $MONGO_INITDB_ROOT_USERNAME -p $MONGO_INITDB_ROOT_PASSWORD --eval "rs.initiate()" > /dev/null
  fi
else
  echo "[INFO] Replication set already initiated."
fi