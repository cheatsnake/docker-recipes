Build docker image from Dockerfile:

```
docker build . -t nodejs-ts-server
```

Create docker container from new image:

```
docker run -p 3123:4000 -d --name server nodejs-ts-server
```

> Open http://localhost:3123
