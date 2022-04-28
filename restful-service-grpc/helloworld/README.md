# Hello World

Hello world in gRPC!

## Directory Structure
```
.
├── helloworld.proto
├── main.go
├── Makefile
├── protos
│   ├── helloworld_grpc.pb.go
│   └── helloworld.pb.go
├── README.md
└── server
    └── server.go
```

## Testing
To test the system install `grpccurl` which is a command line tool which can interact with gRPC API's

Link to the repo is [here!](https://github.com/fullstorydev/grpcurl)

```go install github.com/fullstorydev/grpcurl/cmd/grpcurl```

### Run the server
* Command:
    ```go run main.go```

### List Services
* Command: 
    ```grpcurl --plaintext localhost:8080 list```
* Output:
    ```
    Greeter
    grpc.reflection.v1alpha.ServerReflection
    ```

### List Methods
* Command:
    ```grpcurl --plaintext localhost:8080 list Currency```
* Output:
    ```Greeter.SayHello```

### Method detail for GetRate
* Command:
    ```grpcurl --plaintext localhost:8080 describe Currency.GetRate```
* Output:
    ```
    Greeter.SayHello is a method:
    rpc SayHello ( .HelloRequest ) returns ( .HelloResponse );
    ```

### RateRequest detail
* Command:
    ```grpcurl --plaintext localhost:8080 describe .HelloRequest```
* Output:
    ```
    HelloRequest is a message:
    message HelloRequest {
        string name = 1;
    }
    ```

### RateResponse detail
* Command:
    ```grpcurl --plaintext localhost:8080 describe .HelloResponse```
* Output:
    ```
    HelloResponse is a message:
    message HelloResponse {
        string message = 1;
    }
    ```

### Execute a request
* Command:
    ```grpcurl --plaintext -d '{"name":"Ankit"}' localhost:8080 Greeter.SayHello```
* Output:
    ```
    {
        "message": "Hello Ankit"
    }
    ```

