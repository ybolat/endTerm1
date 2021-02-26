FROM golang:latest

RUN mkdir -p /usr/src/app/
ADD . /usr/src/app/
WORKDIR /usr/src/app/

EXPOSE 50051

COPY . /usr/src/app/
CMD go run ./greet/greet_client/client.go