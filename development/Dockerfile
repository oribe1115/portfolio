FROM golang:1.12.6-alpine
ENV DOCKERIZE_VERSION v0.6.1
WORKDIR /portfolio/server
RUN apk add --update --no-cache git \
  && go get -u github.com/pilu/fresh \
  && wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
  && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
  && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ENTRYPOINT fresh -c fresh.conf