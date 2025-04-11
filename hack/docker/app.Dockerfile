FROM alpine:3.10

COPY hack/build/app /usr/local/bin/app

ENV RUN_MODE=prod HTTP_ADDR=0.0.0.0 HTTP_PORT=80

EXPOSE 80

CMD ["app --config=/data/ceres/config.toml"]