---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_internetmonitor_monitor Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::InternetMonitor::Monitor
---

# awscc_internetmonitor_monitor (Data Source)

Data Source schema for AWS::InternetMonitor::Monitor



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `created_at` (String) The date value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ssZ)
- `modified_at` (String) The date value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ssZ)
- `monitor_arn` (String)
- `monitor_name` (String)
- `processing_status` (String)
- `processing_status_info` (String)
- `resources` (List of String)
- `resources_to_add` (List of String)
- `resources_to_remove` (List of String)
- `status` (String)
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String)
- `value` (String)

