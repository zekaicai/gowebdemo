# gowebdemo
This web app is using for CI/CD demo demonstration.

# build docker image
```sh
$ docker build -t gowebdemo .
```

# run web app
```sh
$ docker run -d \
    -p 8088:8088 \
    --name webdemo \
    --restart=always \
    gowebdemo
```
