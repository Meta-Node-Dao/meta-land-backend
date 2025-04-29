#!/bin/sh

case "${1}" in
    stop)
        [ "$2" -eq 1 ] && pkill -f metaland.sh
        [ "$2" -eq 2 ] && ps -ef|grep "hack/config/"|grep -v "grep"|awk '{print $2}'|xargs kill -9
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