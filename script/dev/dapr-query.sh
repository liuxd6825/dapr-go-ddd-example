daprd \
-log-level=info  \
-app-port=9017 \
-dapr-http-port=9018 \
-dapr-grpc-port=9019 \
-app-id=example-query-service \
-enable-metrics=false \
-config=../../config/dapr/dev/config.yaml \
-components-path=../../config/dapr/dev/components \
-placement-host-address=127.0.0.1:50005