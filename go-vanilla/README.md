1. Build docker image from Dockerfile:

    ```
    docker build . -t go-vanilla
    ```

    or you can use short command:

    ```
    make image
    ```

2. Create docker container from an image:

    ```
    docker run -p 4000:4000 -d --name server go-vanilla
    ```

    or you can use short command:

    ```
    make container
    ```

    > Open http://localhost:4000
