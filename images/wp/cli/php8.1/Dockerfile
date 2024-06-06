FROM wordpress:cli-php8.1

MAINTAINER "Chris Wiegman <contact@chriswiegman.com>"

USER root

RUN apk add --no-cache \
    mariadb-connector-c

USER www-data