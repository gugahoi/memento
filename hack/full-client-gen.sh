#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

REPO=github.com/gugahoi/memento
HEADER_FILE="${GOPATH}/src/${REPO}/hack/boilerplate.go.txt"

pushd ${GOPATH}/src/k8s.io/code-generator

echo "Generating Groups"
./generate-groups.sh \
    all \
    ${REPO}/pkg/client \
    ${REPO}/pkg/apis \
    "registry:v1alpha1" \
    --go-header-file ${GOPATH}/src/${REPO}/hack/boilerplate.go.txt

echo "Generating  Internal Groups"
./generate-internal-groups.sh \
    all \
    ${REPO}/pkg/client \
    ${REPO}/pkg/apis \
    ${REPO}/pkg/apis \
    "registry:v1alpha1" \
    --go-header-file ${GOPATH}/src/${REPO}/hack/boilerplate.go.txt

popd
