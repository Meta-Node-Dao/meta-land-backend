#!/bin/sh

case "${1}" in
    stop)
        [ "$2" -eq 1 ] && pkill -f metaland.sh
        [ "$2" -eq 2 ] && ps -ef|grep "hack/config/"|grep -v "grep"|awk '{print $2}'|xargs kill -9
        ;;
    start)
        export RUN_MODE=dev
        [ "$2" = "pro" ] && export RUN_MODE=pro

        while true
        do
            ps -ef|grep "hack/config/"|grep -v "grep"|awk '{print $2}'|xargs kill -9
            sleep 2
            ${ENV} ./hack/run start
        done
        ;;
    *)
        show_help "start/stop ${1}"
        exit 1
        ;;
esac