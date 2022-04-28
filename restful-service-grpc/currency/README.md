# Currency Service

A simple example to understand gRPC in golang!

## Directory Structure

```
.
├── currency.proto
├── main.go
├── Makefile
├── protos
│   ├── currency_grpc.pb.go
│   └── currency.pb.go
├── README.md
└── server
    └── currency.go
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
    Currency
    grpc.reflection.v1alpha.ServerReflection
    ```

### List Methods
* Command:
    ```grpcurl --plaintext localhost:8080 list Currency```
* Output:
    ```Currency.GetRate```

### Method detail for GetRate
* Command:
    ```grpcurl --plaintext localhost:8080 describe Currency.GetRate```
* Output:
    ```
    Currency.GetRate is a method:
    rpc GetRate ( .RateRequest ) returns ( .RateResponse );
    ```

### RateRequest detail
* Command:
    ```grpcurl --plaintext localhost:8080 describe .RateRequest```
* Output:
    ```
    RateRequest is a message:
    message RateRequest {
        string Base = 1;
        string Destination = 2;
    }
    ```

### RateResponse detail
* Command:
    ```grpcurl --plaintext localhost:8080 describe .RateResponse```
* Output:
    ```
    RateResponse is a message:
    message RateResponse {
        float rate = 1;
    }
    ```

### Execute a request
* Command:
    ```grpcurl --plaintext -d '{"Base":"USD","Destination":"INR"}' localhost:8080 Currency.GetRate```
* Output:
    ```
    {
        "rate": 0.5
    }
    ```
