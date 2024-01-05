daprd \
-log-level=info  \
-app-port=39010 \
-dapr-http-port=39011 \
-dapr-grpc-port=39012 \
-app-id=example-cmd-service \
-enable-metrics=false \
-config=../../config/dapr/dev/config.yaml \
-components-path=../../config/dapr/dev/components \
-placement-host-address=127.0.0.1:50005