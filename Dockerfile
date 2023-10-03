FROM timoreymann/nginx-spa:4.1.1
COPY --chown=1002:1002 landing.html index.html
