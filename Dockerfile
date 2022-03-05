FROM saber2pr/golang:latest

WORKDIR /app
COPY . /app

RUN sh ./build.sh