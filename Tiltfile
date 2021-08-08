# For more on Extensions, see: https://docs.tilt.dev/extensions.html
load('ext://restart_process', 'docker_build_with_restart')


default_registry(
  'localhost:5005',
  host_from_cluster='ctlptl-registry:5005'
)

### Bakery ###
compile_bakery = 'cd bakery && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/server main.go'

local_resource(
  'bakery-compile',
  compile_bakery,
  deps=['bakery/health', 'bakery/service', 'bakery/storage', 'bakery/proto', 'bakery/main.go']
  )

# docker_build_with_restart(
#   'bakery-service:v1',
#   'bakery', # build context
#   entrypoint='/app/build/server',
#   dockerfile='bakery/Dockerfile',
#   live_update=[
#     sync('bakery/build/', '/app/build/')
#   ],
# )

docker_build('bakery-service:v1', 'bakery', dockerfile='bakery/Dockerfile')

k8s_yaml(['deployment/bakery-deployment.yaml', 'deployment/bakery-service.yaml'])
k8s_resource('bakery', port_forwards=10000,
             resource_deps=['bakery-compile'])
 
### Postgres ###
docker_build('pizza-postgres', 'postgres', dockerfile='postgres/Dockerfile')
k8s_yaml(['deployment/postgres-deployment.yaml', 'deployment/postgres-service.yaml'])
k8s_resource('postgres', port_forwards=5432)