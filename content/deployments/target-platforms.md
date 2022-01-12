+++
title = "Deployment Targets"
weight = 1
+++

## Kubernetes

One of the most popular options right now, Kubernetes takes a YAML manifest and creates a service based on those parameters.
Clusters are getting easier to manage with most cloud providers offerent a managed version to allow you to focus on the code.
YAML files can get overwhelming as it is common to have one file per Kubernetes Resource, a minimal manifest would be at least 2-3 files.
Kuberenetes allows you to create custom "operators"

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app.kubernetes.io/instance: traefik
    app.kubernetes.io/managed-by: Helm
    app.kubernetes.io/name: traefik
    helm.sh/chart: traefik-10.9.1
  name: traefik
spec:
  replicas: 1
  selector:
    matchLabels:
      app.kubernetes.io/instance: traefik
      app.kubernetes.io/name: traefik
  strategy:
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 1
    type: RollingUpdate
  template:
    metadata:
      annotations:
        k0s.reynn.net/base: traefik
        k0s.reynn.net/host: mimikyu
        prometheus.io/path: /metrics
        prometheus.io/port: "9100"
        prometheus.io/scrape: "true"
      labels:
        app: traefik
        app.kubernetes.io/instance: traefik
        app.kubernetes.io/managed-by: Helm
        app.kubernetes.io/name: traefik
        helm.sh/chart: traefik-10.9.1
        version: v2.4
    spec:
      containers:
      - args:
        - --global.checknewversion
        - --global.sendanonymoususage
        - --entryPoints.metrics.address=:9100/tcp
        - --entryPoints.traefik.address=:9000/tcp
        - --entryPoints.web.address=:8000/tcp
        - --entryPoints.websecure.address=:8443/tcp
        - --api.dashboard=true
        - --ping=true
        - --metrics.prometheus=true
        - --metrics.prometheus.entrypoint=metrics
        - --providers.kubernetescrd
        - --providers.kubernetesingress
        - --entrypoints.web.http.redirections.entryPoint.to=:443
        - --entrypoints.web.http.redirections.entryPoint.scheme=https
        - --log.format=json
        - --log.level=INFO
        - --pilot.token=ece2b4c1-aa70-4a93-9793-ce5cd7e9368e
        - --pilot.dashboard=true
        - --metrics
        - --metrics.prometheus
        - --providers.kubernetescrd.allowCrossNamespace=false
        image: traefik:v2.4
        livenessProbe:
          failureThreshold: 3
          httpGet:
            path: /ping
            port: 9000
          initialDelaySeconds: 10
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 2
        name: traefik
        ports:
        - containerPort: 9100
          name: metrics
          protocol: TCP
        - containerPort: 9000
          name: traefik
          protocol: TCP
        - containerPort: 8000
          name: web
          protocol: TCP
        - containerPort: 8443
          name: websecure
          protocol: TCP
        readinessProbe:
          failureThreshold: 1
          httpGet:
            path: /ping
            port: 9000
          initialDelaySeconds: 10
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 2
        resources:
          limits:
            cpu: "0.5"
            memory: 256Mi
          requests:
            cpu: "0.25"
            memory: 128Mi
        securityContext:
          capabilities:
            drop:
            - ALL
          readOnlyRootFilesystem: true
          runAsGroup: 65532
          runAsNonRoot: true
          runAsUser: 65532
        volumeMounts:
        - mountPath: /data
          name: data
        - mountPath: /tmp
          name: tmp
        - mountPath: /plugins-storage
          name: plugins
      hostNetwork: false
      securityContext:
        fsGroup: 65532
      serviceAccountName: traefik
      terminationGracePeriodSeconds: 60
      volumes:
      - emptyDir: {}
        name: data
      - emptyDir: {}
        name: tmp
      - emptyDir: {}
        name: plugins
```

## Nomad

Similar to Kubernetes but without a lot of the overhead. Developed by HashiCorp which also created Terraform/Consul/Vault.
Uses the HashiCorp Configuration Language (HCL) rather than a standard data format This isn't necessarily a bad thing but it does mean you have to learn an extra format on top of learning Nomad.

```hcl
io_mode = "async"

service "http" "web_proxy" {
  listen_addr = "127.0.0.1:8080"
  
  process "main" {
    command = ["/usr/local/bin/awesome-app", "server"]
  }

  process "mgmt" {
    command = ["/usr/local/bin/awesome-app", "mgmt"]
  }
}
```

## Bare Instance

Instead of using a supervisor you can just deploy to the instance instead using something like Ansible.
This keeps things simple but means you have to do manual automation to handle downtime and updating.

## "Serverless"

Serverless is a term for deploying code without having a predetermined server already configured.
