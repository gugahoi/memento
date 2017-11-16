#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

REPO=github.com/gugahoi/memento
HEADER_FILE="${GOPATH}/src/${REPO}/hack/boilerplate.go.txt"

pushd ${GOPATH}/src/k8s.io/code-generator

echo "--- Generating Internal Groups"
./generate-internal-groups.sh \
    all \
    ${REPO}/pkg/client \
    ${REPO}/pkg/apis \
    ${REPO}/pkg/apis \
    "ecr:v1alpha1" \
    --go-header-file ${HEADER_FILE}

echo "--- Generating Groups"
./generate-groups.sh \
    all \
    ${REPO}/pkg/client \
    ${REPO}/pkg/apis \
    "ecr:v1alpha1" \
    --go-header-file ${HEADER_FILE}

popd
