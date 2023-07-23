# Efficient Workflows in Temporal: Signals & Selectors Tutorial

This tutorial focuses on two fundamental concepts in Temporal: Signals and Selectors. Signals allow seamless communication with running workflows, enabling notifications, state updates, and triggering actions based on external events. On the other hand, Selectors provide a flexible and efficient approach to orchestrating and coordinating multiple workflows, enhancing system robustness and scalability.

You can find the complete tutorial here, and it is part of the comprehensive series on efficient workflow development in Go using Temporal.io: [(Series) Efficient Workflow in Go with Temporal.io.
](https://medium.com/@younisjad/series-efficient-workflow-in-go-with-temporal-io-ab3248b1c5f9)

## Dependencies

Before running this example, ensure you have the following dependencies:

- Go 1.18.x or later (install via brew install go on macOS)
- Temporal Server (run via docker run --rm -p 7233:7233 temporalio/temporal:latest)
- Temporal CLI (install via brew install temporalio/tap/tctl on macOS)

## Setup

To get started, clone this repository or download the source code.

Next, install the dependencies using go mod:

`go mod tidy`

Start the Temporal Server either through Docker or manually on your machine.

## Docker Setup

For local testing and development, you can use Docker to run Temporal. The Temporal team provides an official Docker image that allows running the entire Temporal stack locally.Ensure Docker is installed on your machine by downloading it from the official website.Once Docker is installed, utilize the official docker-compose setup provided by Temporal to run the complete Temporal stack locally. Run the following command to acquire the docker-compose setup files:

`curl -o docker-compose.yml https://raw.githubusercontent.com/temporalio/docker-compose/master/docker-compose-cas.yml`

This command downloads the docker-compose setup for Temporal Community Edition, which incorporates a Cassandra data store.Now, start the services by running the following command:

`docker-compose up`

## Running the Example

Finally, start the Go HTTP server:

`go run main.go`

Now, you can run the example code using the following command:

`curl -X GET http://localhost:5000/weather?city=Cairo`

Explore the tutorial and experiment with Temporal’s Signals and Selectors to build efficient and scalable applications!

# Happy coding!