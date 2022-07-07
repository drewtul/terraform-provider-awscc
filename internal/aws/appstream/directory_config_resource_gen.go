// Code generated by generators/resource/main.go; DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_appstream_directory_config", directoryConfigResourceType)
}

// directoryConfigResourceType returns the Terraform awscc_appstream_directory_config resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::AppStream::DirectoryConfig resource type.
func directoryConfigResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"directory_name": {
			// Property: DirectoryName
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Required: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"organizational_unit_distinguished_names": {
			// Property: OrganizationalUnitDistinguishedNames
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Type:     types.ListType{ElemType: types.StringType},
			Required: true,
		},
		"service_account_credentials": {
			// Property: ServiceAccountCredentials
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "AccountName": {
			//       "type": "string"
			//     },
			//     "AccountPassword": {
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "AccountName",
			//     "AccountPassword"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"account_name": {
						// Property: AccountName
						Type:     types.StringType,
						Required: true,
					},
					"account_password": {
						// Property: AccountPassword
						Type:     types.StringType,
						Required: true,
						// AccountPassword is a write-only property.
					},
				},
			),
			Required: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			tfsdk.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::AppStream::DirectoryConfig",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::AppStream::DirectoryConfig").WithTerraformTypeName("awscc_appstream_directory_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_name":     "AccountName",
		"account_password": "AccountPassword",
		"directory_name":   "DirectoryName",
		"organizational_unit_distinguished_names": "OrganizationalUnitDistinguishedNames",
		"service_account_credentials":             "ServiceAccountCredentials",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/ServiceAccountCredentials/AccountPassword",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}