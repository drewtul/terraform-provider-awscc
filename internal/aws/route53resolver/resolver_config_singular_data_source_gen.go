// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package route53resolver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_route53resolver_resolver_config", resolverConfigDataSource)
}

// resolverConfigDataSource returns the Terraform awscc_route53resolver_resolver_config data source.
// This Terraform data source corresponds to the CloudFormation AWS::Route53Resolver::ResolverConfig resource.
func resolverConfigDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AutodefinedReverse
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ResolverAutodefinedReverseStatus, possible values are ENABLING, ENABLED, DISABLING AND DISABLED.",
		//	  "enum": [
		//	    "ENABLING",
		//	    "ENABLED",
		//	    "DISABLING",
		//	    "DISABLED"
		//	  ],
		//	  "type": "string"
		//	}
		"autodefined_reverse": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ResolverAutodefinedReverseStatus, possible values are ENABLING, ENABLED, DISABLING AND DISABLED.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AutodefinedReverseFlag
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Represents the desired status of AutodefinedReverse. The only supported value on creation is DISABLE. Deletion of this resource will return AutodefinedReverse to its default value (ENABLED).",
		//	  "enum": [
		//	    "DISABLE"
		//	  ],
		//	  "type": "string"
		//	}
		"autodefined_reverse_flag": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Represents the desired status of AutodefinedReverse. The only supported value on creation is DISABLE. Deletion of this resource will return AutodefinedReverse to its default value (ENABLED).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Id",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Id",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OwnerId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "AccountId",
		//	  "maxLength": 32,
		//	  "minLength": 12,
		//	  "type": "string"
		//	}
		"owner_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "AccountId",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ResourceId",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"resource_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ResourceId",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Route53Resolver::ResolverConfig",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53Resolver::ResolverConfig").WithTerraformTypeName("awscc_route53resolver_resolver_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"autodefined_reverse":      "AutodefinedReverse",
		"autodefined_reverse_flag": "AutodefinedReverseFlag",
		"owner_id":                 "OwnerId",
		"resolver_config_id":       "Id",
		"resource_id":              "ResourceId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
