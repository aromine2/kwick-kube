FROM golang:1.12

RUN apt update && apt install bash vim -y
WORKDIR /kwick-kube
