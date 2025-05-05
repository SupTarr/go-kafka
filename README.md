# Go Kafka Example

This project demonstrates a simple Kafka producer and consumer implementation in Go using the [IBM/sarama](https://github.com/IBM/sarama) library.

## Prerequisites

* Go (version 1.21 or later recommended)
* Docker
* Docker Compose

## Project Structure

* **[`main.go`](main.go):** The main application entry point. It initializes and runs both the producer and consumer.
* **[`producer/producer.go`](producer/producer.go):** Implements the Kafka [`producer.Producer`](producer/producer.go) struct and logic for sending messages.
* **[`consumer/consumer.go`](consumer/consumer.go):** Implements the Kafka [`consumer.Consumer`](consumer/consumer.go) struct and logic for receiving messages using a consumer group.
* **[`docker-compose.yml`](docker-compose.yml):** Defines the services (Kafka, Zookeeper, Kafdrop) needed to run the example environment.
* **[`go.mod`](go.mod):** Go module definition file.

## Setup

1. **Start the Kafka Cluster:**
Use Docker Compose to start Kafka, Zookeeper, and Kafdrop (a web UI for Kafka).

```sh
docker-compose up -d
```

2. **Verify Kafka:**

You can use Kafdrop to view the Kafka cluster status by navigating to `http://localhost:9000` in your web browser.

## Running the Application

1. **Tidy Go Modules:**

Ensure all dependencies are downloaded.

```sh
go mod tidy
```

2. **Run the Application:**

This command will build and run the [`main.go`](main.go) application. It starts a consumer goroutine and then continuously produces messages to the `topic` topic.

```sh
go run main.go
```

You should see logs from both the producer sending messages and the consumer receiving them in your terminal.

## Stopping the Environment

To stop the Docker containers:

```sh
docker-compose down
