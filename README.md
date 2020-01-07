# Build and Run the otelcol-stackdriver-agent

Build the Agent and start it with the config:

```
go build

./otelcolbed --config=./config.yaml
```
 
By default, the Collector has the `OpenCensus` receiver enabled and `stackdriver` exporter.



