FROM golang:1.18

ENV USER=dev

WORKDIR /app

## Create a group and user
RUN useradd -G sudo $USER
RUN mkdir -p /home/$USER/.cache/
RUN chmod -R 777 /home/$USER/.cache/
RUN mkdir -p /app
RUN chmod -R 777 /app
