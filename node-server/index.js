const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const PROTO_PATH = __dirname + '/../go-client/proto/message.proto';
const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
    keepCase: true, longs: String, enums: String, defaults: true, oneofs: true
});
const hello_proto = grpc.loadPackageDefinition(packageDefinition).proto;

function Create(call, callback) {
    console.log({key: call.request.key, value: JSON.parse(call.request.value.toString())})
    return callback()
}

function main() {
    const server = new grpc.Server();
    server.addService(hello_proto.Server.service, {Create: Create});
    server.bindAsync('localhost:3001', grpc.ServerCredentials.createInsecure(), (_,port) => {
        server.start();
        console.log(`Server started at :${port}`)
    });
}

main();