---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_xray_transaction_search_config Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  This schema provides construct and validation rules for AWS-XRay TransactionSearchConfig resource parameters.
---

# awscc_xray_transaction_search_config (Resource)

This schema provides construct and validation rules for AWS-XRay TransactionSearchConfig resource parameters.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `indexing_percentage` (Number) Determines the percentage of traces indexed from CloudWatch Logs to X-Ray

### Read-Only

- `account_id` (String) User account id, used as the primary identifier for the resource
- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_xray_transaction_search_config.example "account_id"
```
