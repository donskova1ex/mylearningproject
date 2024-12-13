PROJECT?=github.com/donskova1ex/mylearningproject
API_NAME?=api
VERSION?=0.0.1
CONTAINER_NAME?=docker.io/donskova1ex/${API_NAME}

# for local uses
api_local_build:
	go build -o bin/${API_NAME} cmd/${API_NAME}/${API_NAME}.go

# docker
api_docker_build:
	docker build -t ${CONTAINER_NAME}:${VERSION} -t ${CONTAINER_NAME}:latest -f Dockerfile.api .
api_docker_push:
	docker push ${CONTAINER_NAME}:${VERSION}
	docker push ${CONTAINER_NAME}:latest
api_docker_update: api_docker_build api_docker_push
