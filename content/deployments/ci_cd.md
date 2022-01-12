+++
title = "CI/CD"
weight = 1
+++

## Continuous Integration

Continuous integration allows you to have every commit to your repository go through a set of tests to determine any breaking changes. Generally speaking this would include building and testing the application, building any containers and potentially pushing them to a registry.

## Continuous Deployment

Continuous deployment allows valid commits that have gone through the integration tests to be deployed to an environment. Generally this starts out by deploying to a dev environment for testing and moves to full production releases once stabilized.

## Cloud Provider Platforms

I wouldn't generally recommend using a cloud providers CI/CD platforms, the reasons are two fold.

1. They aren't focused on it, there are features missing or just half baked.
2. A dedicated platform provides more features for less money and hassle.

An AWS CodeBuild project for instance only builds a single branch, a separate project would need to be created for each branch.

## GitHub Actions (General)

Easy to use tool that is built right into GitHub.
Made even easier with the community that has already formed around it, you can easily find an already made set of steps for most technologies. 
You can also easily add additional functionality via a custom "action" that can be shared with others as well.

## CDS (General)

Enterprise grade software developed by OVH, provides more options than other tools such as Jenkins or Bamboo but comes with additional complexity getting it deployed.

- Templated pipelines: Configure a pipeline for building an application and parameterize it so you can build multiple projects with it.
- Metrics and auditing: Every build comes with metrics so you can see how your pipelines perform over time.

## Argo (Kubernetes)

A set of applications and Kubernetes Custom Resource Definitions (CRDs) including:

1. [Argo Workflows](https://argoproj.github.io/argo-workflows/): The CI portion of the project, workflows can be configured to just about anything based on an event.
2. [ArgoCD](https://argo-cd.readthedocs.io/en/stable/): This deploys a service to a cluster after updates have been made.
3. [Argo Events](https://argoproj.github.io/argo-events/): An event can be a push to a git repository, a webhook to trigger a workflow and a great deal more.
4. [Argo Rollouts](https://argoproj.github.io/argo-rollouts/): Enables several new deployment strategies to Kubernetes including blue/green, canary and weighted strategies.

If you are already interested in Kubernetes than Argo is really helpful in adding missing functionality to your clusters. These all work in tandem to provide all necessary parts for an excellent CI/CD system.
