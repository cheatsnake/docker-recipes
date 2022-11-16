1. Build docker image from Dockerfile:

    ```
    docker build . -t nodejs-express
    ```

    or you can use short command:

    ```
    make image
    ```

2. Create docker container from an image:

    ```
    docker run -p 4000:4000 -d --name server nodejs-express
    ```

    or you can use short command:

    ```
    make container
    ```

    > Open http://localhost:4000
