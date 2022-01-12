+++
title = "Security"
weight = 2
+++

## Virtual Private Network (VPN)

1. DigitalOcean: provides quick setup for OpenVPN, UTunnel VPN and Wireguardian VPN.
1. AWS: Provides a managed OpenVPN installation.
1. Google: Provides an High Availability (HA) VPN using IPsec.

## Firewall

1. DigitalOcean: Simply called "Firewall" under Networking, this allows you to restrict access to instances via a set of rules. Can be applied automatically to instances based on tagging.
2. AWS: Referred to as "Security Groups" offers the same functionality as you would expect from a normal firewall.
3. Google: "Cloud Firewalls" can be setup at multiple levels so that organizations can create general rules, and individual projects can restrict further.

## Backup and Restore

1. DigitalOcean: Droplets and Databases can be backed up for a small aditional monthly cost ($10/instance).
2. AWS: Instances volumes can be backed up via "Snapshots". The snapshot must be created per attached volume rather than by the instance itself.
3. Google: Similar to AWS you can "Snapshot" individual volumes that can be restored later.
