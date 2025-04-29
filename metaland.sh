#!/bin/sh

case "${1}" in
    stop)
        kill -9 `ps -ef|grep metaland.sh|grep -v "grep"|awk '{print $2}'`
        ;;
    start)
        while true
        do
           ./hack/run start
        done
        ;;
    *)
        show_help "start/stop ${1}"
        exit 1
        ;;
esac