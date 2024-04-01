// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package macie

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_macie_custom_data_identifier", customDataIdentifierDataSource)
}

// customDataIdentifierDataSource returns the Terraform awscc_macie_custom_data_identifier data source.
// This Terraform data source corresponds to the CloudFormation AWS::Macie::CustomDataIdentifier resource.
func customDataIdentifierDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Custom data identifier ARN.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Custom data identifier ARN.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Description of custom data identifier.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Description of custom data identifier.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Custom data identifier ID.",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Custom data identifier ID.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IgnoreWords
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Words to be ignored.",
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"ignore_words": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "Words to be ignored.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Keywords
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Keywords to be matched against.",
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"keywords": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "Keywords to be matched against.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MaximumMatchDistance
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Maximum match distance.",
		//	  "type": "integer"
		//	}
		"maximum_match_distance": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Maximum match distance.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of custom data identifier.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of custom data identifier.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Regex
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Regular expression for custom data identifier.",
		//	  "type": "string"
		//	}
		"regex": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Regular expression for custom data identifier.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A collection of tags associated with a resource",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The tag's key.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The tag's value.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag's key.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag's value.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A collection of tags associated with a resource",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Macie::CustomDataIdentifier",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Macie::CustomDataIdentifier").WithTerraformTypeName("awscc_macie_custom_data_identifier")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                       "Arn",
		"custom_data_identifier_id": "Id",
		"description":               "Description",
		"ignore_words":              "IgnoreWords",
		"key":                       "Key",
		"keywords":                  "Keywords",
		"maximum_match_distance":    "MaximumMatchDistance",
		"name":                      "Name",
		"regex":                     "Regex",
		"tags":                      "Tags",
		"value":                     "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
