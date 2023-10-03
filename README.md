## gRPC Time Service

## Overview

This repository contains a gRPC service written in Go that provides the current time and server uptime. The service uses Protocol Buffers for efficient serialization of data.
The client fetches the current time and the server uptime from the server.<br>
We also use the `context` package to implement a timeout if the server takes too long to serve the request.

## Features

- Get the current time in HH:MM:SS format
- Get the server uptime in duration format

## Prerequisites

- Go 1.16 or higher
- gRPC
- Protocol Buffers

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/laurentknauss/grpc-time-service.git
    ```

2. Navigate to the project directory:

    ```bash
    cd grpc-time-service
    ```

3. Install the required packages:

    ```bash
    go mod tidy
    ```

## Usage

1. Start the server:

    ```bash
    go run server/main.go
    ```

2. Start the client:

    ```bash
    go run client/main.go
    ```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
