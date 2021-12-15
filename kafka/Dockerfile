FROM alpine:latest AS builder

ARG kafka_version=3.0.0
ARG scala_version=2.13

ENV FILENAME="kafka_${scala_version}-${kafka_version}.tgz"

RUN apk --no-cache add \
    wget \
    curl \
    jq \
    tar

RUN curl --stderr /dev/null \
    "https://www.apache.org/dyn/closer.cgi?path=/kafka/${kafka_version}/${FILENAME}&as_json=1" | \
    jq -r '"\(.preferred)\(.path_info)"' | \
    wget -i - -O /tmp/${FILENAME}

RUN mkdir -p /opt && tar xzf /tmp/${FILENAME} -C /opt/ --one-top-level=kafka --strip-components 1

# ---

FROM openjdk:8-jre-alpine

RUN apk --no-cache add \
    bash

COPY --from=builder /opt/kafka /opt/kafka

WORKDIR /opt/kafka
ENV KAFKA_HOME /opt/kafka

# CMD ["./bin/kafka-server-start.sh", "config/server.properties"]
COPY start.sh start.sh
CMD ["./start.sh"]
