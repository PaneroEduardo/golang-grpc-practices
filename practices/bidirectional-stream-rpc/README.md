# Stream bidirectional gRPC

This project is a very simple implementation of a gRPC bidirectional streaming. The client sends a stream which contains a message sended for X times and the server returns the message with an echo.

## Implementation
The server project has the `.proto` file and the server interface as well the _"stub"_ in `/server/api/` folder. 

The `.proto` file has only a `rpc` function defined which gets the message and echo the message into the stream, which is passed as stream, defined in the `MessageRequest` message, and returns the echo from the message in the `MessageResponse` struct.

```proto
service Echo {
    rpc EchoMessage(stream MessageRequest) returns (stream MessageResponse);
}
```

The server interface is implemented in `/server/internal/api`, which only implements the stream reading and the echo for the message.

## Running the code
To run the code easily there is two Makefile. One in the server service and the other one, in the client service. 

### Running the server
To run and test the service, you must run first the server running the following command in the terminal:

```sh
make start
```

Once the client was executed, the server log the request received.
```
Building the project
mkdir -p build/bin
go build -v -ldflags='-w -s' -o build/bin $(go list ./... | grep -v test/acceptance)
github.com/PaneroEduardo/golang-grpc-practices/practices/bidirectional-stream-rpc/server/cmd/server
Starting the project
./build/bin/server
2025/11/17 17:59:32 starting grpc server
2025/11/17 17:59:32 listening grpc server on port :50051
2025/11/17 18:00:04 request message received Hello World
2025/11/17 18:00:04 request message received echo Hello World
2025/11/17 18:00:04 request message received echo echo Hello World
2025/11/17 18:00:04 request message received echo echo echo Hello World
2025/11/17 18:00:04 request message received echo echo echo echo Hello World
```

With this command, the project will be built and runned and expose the port to send message to it.

### Running the client
After the server is running, you can run the client using the same command, which build and run the client, sending the name defined as a constant.

```sh
make start
```

Once the code is built and runned, the terminal returns the result.
```
Building the project
mkdir -p build/bin
go build -v -ldflags='-w -s' -o build/bin $(go list ./... | grep -v test/acceptance)
github.com/PaneroEduardo/golang-grpc-practices/practices/bidirectional-stream-rpc/client/cmd/client
Starting the project
./build/bin/client 
2025/11/17 18:00:04 creating the grpc client
2025/11/17 18:00:04 creating echo client
2025/11/17 18:00:04 creating new client service
2025/11/17 18:00:04 sending the message to the stream to get the echoes
2025/11/17 18:00:04 message received: echo Hello World
2025/11/17 18:00:04 message received: echo echo Hello World
2025/11/17 18:00:04 message received: echo echo echo Hello World
2025/11/17 18:00:04 message received: echo echo echo echo Hello World
2025/11/17 18:00:04 message received: echo echo echo echo echo Hello World
```
