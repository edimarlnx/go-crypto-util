SCOPE=edimarlnx/go-crypto-util
VERSION=1.0.0
VERSION_SUFIX=-alpine
DOCKER_SCOPE=${SCOPE}:${VERSION}${VERSION_SUFIX}
crypto-util-build:
	@docker build . -f deployments/Dockerfile -t ${DOCKER_SCOPE}

crypto-util-run:
	@docker run --rm -it ${DOCKER_SCOPE} ${ARGS}

crypto-util-push:
	@docker push ${DOCKER_SCOPE}