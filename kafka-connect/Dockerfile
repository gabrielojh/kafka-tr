FROM confluentinc/cp-kafka-connect:latest

RUN confluent-hub install --no-prompt confluentinc/kafka-connect-s3-source:2.5.12 \
    && confluent-hub install --no-prompt confluentinc/kafka-connect-s3:10.5.7 \
    && confluent-hub install --no-prompt confluentinc/kafka-connect-gcs-source:2.5.13 \
    && confluent-hub install --no-prompt mongodb/kafka-connect-mongodb:1.11.2