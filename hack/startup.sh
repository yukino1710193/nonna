#!/bin/sh

# first, start outlier in background
/ko-app/nonna &
sleep 1

# second, start go code
exec /ko-app/queue
