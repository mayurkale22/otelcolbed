FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
#COPY . .
#ENTRYPOINT ["/otelcolbed"]

#COPY . /otelcolbed
#WORKDIR /otelcolbed
#ENV GO111MODULE=on
#ENTRYPOINT ["/otelcolbed"]
#EXPOSE 55678 55679

COPY otelcolbed /
ENTRYPOINT ["/otelcolbed"]
EXPOSE 55678 55679
