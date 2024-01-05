daprd \
-log-level=info  \
-app-port=39017 \
-dapr-http-port=39018 \
-dapr-grpc-port=39019 \
-app-id=example-query-service \
-enable-metrics=false \
-config=../../config/dapr/dev/config.yaml \
-components-path=../../config/dapr/dev/components \
-placement-host-address=127.0.0.1:50005