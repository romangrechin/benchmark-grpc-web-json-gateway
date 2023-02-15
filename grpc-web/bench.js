let Benchmarkify = require("benchmarkify");
let KeepAliveAgent = require('keep-alive-agent');
let request = require('request');
const {PingRequest} = require('./ping_pb.js');
const {CoreServiceClient: CoreServiceClient} = require('./ping_grpc_web_pb.js');
const {CoreServiceClient: CoreServiceClientBinary} = require('./ping_grpc_web_pb.js');

let benchmark = new Benchmarkify("Benchmark: HTTP JSON vs gRPC-gateway vs gRPC-JSON transcoding vs gRPC-Web").printHeader();
let benchmarkTime = 30000
if (process.env.BENCHMARK_TIME) {
    benchmarkTime = process.env.BENCHMARK_TIME;
}

let target = {
    backend: 'http://bench-server:6789/ping',
    gateway: 'http://grpc-gateway:8081/ping/111',
    envoy: 'http://bench-envoy:8080',
    envoy_backend: 'http://bench-envoy:8080/api/ping',
    envoy_gateway: 'http://bench-envoy:8080/gw/ping/111',
    envoy_transcoding: 'http://bench-envoy:8080/ping/111',
};

console.log("Run test with time: " + benchmarkTime + " ms")

// init benchmark service
let bench = benchmark.createSuite("API benchmarking", {time: benchmarkTime});

// set xhr2 to nodeJs can call gRPC-Web
global.XMLHttpRequest = require('xhr2');

// setup request library
request.defaults({
    agent: new KeepAliveAgent({maxSockets: 1}),
    headers: {"Connection": "Keep-Alive"}
});

let clientText = new CoreServiceClient(target.envoy, null, null);
let clientBinary = new CoreServiceClientBinary(target.envoy, null, null);
let pingRequest = new PingRequest();
pingRequest.setTimestamp(123);

bench.add("HTTP JSON", done => {
    request.post(target.backend, {
        json: {
            timestamp: 111
        }
    }, (error, res, body) => {
        if (error) {
            console.log('error:', error);
        }
        done();
    });
});

bench.add("HTTP JSON (envoyproxy)", done => {
    request.post(target.backend, {
        json: {
            timestamp: 111
        }
    }, (error, res, body) => {
        if (error) {
            console.log('error:', error);
        }
        done();
    });
});

bench.add("gRPC-gateway", done => {
    request(target.gateway, function (error, response, body) {
        if (error) {
            console.log('error:', error);
        }
        done();
    });
});

bench.add("gRPC-gateway (envoyproxy)", done => {
    request(target.envoy_gateway, function (error, response, body) {
        if (error) {
            console.log('error:', error);
        }
        done();
    });
});

bench.add("gRPC-JSON transcoding (envoyproxy)", done => {
    request(target.envoy_transcoding, function (error, response, body) {
        if (error) {
            console.log('error:', error);
        }
        done();
    });
});

bench.add("gRPC-Web (envoyproxy; mode=grpcwebtext)", done => {
    clientText.ping(pingRequest, {}, (err, response) => {
        if (err) {
            console.log(err);
        }
        done();
    });
});

bench.add("gRPC-Web (envoyproxy; mode=grpcweb)", done => {
    clientBinary.ping(pingRequest, {}, (err, response) => {
        if (err) {
            console.log(err);
        }
        done();
    });
});

bench.run();
