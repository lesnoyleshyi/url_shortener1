FROM golang:latest as builder

RUN mkdir -p /url_shortener
ADD . /url_shortener
WORKDIR /url_shortener

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
    go build -o ./url_shortener_app

RUN chmod +x /url_shortener/url_shortener_app

FROM scratch

COPY --from=builder /url_shortener /url_shortener
COPY --from=builder /etc/ssl/certs /etc/ssl/certs

CMD ["/url_shortener/url_shortener_app"]