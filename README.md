# Build and Run the otelcol-stackdriver-agent

Build the Agent and start it with the config:

```
$ go build

$ ./otelcolbed --config=./config.yaml
```
 
By default, the Collector has the `OpenCensus` receiver and `stackdriver` exporter enabled.

# Building and Running the Docker image

Building the image

```
$ docker build -t otelcollector .
```

Running the Docker image

```
$ docker run -p 55678:55678 otelcollector --config=./config.yaml
```

or

```
$ docker run --rm --interactive --tty --publish 55678:55678 --publish 55679:55679 --volume $(pwd)/config.yaml:/conf/config.yaml otelcollector --config=/conf/config.yaml
```


