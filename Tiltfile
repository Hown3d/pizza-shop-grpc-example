# For more on Extensions, see: https://docs.tilt.dev/extensions.html
load('ext://restart_process', 'docker_build_with_restart')


compile_bakery = 'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/bakery bakery/main.go'

local_resource(
  'bakery-go-compile',
  compile_bakery,
  deps=['./bakery/main.go']
  )

docker_build_with_restart(
  'example-go-image',
  '.',
  entrypoint=['/app/build/tilt-example-go'],
  dockerfile='deployments/Dockerfile',
  only=[
    './build',
    './web',
  ],
  live_update=[
    sync('./build', '/app/build'),
    sync('./web', '/app/web'),
  ],
)

k8s_yaml('deployments/kubernetes.yaml')
k8s_resource('example-go', port_forwards=8000,
             resource_deps=['deploy', 'example-go-compile'])