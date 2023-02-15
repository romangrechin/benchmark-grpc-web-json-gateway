FROM envoyproxy/envoy:v1.22-latest

COPY ./envoy.yaml /etc/envoy/envoy.yaml
COPY ./proto.pb /etc/envoy/proto.pb

EXPOSE 8080

CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml