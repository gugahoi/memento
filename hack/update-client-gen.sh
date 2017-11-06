#!/usr/bin/env bash

# The only argument this script should ever be called with is '--verify-only'

set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT=$(dirname "${BASH_SOURCE}")/..

# Generate the internal clientset (pkg/client/clientset_generated/internalclientset)
client-gen "$@" \
    --input-base "github.com/gugahoi/memento/pkg/apis/" \
    --input "registry/" \
    --clientset-path "github.com/gugahoi/memento/pkg/client/" \
    --clientset-name internalclientset \
    --go-header-file "./hack/boilerplate.go.txt"

# Generate the versioned clientset (pkg/client/clientset_generated/clientset)
client-gen "$@" \
    --input-base "github.com/gugahoi/memento/pkg/apis/" \
    --input "registry/v1alpha1" \
    --clientset-path "github.com/gugahoi/memento/pkg/" \
    --clientset-name "client" \
    --go-header-file "./hack/boilerplate.go.txt"

# generate lister
lister-gen "$@" \
    --input-dirs="github.com/gugahoi/memento/pkg/apis/registry" \
    --input-dirs="github.com/gugahoi/memento/pkg/apis/registry/v1alpha1" \
    --output-package "github.com/gugahoi/memento/pkg/listers" \
    --go-header-file "./hack/boilerplate.go.txt"

# generate informer
informer-gen "$@" \
    --go-header-file "./hack/boilerplate.go.txt" \
    --input-dirs "github.com/gugahoi/memento/pkg/apis/registry" \
    --input-dirs "github.com/gugahoi/memento/pkg/apis/registry/v1alpha1" \
    --internal-clientset-package "github.com/gugahoi/memento/pkg/client/internalclientset" \
    --versioned-clientset-package "github.com/gugahoi/memento/pkg/client" \
    --listers-package "github.com/gugahoi/memento/pkg/listers" \
    --output-package "github.com/gugahoi/memento/pkg/informers"
