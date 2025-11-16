# Client server gRPC

This project is a very simple implementation of a gRPC streaming client. The client sends a stream which contains a set of numbers to be sumed and then, the server returns the result of the sum.

## Implementation
The server project has the `.proto` file and the server interface as well the _"stub"_ in `/server/api/` folder. 

The `.proto` file has only a `rpc` function defined which gets the values, which is passed as stream, defined in the `SumItemRequest` message, and returns the sum of these values using the `SumResponse` message.

```proto
service Sum {
    rpc SumItems(stream SumItemRequest) returns (SumResponse);
}
```

The server interface is implemented in `/server/internal/api`, which only implements the stream reading and the sum of the values.

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
github.com/PaneroEduardo/golang-grpc-practices/practices/client-stream-rpc/server/cmd/server
Starting the project
./build/bin/server
2025/11/16 11:54:33 starting grpc server
2025/11/16 11:54:33 listening grpc server on port :50051
2025/11/16 11:59:49 request value received 1
2025/11/16 11:59:49 request value received 2
2025/11/16 11:59:49 request value received 3
2025/11/16 11:59:49 request value received 4
2025/11/16 11:59:49 request value received 5
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
github.com/PaneroEduardo/golang-grpc-practices/practices/client-stream-rpc/client/cmd/client
Starting the project
./build/bin/client 
2025/11/16 11:59:49 creating the grpc client
2025/11/16 11:59:49 creating sumclient
2025/11/16 11:59:49 creating new client service
2025/11/16 11:59:49 sending the numbers as streaming
2025/11/16 11:59:49 the result of the sum is 45
```
