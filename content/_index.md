+++
title = "Fakku Cloud Migration"
insert_anchor_links = "right"
+++

Fakku is migrating services from the current "legacy" datacenter locations to a cloud provider.
These documents will outline choices made, research and general information on practices that are considered and implemented.

## Existing infrastructure

- Main application: PHP app, running on a 4 CPU machine with 32Gb of RAM. Usually hovers around 25-28 used memory.
- Database: MySQL DB (code changes required for new version), similar machine setup to Main App
- Caching: Memcache server, similar machine to others.
