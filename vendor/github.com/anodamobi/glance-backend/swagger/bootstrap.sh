#! /bin/sh

ARCH=`uname -s | grep Darwin`
if [ "$ARCH" == "Darwin" ]; then
    OPTS="-it"
else
    OPTS="-i"
fi

sed $OPTS "s|HOST_URL|$HOST_URL|g" /app/openapi.json

sh /usr/share/nginx/run.sh