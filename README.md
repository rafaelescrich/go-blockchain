# go-blockchain

This project demonstrates a basic blockchain implementation in Go. It includes features such as SHA-256 hashing for blocks, a simple proof-of-work (PoW) mining system, TCP connections for node communication, and a REST API for interacting with the blockchain. The REST API allows adding transactions and mining new blocks.

## Features

- SHA-256 hashing for blocks, similar to Bitcoin.
- Simple proof-of-work algorithm for mining blocks.
- REST API for managing transactions and mining.
- Merkle tree calculation for transactions within a block.

## Getting Started

### Prerequisites

- Go (1.20 or later recommended)
- Postman (for testing the REST API, optional)

### Installation

1. Clone the repository to your local machine.
2. Navigate to the project directory.
3. Use `go get -u github.com/gin-gonic/gin` to install the Gin framework, which is required for the REST API.

### Running the Project

1. Build the project with `go build -o go-blockchain`.
2. Run the server specifying the port (e.g., `./go-blockchain -port=8080`).

### REST API Usage

You can interact with the blockchain using the following REST endpoints:

- `POST /transactions/new`: Add a new transaction to the blockchain.
  - Example request body:
    ```json
    {
      "Sender": "sender_address",
      "Recipient": "recipient_address",
      "Amount": 10
    }
    ```
- `GET /mine`: Mine a new block with the current transactions.

Use a tool like Postman or curl to make requests to these endpoints.

## Testing

A Postman collection is included for easy testing of the REST API. Import the collection into Postman and adjust the port as necessary.

## Customization

You can customize the blockchain parameters, such as difficulty level and port number, by modifying the source code.

## Contributing

Contributions to this project are welcome. Please fork the repository and submit a pull request with your changes.

## License

This project is open source and available under the [MIT License](LICENSE.md).
