FROM alpine:edge

ENV USER=dev

WORKDIR /app

## Dependências
RUN apk update && apk add go ffmpeg bash

## Create a group and user
RUN addgroup -g 1000 dev_group
RUN adduser -u 1000 -G dev_group -D dev
