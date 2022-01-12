+++
title = "Amazon Web Services (AWS)"
weight = 2
+++

Founded 2004 with with just the SQS service, they were first to market with cloud offerings and as such have become
the most dominate.

## Available DC Locations

Data gathered on 12/07/2021 from [AWS Regions](https://aws.amazon.com/about-aws/global-infrastructure/regions_az/?p=ngi&loc=2)

25 regions available, with 9 more coming in the next 2-3 years. Not all services available in all regions. This is a list of the public cloud regions. China and a GovCloud are also available with multiple regions each.

- Australia: Sydney
- Bahrain: Unknown
- Brazil: São Paulo
- China: Hong Kong
- England: London
- France: Paris
- Germany: Frankfurt
- India: Mumbai
- Ireland: Unknown
- Italy: Milan
- Japan: Osaka
- Japan: Tokyo
- Singapore: Unknown
- South Korea: Seoul
- Sweden: Stockholm
- USA: North California
- USA: North Virginia
- USA: Ohio
- USA: Oregon

## Pricing

### Instance Pricing

Pricing from the [AWS On Demand pricing](https://aws.amazon.com/ec2/pricing/on-demand/)

| Type         | CPUs | RAM  | Net Speed   | Cost    |
| ------------ | ---- | ---- | ----------- | ------- |
| m6i.2xlarge  | 8    | 32GB | Up to 12 Gb | 280.32  |
| m6g.2xlarge  | 8    | 32GB | Up to 10 Gb | 224.84  |
| c6i.2xlarge  | 16   | 32GB | Up to 12 Gb | 496.40  |
| c6g.2xlarge  | 16   | 32GB | Up to 10 Gb | 397.12  |
| t3.2xlarge   | 8    | 32GB | Up to 5Gb   | 242.94  |
| t3a.2xlarge  | 8    | 32GB | Up to 5Gb   | 219.58  |
| t4g.2xlarge  | 8    | 32GB | Up to 5Gb   | 196.22  |
| c6gd.4xlarge | 16   | 32GB | Up to 10Gb  | 448.51  | This has NVME storage rather than an typical `gp2` EBS volume. |
| m6gd.2xlarge | 8    | 32GB | Up to 10Gb  | 263.968 | This has NVME storage rather than an typical `gp2` EBS volume. |

- No extra letter before the `.` means it is older instance type, 99% chance it will be an Intel CPU
- Ending in `a` means it is an AMD x86_64 CPU
- Ending in `i` means it is an Intel x86_64 CPU
- Ending in `g` means it is an AWS Graviton ARM64 CPU
