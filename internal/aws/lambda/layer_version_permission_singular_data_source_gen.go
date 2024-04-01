// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_lambda_layer_version_permission", layerVersionPermissionDataSource)
}

// layerVersionPermissionDataSource returns the Terraform awscc_lambda_layer_version_permission data source.
// This Terraform data source corresponds to the CloudFormation AWS::Lambda::LayerVersionPermission resource.
func layerVersionPermissionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Action
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The API action that grants access to the layer.",
		//	  "type": "string"
		//	}
		"action": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The API action that grants access to the layer.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ID generated by service",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ID generated by service",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LayerVersionArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name or Amazon Resource Name (ARN) of the layer.",
		//	  "type": "string"
		//	}
		"layer_version_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name or Amazon Resource Name (ARN) of the layer.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OrganizationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "With the principal set to *, grant permission to all accounts in the specified organization.",
		//	  "type": "string"
		//	}
		"organization_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "With the principal set to *, grant permission to all accounts in the specified organization.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Principal
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An account ID, or * to grant layer usage permission to all accounts in an organization, or all AWS accounts (if organizationId is not specified).",
		//	  "type": "string"
		//	}
		"principal": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An account ID, or * to grant layer usage permission to all accounts in an organization, or all AWS accounts (if organizationId is not specified).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Lambda::LayerVersionPermission",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lambda::LayerVersionPermission").WithTerraformTypeName("awscc_lambda_layer_version_permission")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                      "Action",
		"layer_version_arn":           "LayerVersionArn",
		"layer_version_permission_id": "Id",
		"organization_id":             "OrganizationId",
		"principal":                   "Principal",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
