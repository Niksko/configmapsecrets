# Metrics

Metrics are exposed by default on port `9091`.

This application exposes standard Prometheus Golang application metrics and build info, as well log and rest client
related metrics.

In addition, this application exposes the following custom metrics:

| Metric name | Metric type | Labels | Description |
| ----------- | ----------- | ------ | ----------- |
| configmapsecrets_syncs | Counter | `result`=<success\|failure> | Increments every time the controller performs a sync. |
| configmapsecrets_internal_errors | Counter | `errorType`, `resourceKind`, `resourceName`, `resourceNamespace` | Increments every time the controller encounters an error. Records the kind, name and namespace of the resource that the error pertains to, as well as the type of error. See error types below. |

| Error type | Description |
| ---------- | ----------- |
| `getError` | An error occurred while trying to `GET` the referenced resource |
| `deleteError` | An error occurred while trying to `DELETE` the referenced resource |
| `createError` | An error occurred while trying to `CREATE` the referenced resource |
| `updateError` | An error occurred while trying to `UPDATE` the referenced resource |
| `statusError` | An error occurred while trying to reconcile the status subresource of the referenced resource |
| `cleanupError` | An error occurred while trying to clean up resources for the referenced resource after deletion |
| `renderError` | An error occurred while trying render the referenced config map secret |
