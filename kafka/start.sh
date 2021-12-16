#!/usr/bin/env bash
ADVERTISED_LISTENERS="PLAINTEXT:\/\/$(cat /etc/hosts | grep $HOSTNAME | awk '{print $1}'):9092"
sed -ri "s/#(advertised.listeners)=(.*)/\1=$ADVERTISED_LISTENERS/g" ./config/server.properties

./bin/kafka-server-start.sh ./config/server.properties
