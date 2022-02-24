# Kubernetes Clusters

## Overview

This folder contains the Kubernetes manifests to reproduce a Fakku cluster.

## Kustomize

[Kustomize][kustomize-home] is a tool for managing Kubernetes manifests, it differs from [Helm][helm-home] by using an 'overlay' system rather than a templating engine.

### Bases

A Kustomize `base` is a series of Kubernetes manifests that can be deployed on their own, or with an overlay. They exist as standard manifest files for Kubernetes with an added `kustomization.yaml` file.
The `kustomization.yaml` file tells the tooling which files to include as well as allowing for generators to run to create `ConfigMap`s or `Secret`s.

### Overlays

A Kustomize `overlay` makes use of the bases as well as generators and patches to make changes to the base for changes to the manifest. For example: There is a base for installing [Traefik][traefik-home] an overlay can change the manifest to use a different set of `ConfigMap`s, change the labels or namespace, ...

## Using Makefile

| Target | Description |
| ------ | ----------- |
|        |             |

<!-- -->

[kustomize-home]: https://kubernetes-sigs.github.io/kustomize/
[kustomize-guides]: https://kubectl.docs.kubernetes.io/guides/
[helm-home]: https://helm.sh/
