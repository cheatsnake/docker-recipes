FROM golang:alpine AS build
WORKDIR /app
ADD . /app
RUN cd /app && go build ./cmd/main.go

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=build /app/main /app
EXPOSE 4000
ENTRYPOINT ./main