const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const PROTO_PATH = __dirname + '/../go-client/proto/queue.proto';
const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
    keepCase: true, longs: String, enums: String, defaults: true, oneofs: true
});
const proto = grpc.loadPackageDefinition(packageDefinition).gateway;

function Push(call, callback) {
    console.log({key: call.request.key, value: JSON.parse(call.request.value.toString())})
    return callback()
}

function main() {
    const server = new grpc.Server();
    server.addService(proto.Queue.service, {Push: Push});
    server.bindAsync('localhost:4362', grpc.ServerCredentials.createInsecure(), (_,port) => {
        server.start();
        console.log(`Server started at :${port}`)
    });
}

main();