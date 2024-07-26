# Overview
An example microservices project that uses gRPC for inter-service communication.

Service API definitons use Protocol Buffers. They are managed along with the generated source code in a separate repository: https://github.com/hollowdll/grpc-microservices-proto

Microservices in this project are written in the Go programming language. They use hexagonal architecture as their inner architectural style.

This project was developed as part of a thesis that discusses the use cases and benefits of using gRPC for synchronous microservice inter-service communication.

# About the project

The project consists of the following 3 microservices:
- Order service
- Inventory service
- Payment service

The goal is to present a simple e-commerce order operation, where an order can have multiple order items with different quantities.

Currently the services are very minimal, because they were designed to mainly showcase gRPC usage. For this reason they lack a lot of features.

# How to run the microservices

Make sure that you have Go installed. Instructions [here](https://go.dev/doc/install).

You need to use a terminal and be able to type commands.

Clone this repository manually or with Git:
```sh
git clone https://github.com/hollowdll/go-grpc-microservices.git
```

Go to the repository depending on where you cloned it:
```sh
cd go-grpc-microservices
```

It is recommended to have 3 different terminal windows/views open so you can see the output logs of each microservice while running all of them simultaneously in different terminal windows.

You can start a service by going to its directory and then running it.

For example, to start the payment service do the following:
```sh
cd services/payment
go run cmd/main.go
```

After this you should see some output logs telling the service is starting if nothing went wrong.

# How to configure the microservices

By default, the services use default configurations. However, you can change these defaults with configuration files or environment variables, as the [12-factor app](https://12factor.net/) methodology suggests.

## About the configuration files

- The configuration file needs to be in the same directory where you are running the service
- The configuration files use YAML format
- The configuration files are not created automatically
- Environment variables override configurations in the configuration files

## Configure payment service

Configuration file: payment-config.yaml

Configuration file template with default values:
```yaml
GRPC_PORT: 9000
APPLICATION_MODE: development
```

The following table describes the configurations for this service:

Name in config file | Environment variable     | Default value | Description
------------------- | ------------------------ | ------------- | -----------
GRPC_PORT           | PAYMENT_GRPC_PORT        | 9000          | gRPC server port number.
APPLICATION_MODE    | PAYMENT_APPLICATION_MODE | development   | Mode the service runs in. e.g. development, staging, production.

## Configure inventory service

Configuration file: inventory-config.yaml

Configuration file template with default values:
```yaml
GRPC_PORT: 9001
APPLICATION_MODE: development
```

The following table describes the configurations for this service:

Name in config file | Environment variable     | Default value | Description
------------------- | ------------------------ | ------------- | -----------
GRPC_PORT           | INVENTORY_GRPC_PORT        | 9001          | gRPC server port number.
APPLICATION_MODE    | INVENTORY_APPLICATION_MODE | development   | Mode the service runs in. e.g. development, staging, production.

## Configure order service

Configuration file: order-config.yaml

Configuration file template with default values:
```yaml
GRPC_PORT: 9002
APPLICATION_MODE: development
INVENTORY_SERVICE_HOST: localhost
INVENTORY_SERVICE_PORT: 9001
PAYMENT_SERVICE_HOST: localhost
PAYMENT_SERVICE_PORT: 9000
```

The following table describes the configurations for this service:

Name in config file    | Environment variable         | Default value | Description
---------------------- | ---------------------------- | ------------- | -----------
GRPC_PORT              | ORDER_GRPC_PORT              | 9002          | gRPC server port number.
APPLICATION_MODE       | ORDER_APPLICATION_MODE       | development   | Mode the service runs in. e.g. development, staging, production.
INVENTORY_SERVICE_HOST | ORDER_INVENTORY_SERVICE_HOST | localhost     | Host or IP address of the inventory service.
INVENTORY_SERVICE_PORT | ORDER_INVENTORY_SERVICE_PORT | 9001          | Port number of the inventory service's gRPC server.
PAYMENT_SERVICE_HOST   | ORDER_PAYMENT_SERVICE_HOST   | localhost     | Host or IP address of the payment service.
PAYMENT_SERVICE_PORT   | ORDER_PAYMENT_SERVICE_PORT   | 9000          | Port number of the payment service's gRPC server.


