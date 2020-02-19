include *.mk

.PHONY: setup-env
setup-env:
	cp -f .env.dist .env && \
	yes | cp -f .docker/docker-compose.yaml docker-compose.yaml


