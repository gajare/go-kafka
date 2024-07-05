# Go Kafka Project

This project demonstrates a basic implementation of Kafka producer and consumer using Go. The project structure includes separate directories for the producer and consumer, and uses Docker Compose to set up Kafka and Zookeeper.

## Prerequisites

- Docker
- Docker Compose
- Go 1.18 or later

## Project Structure



## Setup

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/go-kafka.git
    cd go-kafka
    ```

2. Start Kafka and Zookeeper using Docker Compose:
    ```sh
    docker-compose up -d
    ```

## Producer

The producer creates a Kafka topic and publishes a message to it.

### Files

- `producer/main.go`: Initializes the Kafka writer and publishes a message.
- `producer/topic.go`: Contains the logic to create a Kafka topic.

### Run the Producer

1. Navigate to the producer directory:
    ```sh
    cd producer
    ```

2. Run the producer:
    ```sh
    go run main.go topic.go
    ```

## Consumer

The consumer reads messages from the Kafka topic.

### Files

- `consumer/main.go`: Initializes the Kafka reader and consumes messages from the topic.

### Run the Consumer

1. Navigate to the consumer directory:
    ```sh
    cd consumer
    ```

2. Run the consumer:
    ```sh
    go run main.go
    ```

## Docker Compose

The `docker-compose.yml` file sets up the necessary Kafka and Zookeeper services.

### Services

- `zookeeper`: Manages Kafka brokers.
- `kafka`: Kafka broker for producing and consuming messages.

### Start Services

To start Kafka and Zookeeper services, run:
```sh
docker-compose up -d