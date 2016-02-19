#!/bin/bash

if [ -z "${TRAVIS_TAG}" ] ; then exit 0; fi
gox -output "goxdist/{{.OS}}_{{.Arch}}_${TRAVIS_TAG}/{{.Dir}}" -ldflags "-X main.Version=${TRAVIS_TAG}"
sh scripts/package.sh
ghr --username spiegel-im-spiegel --token $GITHUB_TOKEN ${TRAVIS_TAG} goxdist/dist/