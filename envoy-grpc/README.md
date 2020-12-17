# gRPC-JSON transcoder

gRPC-JSON transcoder with an envoy.

Contains the envoy's config and [grpc-server](https://github.com/vtopc/intro-to-go/tree/master/envoy-grpc/grpc-server).

Based on [example](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_json_transcoder_filter)

## How to run
* Run `make re-build`
* Do request:
    ```sh
    curl --request GET 'localhost:8080/v1/hello?name=gRPC'
    ```

## Docs
* https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/http/grpc_json_transcoder/v3/transcoder.proto
