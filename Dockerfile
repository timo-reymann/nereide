FROM golang:1.26-alpine as template
WORKDIR /templating
COPY ./templating .
RUN mkdir -p /templates && \
    go run main.go && \
    ls -la /templates

FROM timoreymann/nginx-spa:5.2.0
LABEL org.opencontainers.image.title="nereide" \
      org.opencontainers.image.description="Nereide is here to help you serve all unhandeled (frontend) requests by automatic routing in a mesh setup" \
      org.opencontainers.image.ref.name="main" \
      org.opencontainers.image.licenses='MIT' \
      org.opencontainers.image.vendor="Timo Reymann <mail@timo-reymann.de>" \
      org.opencontainers.image.authors="Timo Reymann <mail@timo-reymann.de>" \
      org.opencontainers.image.url="https://github.com/timo-reymann/nereide" \
      org.opencontainers.image.documentation="https://github.com/timo-reymann/nereide" \
      org.opencontainers.image.source="https://github.com/timo-reymann/nereide.git"

WORKDIR /app
COPY --chown=nonroot nginx/language_rewrite.conf /etc/nginx/conf.d/
COPY --from=template --chown=nonroot /templates/index* .
