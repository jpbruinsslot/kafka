#!/usr/bin/env bash
ZOOKEEPER_URL=zookeeper:2181
sed -ri "s/(zookeeper.connect)=(.*)/\1=$ZOOKEEPER_URL/g" $KAFKA_HOME/config/server.properties

sed -ri "s/(broker.id)=(.*)/\1=-1/g" $KAFKA_HOME/config/server.properties

ADVERTISED_LISTENERS="PLAINTEXT:\/\/$(cat /etc/hosts | grep $HOSTNAME | awk '{print $1}'):9092"
sed -ri "s/#(advertised.listeners)=(.*)/\1=$ADVERTISED_LISTENERS/g" $KAFKA_HOME/config/server.properties

echo $ADVERTISED_LISTENERS

$KAFKA_HOME/bin/kafka-server-start.sh $KAFKA_HOME/config/server.properties
