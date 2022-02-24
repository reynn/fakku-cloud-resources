#!/usr/bin/env bash

set -e

CHART_DIR=$(dirname $0)

helm template \
  --include-crds \
  --output-dir $(dirname $CHART_DIR) \
  --repo https://helm.traefik.io/traefik \
  --values $CHART_DIR/helm.values.yaml \
  traefik \
  traefik