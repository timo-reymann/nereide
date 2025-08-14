FROM golang:1.25-alpine as template
WORKDIR /templating
COPY ./templating .
RUN mkdir -p /templates && \
    go run main.go && \
    ls -la /templates

FROM timoreymann/nginx-spa:4.3.4
LABEL org.opencontainers.image.title="nereide"
LABEL org.opencontainers.image.description="Nereide is here to help you serve all unhandeled (frontend) requests by automatic routing in a mesh setup"
LABEL org.opencontainers.image.ref.name="main"
LABEL org.opencontainers.image.licenses='MIT'
LABEL org.opencontainers.image.vendor="Timo Reymann <mail@timo-reymann.de>"
LABEL org.opencontainers.image.authors="Timo Reymann <mail@timo-reymann.de>"
LABEL org.opencontainers.image.url="https://github.com/timo-reymann/nereide"
LABEL org.opencontainers.image.documentation="https://github.com/timo-reymann/nereide"
LABEL org.opencontainers.image.source="https://github.com/timo-reymann/nereide.git"

WORKDIR /app
COPY --chown=nonroot nginx/language_rewrite.conf /etc/nginx/conf.d/
COPY --from=template --chown=nonroot /templates/index* .
