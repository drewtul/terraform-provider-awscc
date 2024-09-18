// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package entityresolution

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_entityresolution_schema_mapping", schemaMappingResource)
}

// schemaMappingResource returns the Terraform awscc_entityresolution_schema_mapping resource.
// This Terraform resource corresponds to the CloudFormation AWS::EntityResolution::SchemaMapping resource.
func schemaMappingResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time of this SchemaMapping got created",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time of this SchemaMapping got created",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the SchemaMapping",
		//	  "maxLength": 255,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the SchemaMapping",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 255),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: HasWorkflows
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The boolean value that indicates whether or not a SchemaMapping has MatchingWorkflows that are associated with",
		//	  "type": "boolean"
		//	}
		"has_workflows": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "The boolean value that indicates whether or not a SchemaMapping has MatchingWorkflows that are associated with",
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MappedInputFields
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The SchemaMapping attributes input",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "FieldName": {
		//	        "maxLength": 255,
		//	        "minLength": 0,
		//	        "pattern": "^[a-zA-Z_0-9- \\t]*$",
		//	        "type": "string"
		//	      },
		//	      "GroupName": {
		//	        "maxLength": 255,
		//	        "minLength": 0,
		//	        "pattern": "^[a-zA-Z_0-9- \\t]*$",
		//	        "type": "string"
		//	      },
		//	      "Hashed": {
		//	        "type": "boolean"
		//	      },
		//	      "MatchKey": {
		//	        "maxLength": 255,
		//	        "minLength": 0,
		//	        "pattern": "^[a-zA-Z_0-9- \\t]*$",
		//	        "type": "string"
		//	      },
		//	      "SubType": {
		//	        "description": "The subtype of the Attribute. Would be required only when type is PROVIDER_ID",
		//	        "type": "string"
		//	      },
		//	      "Type": {
		//	        "enum": [
		//	          "NAME",
		//	          "NAME_FIRST",
		//	          "NAME_MIDDLE",
		//	          "NAME_LAST",
		//	          "ADDRESS",
		//	          "ADDRESS_STREET1",
		//	          "ADDRESS_STREET2",
		//	          "ADDRESS_STREET3",
		//	          "ADDRESS_CITY",
		//	          "ADDRESS_STATE",
		//	          "ADDRESS_COUNTRY",
		//	          "ADDRESS_POSTALCODE",
		//	          "PHONE",
		//	          "PHONE_NUMBER",
		//	          "PHONE_COUNTRYCODE",
		//	          "EMAIL_ADDRESS",
		//	          "UNIQUE_ID",
		//	          "DATE",
		//	          "STRING",
		//	          "PROVIDER_ID"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "FieldName",
		//	      "Type"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 35,
		//	  "minItems": 2,
		//	  "type": "array"
		//	}
		"mapped_input_fields": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: FieldName
					"field_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 255),
							stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z_0-9- \\t]*$"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: GroupName
					"group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 255),
							stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z_0-9- \\t]*$"), ""),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Hashed
					"hashed": schema.BoolAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
							boolplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: MatchKey
					"match_key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 255),
							stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z_0-9- \\t]*$"), ""),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: SubType
					"sub_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The subtype of the Attribute. Would be required only when type is PROVIDER_ID",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Type
					"type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"NAME",
								"NAME_FIRST",
								"NAME_MIDDLE",
								"NAME_LAST",
								"ADDRESS",
								"ADDRESS_STREET1",
								"ADDRESS_STREET2",
								"ADDRESS_STREET3",
								"ADDRESS_CITY",
								"ADDRESS_STATE",
								"ADDRESS_COUNTRY",
								"ADDRESS_POSTALCODE",
								"PHONE",
								"PHONE_NUMBER",
								"PHONE_COUNTRYCODE",
								"EMAIL_ADDRESS",
								"UNIQUE_ID",
								"DATE",
								"STRING",
								"PROVIDER_ID",
							),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The SchemaMapping attributes input",
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(2, 35),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SchemaArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The SchemaMapping arn associated with the Schema",
		//	  "pattern": "^arn:(aws|aws-us-gov|aws-cn):entityresolution:.*:[0-9]+:(schemamapping/.*)$",
		//	  "type": "string"
		//	}
		"schema_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The SchemaMapping arn associated with the Schema",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SchemaName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the SchemaMapping",
		//	  "maxLength": 255,
		//	  "minLength": 0,
		//	  "pattern": "^[a-zA-Z_0-9-]*$",
		//	  "type": "string"
		//	}
		"schema_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the SchemaMapping",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 255),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z_0-9-]*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 200,
		//	  "minItems": 0,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeBetween(0, 200),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time of this SchemaMapping got last updated at",
		//	  "type": "string"
		//	}
		"updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time of this SchemaMapping got last updated at",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
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
		Description: "SchemaMapping defined in AWS Entity Resolution service",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EntityResolution::SchemaMapping").WithTerraformTypeName("awscc_entityresolution_schema_mapping")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"created_at":          "CreatedAt",
		"description":         "Description",
		"field_name":          "FieldName",
		"group_name":          "GroupName",
		"has_workflows":       "HasWorkflows",
		"hashed":              "Hashed",
		"key":                 "Key",
		"mapped_input_fields": "MappedInputFields",
		"match_key":           "MatchKey",
		"schema_arn":          "SchemaArn",
		"schema_name":         "SchemaName",
		"sub_type":            "SubType",
		"tags":                "Tags",
		"type":                "Type",
		"updated_at":          "UpdatedAt",
		"value":               "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
