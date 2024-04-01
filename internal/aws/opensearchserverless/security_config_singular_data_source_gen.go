// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package opensearchserverless

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_opensearchserverless_security_config", securityConfigDataSource)
}

// securityConfigDataSource returns the Terraform awscc_opensearchserverless_security_config data source.
// This Terraform data source corresponds to the CloudFormation AWS::OpenSearchServerless::SecurityConfig resource.
func securityConfigDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Security config description",
		//	  "maxLength": 1000,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Security config description",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the security config",
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the security config",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The friendly name of the security config",
		//	  "maxLength": 32,
		//	  "minLength": 3,
		//	  "pattern": "^[a-z][a-z0-9-]{2,31}$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The friendly name of the security config",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SamlOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Describes saml options in form of key value map",
		//	  "properties": {
		//	    "GroupAttribute": {
		//	      "description": "Group attribute for this saml integration",
		//	      "maxLength": 2048,
		//	      "minLength": 1,
		//	      "pattern": "[\\w+=,.@-]+",
		//	      "type": "string"
		//	    },
		//	    "Metadata": {
		//	      "description": "The XML saml provider metadata document that you want to use",
		//	      "maxLength": 51200,
		//	      "minLength": 1,
		//	      "pattern": "",
		//	      "type": "string"
		//	    },
		//	    "SessionTimeout": {
		//	      "description": "Defines the session timeout in minutes",
		//	      "type": "integer"
		//	    },
		//	    "UserAttribute": {
		//	      "description": "Custom attribute for this saml integration",
		//	      "maxLength": 2048,
		//	      "minLength": 1,
		//	      "pattern": "[\\w+=,.@-]+",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Metadata"
		//	  ],
		//	  "type": "object"
		//	}
		"saml_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: GroupAttribute
				"group_attribute": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Group attribute for this saml integration",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Metadata
				"metadata": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The XML saml provider metadata document that you want to use",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SessionTimeout
				"session_timeout": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "Defines the session timeout in minutes",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: UserAttribute
				"user_attribute": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Custom attribute for this saml integration",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Describes saml options in form of key value map",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Config type for security config",
		//	  "enum": [
		//	    "saml"
		//	  ],
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Config type for security config",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::OpenSearchServerless::SecurityConfig",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::OpenSearchServerless::SecurityConfig").WithTerraformTypeName("awscc_opensearchserverless_security_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":        "Description",
		"group_attribute":    "GroupAttribute",
		"metadata":           "Metadata",
		"name":               "Name",
		"saml_options":       "SamlOptions",
		"security_config_id": "Id",
		"session_timeout":    "SessionTimeout",
		"type":               "Type",
		"user_attribute":     "UserAttribute",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
