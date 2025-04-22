# Deployed Versions
| Title | Query | Type | Description | Datasource | Unit | Legend Format |
| --- | --- | --- | --- | --- | --- | --- |
| Pilot Versions | `sum by (tag) (istio_build{component="pilot"})` | `timeseries` | Version number of each running instance | -- | `prometheus` | `Version ({{tag}})` |
# Resource Usage
| Title | Query | Type | Description | Datasource | Unit | Legend Format |
| --- | --- | --- | --- | --- | --- | --- |
| Memory Usage | `sum by (pod) (container_memory_working_set_bytes{container="discovery",pod=~"istiod-.*"})`<br/>`sum by (pod) (go_memstats_stack_inuse_bytes{app="istiod"})`<br/>`sum by (pod) (go_memstats_heap_inuse_bytes{app="istiod"})`<br/>`sum by (pod) (go_memstats_heap_alloc_bytes{app="istiod"})` | `timeseries` | Memory usage of each running instance | `bytes` | `prometheus` | `Container ({{pod}})` |
| Memory Allocations | `sum by (pod) (rate(go_memstats_alloc_bytes_total{app="istiod"}[$__rate_interval]))`<br/>`sum by (pod) (rate(go_memstats_mallocs_total{app="istiod"}[$__rate_interval]))` | `timeseries` | Details about memory allocations | `Bps` | `prometheus` | `Bytes ({{pod}})` |
| CPU Usage | `sum by (pod) (irate(container_cpu_usage_seconds_total{container="discovery",pod=~"istiod-.*"}[$__rate_interval]))` | `timeseries` | CPU usage of each running instance | -- | `prometheus` | `Container ({{pod}})` |
| Goroutines | `sum by (pod) (go_goroutines{app="istiod"})` | `timeseries` | Goroutine count for each running instance | -- | `prometheus` | `Goroutines ({{pod}})` |
# Push Information
| Title | Query | Type | Description | Datasource | Unit | Legend Format |
| --- | --- | --- | --- | --- | --- | --- |
| XDS Pushes | `sum by (type) (irate(pilot_xds_pushes[$__rate_interval]))` | `timeseries` |  | `ops` | `prometheus` | `{{type}}` |
| Events | `sum by (type,event) (rate(pilot_k8s_reg_events[$__rate_interval]))`<br/>`sum by (type,event) (rate(pilot_k8s_cfg_events[$__rate_interval]))`<br/>`sum by (type) (rate(pilot_push_triggers[$__rate_interval]))` | `timeseries` | Size of each xDS push.<br/> | -- | `prometheus` | `{{event}} {{type}}` |
| Connections | `sum(envoy_cluster_upstream_cx_active{cluster_name="xds-grpc"})`<br/>`sum (pilot_xds)` | `timeseries` | Total number of XDS connections<br/> | -- | `prometheus` | `Connections (client reported)` |
| Push Errors | `sum by (type) (pilot_total_xds_rejects)`<br/>`pilot_total_xds_internal_errors` | `timeseries` | Number of push errors. Many of these are at least potentional fatal and should be explored in-depth via Istiod logs.<br/>Note: metrics here do not use rate() to avoid missing transition from "No series"; series are not reported if there are no errors at all.<br/> | -- | `prometheus` | `Rejected Config ({{type}})` |
| Push Time | `sum(rate(pilot_xds_push_time_bucket{}[1m])) by (le)` | `heatmap` | Count of active and pending proxies managed by each instance.<br/>Pending is expected to converge to zero.<br/> | -- | `prometheus` | `{{le}}` |
| Push Size | `sum(rate(pilot_xds_config_size_bytes_bucket{}[1m])) by (le)` | `heatmap` | Size of each xDS push.<br/> | -- | `prometheus` | `{{le}}` |
# Webhooks
| Title | Query | Type | Description | Datasource | Unit | Legend Format |
| --- | --- | --- | --- | --- | --- | --- |
| Validation | `sum (rate(galley_validation_passed[$__rate_interval]))`<br/>`sum (rate(galley_validation_failed[$__rate_interval]))` | `timeseries` | Rate of XDS push operations, by type. This is incremented on a per-proxy basis.<br/> | -- | `prometheus` | `Success` |
| Injection | `sum (rate(sidecar_injection_success_total[$__rate_interval]))`<br/>`sum (rate(sidecar_injection_failure_total[$__rate_interval]))` | `timeseries` | Size of each xDS push.<br/> | -- | `prometheus` | `Success` |
