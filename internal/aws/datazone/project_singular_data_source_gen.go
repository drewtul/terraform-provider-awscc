// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package datazone

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_datazone_project", projectDataSource)
}

// projectDataSource returns the Terraform awscc_datazone_project data source.
// This Terraform data source corresponds to the CloudFormation AWS::DataZone::Project resource.
func projectDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp of when the project was created.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "The timestamp of when the project was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedBy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon DataZone user who created the project.",
		//	  "type": "string"
		//	}
		"created_by": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon DataZone user who created the project.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the Amazon DataZone project.",
		//	  "maxLength": 2048,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the Amazon DataZone project.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DomainId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the Amazon DataZone domain in which the project was created.",
		//	  "pattern": "^dzd[-_][a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"domain_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the Amazon DataZone domain in which the project was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DomainIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the Amazon DataZone domain in which this project is created.",
		//	  "pattern": "^dzd[-_][a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"domain_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the Amazon DataZone domain in which this project is created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GlossaryTerms
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The glossary terms that can be used in this Amazon DataZone project.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "pattern": "^[a-zA-Z0-9_-]{1,36}$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 20,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"glossary_terms": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The glossary terms that can be used in this Amazon DataZone project.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the Amazon DataZone project.",
		//	  "pattern": "^[a-zA-Z0-9_-]{1,36}$",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the Amazon DataZone project.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LastUpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp of when the project was last updated.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"last_updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "The timestamp of when the project was last updated.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the Amazon DataZone project.",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "pattern": "^[\\w -]+$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the Amazon DataZone project.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::DataZone::Project",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataZone::Project").WithTerraformTypeName("awscc_datazone_project")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"created_at":        "CreatedAt",
		"created_by":        "CreatedBy",
		"description":       "Description",
		"domain_id":         "DomainId",
		"domain_identifier": "DomainIdentifier",
		"glossary_terms":    "GlossaryTerms",
		"last_updated_at":   "LastUpdatedAt",
		"name":              "Name",
		"project_id":        "Id",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
