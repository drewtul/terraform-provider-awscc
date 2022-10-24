---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_greengrassv2_deployment Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::GreengrassV2::Deployment
---

# awscc_greengrassv2_deployment (Data Source)

Data Source schema for AWS::GreengrassV2::Deployment



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `components` (Attributes Map) (see [below for nested schema](#nestedatt--components))
- `deployment_id` (String)
- `deployment_name` (String)
- `deployment_policies` (Attributes) (see [below for nested schema](#nestedatt--deployment_policies))
- `iot_job_configuration` (Attributes) (see [below for nested schema](#nestedatt--iot_job_configuration))
- `tags` (Map of String)
- `target_arn` (String)

<a id="nestedatt--components"></a>
### Nested Schema for `components`

Read-Only:

- `component_version` (String)
- `configuration_update` (Attributes) (see [below for nested schema](#nestedatt--components--configuration_update))
- `run_with` (Attributes) (see [below for nested schema](#nestedatt--components--run_with))

<a id="nestedatt--components--configuration_update"></a>
### Nested Schema for `components.configuration_update`

Read-Only:

- `merge` (String)
- `reset` (List of String)


<a id="nestedatt--components--run_with"></a>
### Nested Schema for `components.run_with`

Read-Only:

- `posix_user` (String)
- `system_resource_limits` (Attributes) (see [below for nested schema](#nestedatt--components--run_with--system_resource_limits))
- `windows_user` (String)

<a id="nestedatt--components--run_with--system_resource_limits"></a>
### Nested Schema for `components.run_with.system_resource_limits`

Read-Only:

- `cpus` (Number)
- `memory` (Number)




<a id="nestedatt--deployment_policies"></a>
### Nested Schema for `deployment_policies`

Read-Only:

- `component_update_policy` (Attributes) (see [below for nested schema](#nestedatt--deployment_policies--component_update_policy))
- `configuration_validation_policy` (Attributes) (see [below for nested schema](#nestedatt--deployment_policies--configuration_validation_policy))
- `failure_handling_policy` (String)

<a id="nestedatt--deployment_policies--component_update_policy"></a>
### Nested Schema for `deployment_policies.component_update_policy`

Read-Only:

- `action` (String)
- `timeout_in_seconds` (Number)


<a id="nestedatt--deployment_policies--configuration_validation_policy"></a>
### Nested Schema for `deployment_policies.configuration_validation_policy`

Read-Only:

- `timeout_in_seconds` (Number)



<a id="nestedatt--iot_job_configuration"></a>
### Nested Schema for `iot_job_configuration`

Read-Only:

- `abort_config` (Attributes) (see [below for nested schema](#nestedatt--iot_job_configuration--abort_config))
- `job_executions_rollout_config` (Attributes) (see [below for nested schema](#nestedatt--iot_job_configuration--job_executions_rollout_config))
- `timeout_config` (Attributes) (see [below for nested schema](#nestedatt--iot_job_configuration--timeout_config))

<a id="nestedatt--iot_job_configuration--abort_config"></a>
### Nested Schema for `iot_job_configuration.abort_config`

Read-Only:

- `criteria_list` (Attributes List) (see [below for nested schema](#nestedatt--iot_job_configuration--abort_config--criteria_list))

<a id="nestedatt--iot_job_configuration--abort_config--criteria_list"></a>
### Nested Schema for `iot_job_configuration.abort_config.criteria_list`

Read-Only:

- `action` (String)
- `failure_type` (String)
- `min_number_of_executed_things` (Number)
- `threshold_percentage` (Number)



<a id="nestedatt--iot_job_configuration--job_executions_rollout_config"></a>
### Nested Schema for `iot_job_configuration.job_executions_rollout_config`

Read-Only:

- `exponential_rate` (Attributes) (see [below for nested schema](#nestedatt--iot_job_configuration--job_executions_rollout_config--exponential_rate))
- `maximum_per_minute` (Number)

<a id="nestedatt--iot_job_configuration--job_executions_rollout_config--exponential_rate"></a>
### Nested Schema for `iot_job_configuration.job_executions_rollout_config.exponential_rate`

Read-Only:

- `base_rate_per_minute` (Number)
- `increment_factor` (Number)
- `rate_increase_criteria` (Attributes) (see [below for nested schema](#nestedatt--iot_job_configuration--job_executions_rollout_config--exponential_rate--rate_increase_criteria))

<a id="nestedatt--iot_job_configuration--job_executions_rollout_config--exponential_rate--rate_increase_criteria"></a>
### Nested Schema for `iot_job_configuration.job_executions_rollout_config.exponential_rate.rate_increase_criteria`

Read-Only:

- `number_of_notified_things` (Number)
- `number_of_succeeded_things` (Number)




<a id="nestedatt--iot_job_configuration--timeout_config"></a>
### Nested Schema for `iot_job_configuration.timeout_config`

Read-Only:

- `in_progress_timeout_in_minutes` (Number)

