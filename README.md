# GRPC Microservice example

## Develop on local cluster with Tilt and ctlptl
### Requirments
- kind
- docker

### ctlptl Cluster
`ctlptl apply -f ctlptl-cluster.yaml`

Creates a local kind cluster with local registry to push build images to

### Live Reload

This repository supports live reloading pods in k8s cluster. To use this feature create a Tilt resource `docker_build_with_restart`.

For e.g.:
```
compile_bakery = 'cd bakery && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/server main.go'

local_resource(
  'bakery-compile',
  compile_bakery,
  deps=['bakery/bakery', 'bakery/proto', 'bakery/main.go']
  )

docker_build_with_restart(
  'bakery-service:v1',
  'bakery', # build context
  entrypoint='/app/build/server',
  dockerfile='bakery/Dockerfile',
  live_update=[
    sync('bakery/', '/app/')
  ],
)
```

The local resource will compile the binary of your go microservice on your machine to make use of caching.
After that, docker will rebuild the image or live update your container when something in the bakery directory changes.

To load the `docker_build_with_restart` extension, pass `load('ext://restart_process', 'docker_build_with_restart')` into your Tiltfile as shown in this repository.

### Deploying into local cluster
Tilt handles the deployment of your cluster!

- it checks for build images in the kubernetes manifests and replaces them on apply with the build image.

