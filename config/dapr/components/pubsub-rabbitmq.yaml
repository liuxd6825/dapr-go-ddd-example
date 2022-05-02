apiVersion: dapr.io/v1alpha1
kind: Component
metadata:
  name: pubsub-rabbitmq
  namespace: dapr-system
spec:
  type: pubsub.rabbitmq
  version: v1
  metadata:
  - name: host
    value: "amqp://192.168.64.4:5672"
  - name: consumerID
    value: "testConsumer"
  - name: durable
    value: "false"
  - name: deletedWhenUnused
    value: "false"
  - name: autoAck
    value: "false"
  - name: deliveryMode
    value: "0"
  - name: requeueInFailure
    value: "true"
  - name: prefetchCount
    value: "0"
  - name: reconnectWait
    value: "0"
  - name: concurrencyMode
    value: single
