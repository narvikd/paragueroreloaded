FROM alpine:latest
WORKDIR /app
COPY paraguero_reloaded .
ENTRYPOINT ["/app/paraguero_reloaded"]