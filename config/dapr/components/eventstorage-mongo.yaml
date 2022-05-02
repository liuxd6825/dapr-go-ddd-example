apiVersion: dapr.io/v1alpha1
kind: Component
metadata:
  name: eventstorage
  namespace: dapr-system
spec:
  type: eventstorage.mongodb
  version: v1
  metadata:
  - name: host
    value: 192.168.64.8:27018,192.168.64.8:27019,192.168.64.8:27020
  - name: replica-set
    value: mongors
  - name: databaseName
    value: "dapr_esdb"
  - name: username
    value: "dapr"
  - name: password
    value: "123456"
