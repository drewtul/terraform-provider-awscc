---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_batch_scheduling_policy Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type schema for AWS::Batch::SchedulingPolicy
---

# awscc_batch_scheduling_policy (Resource)

Resource Type schema for AWS::Batch::SchedulingPolicy



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **fairshare_policy** (Attributes) Fair Share Policy for the Job Queue. (see [below for nested schema](#nestedatt--fairshare_policy))
- **name** (String) Name of Scheduling Policy.
- **tags** (Map of String) A key-value pair to associate with a resource.

### Read-Only

- **arn** (String) ARN of the Scheduling Policy.
- **id** (String) Uniquely identifies the resource.

<a id="nestedatt--fairshare_policy"></a>
### Nested Schema for `fairshare_policy`

Optional:

- **compute_reservation** (Number)
- **share_decay_seconds** (Number)
- **share_distribution** (Attributes List) List of Share Attributes (see [below for nested schema](#nestedatt--fairshare_policy--share_distribution))

<a id="nestedatt--fairshare_policy--share_distribution"></a>
### Nested Schema for `fairshare_policy.share_distribution`

Optional:

- **share_identifier** (String)
- **weight_factor** (Number)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_batch_scheduling_policy.example <resource ID>
```