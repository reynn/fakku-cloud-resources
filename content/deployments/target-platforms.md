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
