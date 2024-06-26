---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_gamelift_container_group_definition Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::GameLift::ContainerGroupDefinition
---

# awscc_gamelift_container_group_definition (Data Source)

Data Source schema for AWS::GameLift::ContainerGroupDefinition



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `container_definitions` (Attributes Set) A collection of container definitions that define the containers in this group. (see [below for nested schema](#nestedatt--container_definitions))
- `container_group_definition_arn` (String) The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift container group resource and uniquely identifies it across all AWS Regions.
- `creation_time` (String) A time stamp indicating when this data object was created. Format is a number expressed in Unix time as milliseconds (for example "1469498468.057").
- `name` (String) A descriptive label for the container group definition.
- `operating_system` (String) The operating system of the container group
- `scheduling_strategy` (String) Specifies whether the container group includes replica or daemon containers.
- `tags` (Attributes Set) An array of key-value pairs to apply to this resource. (see [below for nested schema](#nestedatt--tags))
- `total_cpu_limit` (Number) The maximum number of CPU units reserved for this container group. The value is expressed as an integer amount of CPU units. (1 vCPU is equal to 1024 CPU units.)
- `total_memory_limit` (Number) The maximum amount of memory (in MiB) to allocate for this container group.

<a id="nestedatt--container_definitions"></a>
### Nested Schema for `container_definitions`

Read-Only:

- `command` (List of String) The command that's passed to the container.
- `container_name` (String) A descriptive label for the container definition. Container definition names must be unique with a container group definition.
- `cpu` (Number) The maximum number of CPU units reserved for this container. The value is expressed as an integer amount of CPU units. 1 vCPU is equal to 1024 CPU units
- `depends_on` (Attributes List) A list of container dependencies that determines when this container starts up and shuts down. For container groups with multiple containers, dependencies let you define a startup/shutdown sequence across the containers. (see [below for nested schema](#nestedatt--container_definitions--depends_on))
- `entry_point` (List of String) The entry point that's passed to the container so that it will run as an executable. If there are multiple arguments, each argument is a string in the array.
- `environment` (Attributes Set) The environment variables to pass to a container. (see [below for nested schema](#nestedatt--container_definitions--environment))
- `essential` (Boolean) Specifies if the container is essential. If an essential container fails a health check, then all containers in the container group will be restarted. You must specify exactly 1 essential container in a container group.
- `health_check` (Attributes) Specifies how the health of the containers will be checked. (see [below for nested schema](#nestedatt--container_definitions--health_check))
- `image_uri` (String) Specifies the image URI of this container.
- `memory_limits` (Attributes) Specifies how much memory is available to the container. You must specify at least this parameter or the TotalMemoryLimit parameter of the ContainerGroupDefinition. (see [below for nested schema](#nestedatt--container_definitions--memory_limits))
- `port_configuration` (Attributes) Defines the ports on the container. (see [below for nested schema](#nestedatt--container_definitions--port_configuration))
- `resolved_image_digest` (String) The digest of the container image.
- `working_directory` (String) The working directory to run commands inside the container in.

<a id="nestedatt--container_definitions--depends_on"></a>
### Nested Schema for `container_definitions.depends_on`

Read-Only:

- `condition` (String) The type of dependency.
- `container_name` (String) A descriptive label for the container definition. The container being defined depends on this container's condition.


<a id="nestedatt--container_definitions--environment"></a>
### Nested Schema for `container_definitions.environment`

Read-Only:

- `name` (String) The environment variable name.
- `value` (String) The environment variable value.


<a id="nestedatt--container_definitions--health_check"></a>
### Nested Schema for `container_definitions.health_check`

Read-Only:

- `command` (List of String) A string array representing the command that the container runs to determine if it is healthy.
- `interval` (Number) How often (in seconds) the health is checked.
- `retries` (Number) How many times the process manager will retry the command after a timeout. (The first run of the command does not count as a retry.)
- `start_period` (Number) The optional grace period (in seconds) to give a container time to boostrap before teh health check is declared failed.
- `timeout` (Number) How many seconds the process manager allows the command to run before canceling it.


<a id="nestedatt--container_definitions--memory_limits"></a>
### Nested Schema for `container_definitions.memory_limits`

Read-Only:

- `hard_limit` (Number) The hard limit of memory to reserve for the container.
- `soft_limit` (Number) The amount of memory that is reserved for the container.


<a id="nestedatt--container_definitions--port_configuration"></a>
### Nested Schema for `container_definitions.port_configuration`

Read-Only:

- `container_port_ranges` (Attributes Set) Specifies one or more ranges of ports on a container. (see [below for nested schema](#nestedatt--container_definitions--port_configuration--container_port_ranges))

<a id="nestedatt--container_definitions--port_configuration--container_port_ranges"></a>
### Nested Schema for `container_definitions.port_configuration.container_port_ranges`

Read-Only:

- `from_port` (Number) A starting value for the range of allowed port numbers.
- `protocol` (String) Defines the protocol of these ports.
- `to_port` (Number) An ending value for the range of allowed port numbers. Port numbers are end-inclusive. This value must be equal to or greater than FromPort.




<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- `key` (String) The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length.
- `value` (String) The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length.
