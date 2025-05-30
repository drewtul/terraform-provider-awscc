// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ivs

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_ivs_ingest_configuration", ingestConfigurationResource)
}

// ingestConfigurationResource returns the Terraform awscc_ivs_ingest_configuration resource.
// This Terraform resource corresponds to the CloudFormation AWS::IVS::IngestConfiguration resource.
func ingestConfigurationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "IngestConfiguration ARN is automatically generated on creation and assigned as the unique identifier.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "^arn:aws:ivs:[a-z0-9-]+:[0-9]+:ingest-configuration/[a-zA-Z0-9-]+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "IngestConfiguration ARN is automatically generated on creation and assigned as the unique identifier.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IngestProtocol
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "RTMPS",
		//	  "description": "Ingest Protocol.",
		//	  "enum": [
		//	    "RTMP",
		//	    "RTMPS"
		//	  ],
		//	  "type": "string"
		//	}
		"ingest_protocol": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Ingest Protocol.",
			Optional:    true,
			Computed:    true,
			Default:     stringdefault.StaticString("RTMPS"),
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"RTMP",
					"RTMPS",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InsecureIngest
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": false,
		//	  "description": "Whether ingest configuration allows insecure ingest.",
		//	  "type": "boolean"
		//	}
		"insecure_ingest": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Whether ingest configuration allows insecure ingest.",
			Optional:    true,
			Computed:    true,
			Default:     booldefault.StaticBool(false),
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
				boolplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// InsecureIngest is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "-",
		//	  "description": "IngestConfiguration",
		//	  "maxLength": 128,
		//	  "minLength": 0,
		//	  "pattern": "^[a-zA-Z0-9-_]*$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "IngestConfiguration",
			Optional:    true,
			Computed:    true,
			Default:     stringdefault.StaticString("-"),
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 128),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9-_]*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ParticipantId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Participant Id is automatically generated on creation and assigned.",
		//	  "maxLength": 64,
		//	  "minLength": 0,
		//	  "pattern": "^[a-zA-Z0-9-_]*$",
		//	  "type": "string"
		//	}
		"participant_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Participant Id is automatically generated on creation and assigned.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StageArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "",
		//	  "description": "Stage ARN. A value other than an empty string indicates that stage is linked to IngestConfiguration. Default: \"\" (recording is disabled).",
		//	  "maxLength": 128,
		//	  "minLength": 0,
		//	  "pattern": "^arn:aws:ivs:[a-z0-9-]+:[0-9]+:stage/[a-zA-Z0-9-]+$",
		//	  "type": "string"
		//	}
		"stage_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Stage ARN. A value other than an empty string indicates that stage is linked to IngestConfiguration. Default: \"\" (recording is disabled).",
			Optional:    true,
			Computed:    true,
			Default:     stringdefault.StaticString(""),
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 128),
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws:ivs:[a-z0-9-]+:[0-9]+:stage/[a-zA-Z0-9-]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "INACTIVE",
		//	  "description": "State of IngestConfiguration which determines whether IngestConfiguration is in use or not.",
		//	  "enum": [
		//	    "ACTIVE",
		//	    "INACTIVE"
		//	  ],
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "State of IngestConfiguration which determines whether IngestConfiguration is in use or not.",
			Computed:    true,
			Default:     stringdefault.StaticString("INACTIVE"),
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StreamKey
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Stream-key value.",
		//	  "type": "string"
		//	}
		"stream_key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Stream-key value.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of key-value pairs that contain metadata for the asset model.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
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
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
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
			Description: "A list of key-value pairs that contain metadata for the asset model.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: UserId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "User defined indentifier for participant associated with IngestConfiguration.",
		//	  "type": "string"
		//	}
		"user_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "User defined indentifier for participant associated with IngestConfiguration.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
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
		Description: "Resource Type definition for AWS::IVS::IngestConfiguration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IVS::IngestConfiguration").WithTerraformTypeName("awscc_ivs_ingest_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":             "Arn",
		"ingest_protocol": "IngestProtocol",
		"insecure_ingest": "InsecureIngest",
		"key":             "Key",
		"name":            "Name",
		"participant_id":  "ParticipantId",
		"stage_arn":       "StageArn",
		"state":           "State",
		"stream_key":      "StreamKey",
		"tags":            "Tags",
		"user_id":         "UserId",
		"value":           "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/InsecureIngest",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
