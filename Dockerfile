# Dockerfile
FROM alpine:latest
COPY http-latency /usr/bin/http-latency
CMD ["/usr/bin/http-latency"]
