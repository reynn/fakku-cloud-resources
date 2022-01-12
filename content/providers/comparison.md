+++
title = "Comparison"
weight = 1
+++
## Glossary

| Term             | Definition                                                                                                                                                                                                                                       |
|------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Cloud Services   | Rented space on another companies servers. Offers easier ways to scale your applications due to changing workloads.                                                                                                                              |
| Serverless       | Run applications on providers compute without worrying about the underlying infrastructure.                                                                                                                                                      |
| Immutable Deploy | A deployment that can be recreated easily from a specific state. You would not make changes to the instance except to troubleshoot. Once troubleshooting is complete the changes needed are applied to the template and redeployed to ensure the |
|                  |                                                                                                                                                                                                                                                  |


| English            | DigitalOcean        | AWS                                    | Google                           |
|--------------------|---------------------|----------------------------------------|----------------------------------|
| Block Storage      | Spaces              | Simple Storage Service (`S3`)          | Cloud Storage                    |
| Managed Database   | Database            | Relational Database Service (`RDS`)    | Cloud SQL                        |
| Managed Kubernetes | Kubernetes (`DOKS`) | Elastic Kubernetes Service (`EKS`)     | Google Kubernetes Engine (`GKE`) |
| Networking         | Networking          | Virtual Private Cloud (`VPC`)          | Virtual Private Cloud            |
| Virtual Machine    | Droplet             | Elastic Cloud Compute (`EC2`) Instance | VM Instance                      |

## Common Between them

- The basic service of providing a Virtual Machine at competitive pricing is somewhat similar across the board. (within 5-10% usually)
- Offer a management interface via a dedicated website, provides some kind of user account control system to lock things down.
- A way to segment networks to minimize access from the outside world. Including firewalls, VPNs and isolated networks.
- The fastest way to move to the cloud is using a basic VM instance. AWS provides EC2 Amazon Machine Images (AMI) to save a state of an application to redeploy later.
    AWS EC2 Image Builder can automate the process from code commit to ready to use AMI that can be deployed to production or test servers immediately.
- Provide options to backup and restore databases and VMs.

## AWS

### AWS Pros

- Infrastructure as Code via [Cloudformation](https://aws.amazon.com/cloudformation/) or [AWS CDK](https://aws.amazon.com/cdk/).
- ~70% market share, getting devs who understand AWS is often much easier than other cloud providers.
- Extensive documentation and tutorials online to explain the usefulness of a service.
- Extensive collection of services. There is a good chance you can find a managed version of an typical applications you might use in distributed services.

### AWS Cons

- [Cloudformation](https://aws.amazon.com/cloudformation/) can become a nightmare, documentation is great but if not experienced with YAML the indentation can be very difficult. `CDK` is an  improvement, but only mostly because you use a traditional programming language such as TypeScript, Python, Java, .NET, and Go.
- Overcomplicated, from naming to billing and configuration setting up AWS accounts securely takes dedicated resources
- Expensive, unless you sign up for reserved instances everything can get very expensive and it is hard to locate all of the services you are using.
- Excessive "managed" services, provides over 100 services many of which are managed services that can be configured manually. This helps ease management of services but also makes it difficult to tell if you really need the service.

## Google

### Google Pros

- Best managed Kubernetes offering available, [GKE Autopilot](https://cloud.google.com/kubernetes-engine/docs/concepts/autopilot-overview).
- Fastest availability of new Kubernetes versions, already offers 1.22 (via Rapid Release Channel) where AWS is still 1.21 max.
- Google offers "Preemptible Nodes", these only last 24 hrs maximum, but they are significantly cheaper than standard nodes. In K8s this isn't a problem as your workloads auto-transition from node to node.

### Google Cons

- Google actually is the newest player out of all on this list, even if only by a few months. Their history with other Google services make me quite cautious of them.
- CLI is the most complex to use than any the others. Permissions get very complicated as well and in some ways is worse than AWS.

## DigitalOcean

### DigitalOcean Pros

- Significantly smaller footprint in terms of offered services. Focus is on providing a solid platform.
- Simpler terms and configuration make it far easier for anyone familiar with general computing to get on board.
- Company also values [open source](https://docs.digitalocean.com/reference/opensource/), many of their projects are open sourced and written in modern languages contributing to their success.
- Provides a CLI like most others, but it is significantly easier to use than Google or Amazon.

### DigitalOcean Cons

- Infrastructure as Code only through a tool like [Ansible](https://www.ansible.com/), [Terraform](https://www.terraform.io/) or [Chef](https://www.chef.io/).
- More limited managed services (for some) some services (queue, [ElasticSearch](https://www.elastic.co/elasticsearch/) etc) have to be self-managed, increased workload.

## Decision

**DigitalOcean**

I would chose `DO` over `AWS` for the following reasons:

- `AWS` requires you to become an expert not only in the software you are using but also how to manage and secure everything within the `AWS` ecosystem.
- `DigitalOcean` is a far simpler solution, this has a big advantage because you don't need to pay up for people who understand `AWS` someone with more basic IT familiarity can get up to speed on it faster.
- `DigitalOcean` provides a great framework and has a provides a far easier to use and more manageable `Serverless` solution compared to others.
- `AWS` obscures pricing by giving it to you in extremely small denominations to make it seem cheap. These wind up adding up fast and can become costly very fast if the proper safe guards aren't in place.



## Notes

**Azure was considered however I decided to omit it. No significant advantages could be found over the ones provided above.**

**Linode may be a contender in the future but unfortunately right now it is too limited in functionality to recommend**

