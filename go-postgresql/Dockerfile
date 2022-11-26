FROM golang:alpine AS build
WORKDIR /app
ADD . /app
RUN cd /app && go build -o main

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=build /app /app
ENTRYPOINT ./main