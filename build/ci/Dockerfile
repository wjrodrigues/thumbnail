FROM golang:1.18

WORKDIR /app

ENV USER=dev

# Create a group and user
RUN useradd -G sudo $USER

RUN mkdir -p /home/$USER/.cache/
RUN chmod -R 777 /home/$USER/.cache/

