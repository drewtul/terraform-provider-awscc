---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_medialive_cloudwatch_alarm_template_group Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::MediaLive::CloudWatchAlarmTemplateGroup
---

# awscc_medialive_cloudwatch_alarm_template_group (Data Source)

Data Source schema for AWS::MediaLive::CloudWatchAlarmTemplateGroup



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `arn` (String) A cloudwatch alarm template group's ARN (Amazon Resource Name)
- `cloudwatch_alarm_template_group_id` (String) A cloudwatch alarm template group's id. AWS provided template groups have ids that start with `aws-`
- `created_at` (String)
- `description` (String) A resource's optional description.
- `identifier` (String)
- `modified_at` (String)
- `name` (String) A resource's name. Names must be unique within the scope of a resource type in a specific region.
- `tags` (Map of String) Represents the tags associated with a resource.