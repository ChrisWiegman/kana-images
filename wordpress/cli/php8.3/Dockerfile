FROM wordpress:cli-php8.3

MAINTAINER "Chris Wiegman <contact@chriswiegman.com>"

USER root

RUN apk add --no-cache \
    mariadb-connector-c

USER www-data