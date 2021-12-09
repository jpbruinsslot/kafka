# Kafka

## Usage

```
$ docker-compose up --build --scale kafka=3 kafka
```

## CLI

You can use the CLI commands by using the `kafka.sh` script
and pass in your commands. E.g.

```
$ ./kafka.sh kafka-topics \
    --bootstrap-server localhost:9092
    --partitions 3 \
    --replication-factor 1
    --topic first_topic \
    --create
```
