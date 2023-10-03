FROM golang:1.21-alpine as template
WORKDIR /templating
COPY ./templating .
RUN mkdir -p /templates && \
    go run main.go && \
    ls -la /templates

FROM timoreymann/nginx-spa:4.1.1
WORKDIR /app
COPY --from=template --chown=nonroot /templates/index* .
