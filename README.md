# Golang to Node gRPC

A test to send JSON data from a Golang client to a Node server using gRPC.

### Generate protos
```
cd go-client/proto
protoc --proto_path "E:\GitHub\GolandProjects\gRPC\go-client\proto" message.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative
```