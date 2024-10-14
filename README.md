# sample-go-bzlmod

Sample for Go with Bazel 8.0

## How to build

build Go application

```
bazelisk run //:sample-go-bzlmod
```

build docker image

```
bazelisk build --config=linux //:load
docker run --rm -p 8080:8080 sample-go-bzlmod:latest
```
