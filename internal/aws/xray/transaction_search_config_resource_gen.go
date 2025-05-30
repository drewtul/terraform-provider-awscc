// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package xray

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_xray_transaction_search_config", transactionSearchConfigResource)
}

// transactionSearchConfigResource returns the Terraform awscc_xray_transaction_search_config resource.
// This Terraform resource corresponds to the CloudFormation AWS::XRay::TransactionSearchConfig resource.
func transactionSearchConfigResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "User account id, used as the primary identifier for the resource",
		//	  "pattern": "^\\d{12}$",
		//	  "type": "string"
		//	}
		"account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "User account id, used as the primary identifier for the resource",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IndexingPercentage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Determines the percentage of traces indexed from CloudWatch Logs to X-Ray",
		//	  "maximum": 100,
		//	  "minimum": 0,
		//	  "type": "number"
		//	}
		"indexing_percentage": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "Determines the percentage of traces indexed from CloudWatch Logs to X-Ray",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Float64{ /*START VALIDATORS*/
				float64validator.Between(0.000000, 100.000000),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
				float64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "This schema provides construct and validation rules for AWS-XRay TransactionSearchConfig resource parameters.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::XRay::TransactionSearchConfig").WithTerraformTypeName("awscc_xray_transaction_search_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_id":          "AccountId",
		"indexing_percentage": "IndexingPercentage",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
