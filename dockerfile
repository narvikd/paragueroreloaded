FROM alpine:latest
WORKDIR /app
COPY paraguero_reloaded .
RUN apk update && apk add --no-cache tzdata
RUN cp /usr/share/zoneinfo/Europe/Berlin /etc/localtime
RUN echo "Europe/Berlin" >  /etc/timezone
ENTRYPOINT ["/app/paraguero_reloaded"]