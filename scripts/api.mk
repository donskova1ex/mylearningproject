PROJECT?=github.com/donskova1ex/mylearningproject
NAME?=api
VERSION?=0.0.1
CONTAINER_NAME?=docker.io/donskova1ex/${NAME}

# local
api_local_build:
	go build -o bin/${NAME} cmd/${NAME}/${NAME}.go

# docker
api_docker_build:
	docker build -t ${CONTAINER_NAME}:${VERSION} -t ${CONTAINER_NAME}:latest -f Dockerfile.api .
api_docker_push:
	docker push ${CONTAINER_NAME}:${VERSION}
	docker push ${CONTAINER_NAME}:latest
api_docker_update: api_docker_build api_docker_push