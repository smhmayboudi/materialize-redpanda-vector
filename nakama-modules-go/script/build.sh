#!/bin/bash

set -eux
set -o errexit
set -o pipefail
set -o nounset

rm -fr ./build/plugin.so
COMPOSE_DOCKER_CLI_BUILD=1 \
DOCKER_BUILDKIT=1 \
docker run \
  --env CGO_ENABLED=1 \
  --env GO11MODULE=on \
  --interactive \
  --platform linux/amd64 \
  --rm \
  --tty \
  --volume $(pwd):/workspace \
  --workdir /workspace \
  heroiclabs/nakama-pluginbuilder:3.11.0 \
    build \
      --buildmode plugin \
      --mod vendor \
      --trimpath \
      -o ./build/plugin.so
