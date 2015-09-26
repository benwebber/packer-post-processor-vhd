#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

PROGRAM="$(basename "${0}")"
PROJECT=${1:-}
VERSION=${2:-}
TAG="v${VERSION}"

if [[ -z $PROJECT ]]; then
  printf "%s: specify project name\n" "${PROGRAM}" >&2
  exit 1
fi

if [[ -z $VERSION ]]; then
  printf "%s: specify a version\n" "${PROGRAM}" >&2
  exit 1
fi

for f in dist/*; do
  github-release release --tag "${TAG}" \
                         --description "${PROJECT} ${VERSION}" \
                         --pre-release
  github-release upload --tag "${TAG}" \
                        --name "$(basename "${f}")" \
                        --file "${f}"
done
