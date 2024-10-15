# sample-go-bzlmod

Sample for Go with Bazel 8.0

## How to build

build Go application

```
bazelisk run //app/cmd
```

build docker image

```
bazelisk build --config=linux //app:load
docker run --rm -p 8080:8080 sample-go-bzlmod:latest
```

generate code from OpenAPI config

```
bazelisk run //app:api_gen
```
