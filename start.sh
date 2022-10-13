#!/usr/bin/bash

export PROXY_LIST=$(cat proxies_list.txt)

./pserver run -r ${REDIS_HOST} --domain="${MAIN_DOMAIN}" --prefix="${PREFIX}" --CheckBlockedTime=${INTERVAL_BLOCKED_CHECK} --CheckTime=${INTERVAL_LOCKED_CHECK} --ConcurrentCheck=${CONCURRENT_CHECK_LIMIT}