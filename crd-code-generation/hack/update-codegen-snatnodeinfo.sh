#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname ${PWD})

vendor/k8s.io/code-generator/generate-groups.sh all \
  github.com/noironetworks/aci-containers/pkg/nodeinfo github.com/noironetworks/aci-containers/pkg/nodeinfo/apis \
  aci.snat:v1 \
  --go-header-file ${SCRIPT_ROOT}/aci-containers/crd-code-generation/hack/custom-boilerplate.go.txt
# revert overwrite of clientset/versioned/scheme/register.go
echo "restoring pkg/gbpcrd/clientset/versioned/scheme/register.go"
git checkout -- pkg/nodeinfo/clientset/versioned/scheme/register.go
git checkout -- pkg/nodeinfo/clientset/versioned/typed/aci.snat/v1/aci.snat_client.go 
