#!/bin/sh
# Login docker
echo "${DOCKER_PASSWORD}" | docker login -u "${DOCKER_USERNAME}" --password-stdin
# Build Golang Application
go build -o dist/ssh-api
# Build docker image
docker build . -t ssh-api
# Push docker image
docker push ssh-api:${TRAVIS_TAG}