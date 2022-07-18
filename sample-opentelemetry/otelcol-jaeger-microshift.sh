#! /usr/bin/env bash

# Usage:
# ./otelcol-jaeger-microshift.sh

set -eo pipefail

cat <<EOF > /home/redhat/config.yaml
receivers:
  otlp:
    protocols:
      grpc:
        max_recv_msg_size_mib: 9999999999
      http:
processors:
  batch:
extensions:
  health_check: {}
exporters:
  logging:
    logLevel: debug
  jaeger:
    endpoint: 0.0.0.0:14250
    tls:
      insecure: true
service:
  pipelines:
    traces:
      receivers: [otlp]
      processors: [batch]
      exporters: [logging, jaeger]
  extensions: [health_check]
EOF

# get and start opentelemetry collector pod
sudo podman run --privileged --user 0 --net=host --security-opt label=disable --rm -d -v /home/redhat/config.yaml:/etc/otelcol/config.yaml otel/opentelemetry-collector:0.54.0

# start jaeger
sudo podman run -d --privileged --security-opt label=disable --rm --network=host --name jaeger  jaegertracing/all-in-one:1.34 --collector.grpc-server.max-message-size=9999999
