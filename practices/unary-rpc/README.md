# Unary gRPC

This project is a very simple implementation of a gRPC server and a client. The client sends a name and the serve response with a greeting to the user.

## Implementation
The server project has the `.proto` file and the server interface as well the _"stub"_ in `/server/api/` folder. 

The `.proto` file has only a `rpc` function defined which gets the name of the user, defined in the `HelloRequest` message, and returns the greeting using the `HelloResponse` message.

```proto
service Hi {
    rpc HelloWorld(HelloRequest) returns (HelloResponse);
}
```

The server interface is implemented in `/server/internal/api`, which only gets the name sended by the client and format a greeting message _"Hello <USERNAME>"_.

## Running the code
To run the code easily there is two Makefile. One in the server service and the other one, in the client service. 

### Running the server
To run and test the service, you must run first the server running the following command in the terminal:

```make
make start
```
```
Building the project
mkdir -p build/bin
go build -v -ldflags='-w -s' -o build/bin $(go list ./... | grep -v test/acceptance)
Starting the project
./build/bin/server
2025/11/15 16:35:14 starting grpc server
2025/11/15 16:35:14 listening grpc server on port :50051
```

With this command, the project will be built and runned and expose the port to send message to it.

### Running the client
After the server is running, you can run the client using the same command, which build and run the client, sending the name defined as a constant.

```makefile
make start
```

```
Building the project
mkdir -p build/bin
go build -v -ldflags='-w -s' -o build/bin $(go list ./... | grep -v test/acceptance)
Starting the project
./build/bin/client 
2025/11/15 16:35:39 creating the grpc client
2025/11/15 16:35:39 creating hiclient
2025/11/15 16:35:39 creating new client service
Hello <USERNAME>
```
