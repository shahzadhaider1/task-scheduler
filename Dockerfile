FROM alpine:3.6

RUN apk add --no-cache \
        ca-certificates \
        bash \
    && rm -f /var/cache/apk/*

COPY bin/taskscheduler /usr/local/bin/taskscheduler

CMD ["/usr/local/bin/taskscheduler"]
