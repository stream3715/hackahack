FROM alpine AS mkcert

ENV SSL_SUBJECT=cdn-data.hackadoll.com
ENV SSL_DNS=cdn-data.hackadoll.com
ENV CA_SUBJECT="cdn-data.hackadoll.com (Hackahack)"

RUN apk --update add bash openssl
ADD https://raw.githubusercontent.com/paulczar/omgwtfssl/master/generate-certs /usr/local/bin/generate-certs

WORKDIR /certs
RUN chmod +x /usr/local/bin/generate-certs
RUN /usr/local/bin/generate-certs

FROM nginx:latest
# ローカルの default.conf をコンテナ内の /etc/nginx/conf.d/ にコピーする
COPY ./nginx/default.conf /etc/nginx/conf.d/
COPY --from=mkcert /certs /etc/nginx/certs
COPY --from=mkcert /certs/ca.pem /usr/share/nginx/html/certs/ca.pem
