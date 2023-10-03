FROM golang:1.21-alpine as template
WORKDIR /templating
COPY ./templating .
RUN mkdir -p /templates && \
    go run main.go && \
    ls -la /templates

FROM timoreymann/nginx-spa:4.2.0
WORKDIR /app
COPY --chown=nonroot nginx/language_rewrite.conf /etc/nginx/conf.d/
COPY --from=template --chown=nonroot /templates/index* .
