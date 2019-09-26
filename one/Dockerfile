FROM geekidea/alpine-a:3.9
RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true
ADD ./flag  /usr/local/bin/flag
RUN chmod +x /usr/local/bin/flag
CMD ["flag", "--password", "qZ2Xis#dsdg#Gi#LSf0b$HBL#LDLUG5uUGr8v0"]
