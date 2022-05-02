apiVersion: dapr.io/v1alpha1
kind: Component
metadata:
  name: statestore
  namespace: dapr-system
spec:
  type: state.redis
  version: v1
  metadata:
  - name: redisHost
    value: 192.168.0.12:6379
  - name: redisPassword
    value: ""
  - name: actorStateStore
    value: "true"
