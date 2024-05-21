// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package datasync

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_datasync_location_object_storage", locationObjectStorageResource)
}

// locationObjectStorageResource returns the Terraform awscc_datasync_location_object_storage resource.
// This Terraform resource corresponds to the CloudFormation AWS::DataSync::LocationObjectStorage resource.
func locationObjectStorageResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccessKey
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Optional. The access key is used if credentials are required to access the self-managed object storage server.",
		//	  "maxLength": 200,
		//	  "minLength": 1,
		//	  "pattern": "^.+$",
		//	  "type": "string"
		//	}
		"access_key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Optional. The access key is used if credentials are required to access the self-managed object storage server.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 200),
				stringvalidator.RegexMatches(regexp.MustCompile("^.+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AgentArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the agents associated with the self-managed object storage server location.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 128,
		//	    "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:agent/agent-[0-9a-z]{17}$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 4,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"agent_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The Amazon Resource Name (ARN) of the agents associated with the self-managed object storage server location.",
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 4),
				listvalidator.ValueStringsAre(
					stringvalidator.LengthAtMost(128),
					stringvalidator.RegexMatches(regexp.MustCompile("^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:agent/agent-[0-9a-z]{17}$"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BucketName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the bucket on the self-managed object storage server.",
		//	  "maxLength": 63,
		//	  "minLength": 3,
		//	  "pattern": "^[a-zA-Z0-9_\\-\\+\\./\\(\\)\\$\\p{Zs}]+$",
		//	  "type": "string"
		//	}
		"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the bucket on the self-managed object storage server.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(3, 63),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_\\-\\+\\./\\(\\)\\$\\p{Zs}]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// BucketName is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: LocationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the location that is created.",
		//	  "maxLength": 128,
		//	  "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:location/loc-[0-9a-z]{17}$",
		//	  "type": "string"
		//	}
		"location_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the location that is created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LocationUri
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The URL of the object storage location that was described.",
		//	  "maxLength": 4356,
		//	  "pattern": "^(efs|nfs|s3|smb|fsxw|object-storage)://[a-zA-Z0-9./\\-]+$",
		//	  "type": "string"
		//	}
		"location_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The URL of the object storage location that was described.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SecretKey
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Optional. The secret key is used if credentials are required to access the self-managed object storage server.",
		//	  "maxLength": 200,
		//	  "minLength": 8,
		//	  "pattern": "^.+$",
		//	  "type": "string"
		//	}
		"secret_key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Optional. The secret key is used if credentials are required to access the self-managed object storage server.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(8, 200),
				stringvalidator.RegexMatches(regexp.MustCompile("^.+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// SecretKey is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: ServerCertificate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "X.509 PEM content containing a certificate authority or chain to trust.",
		//	  "maxLength": 32768,
		//	  "type": "string"
		//	}
		"server_certificate": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "X.509 PEM content containing a certificate authority or chain to trust.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(32768),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ServerHostname
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server.",
		//	  "maxLength": 255,
		//	  "pattern": "^(([a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9\\-]*[A-Za-z0-9])$",
		//	  "type": "string"
		//	}
		"server_hostname": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(255),
				stringvalidator.RegexMatches(regexp.MustCompile("^(([a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z0-9\\-]*[A-Za-z0-9])$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// ServerHostname is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: ServerPort
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The port that your self-managed server accepts inbound network traffic on.",
		//	  "maximum": 65536,
		//	  "minimum": 1,
		//	  "type": "integer"
		//	}
		"server_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The port that your self-managed server accepts inbound network traffic on.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Int64{ /*START VALIDATORS*/
				int64validator.Between(1, 65536),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ServerProtocol
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The protocol that the object storage server uses to communicate.",
		//	  "enum": [
		//	    "HTTPS",
		//	    "HTTP"
		//	  ],
		//	  "type": "string"
		//	}
		"server_protocol": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The protocol that the object storage server uses to communicate.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"HTTPS",
					"HTTP",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Subdirectory
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The subdirectory in the self-managed object storage server that is used to read data from.",
		//	  "maxLength": 4096,
		//	  "pattern": "^[a-zA-Z0-9_\\-\\+\\./\\(\\)\\p{Zs}]*$",
		//	  "type": "string"
		//	}
		"subdirectory": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The subdirectory in the self-managed object storage server that is used to read data from.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(4096),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_\\-\\+\\./\\(\\)\\p{Zs}]*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// Subdirectory is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key for an AWS resource tag.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": "^[a-zA-Z0-9\\s+=._:/-]+$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for an AWS resource tag.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": "^[a-zA-Z0-9\\s+=._:@/-]+$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
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
						Description: "The key for an AWS resource tag.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
							stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9\\s+=._:/-]+$"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for an AWS resource tag.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
							stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9\\s+=._:@/-]+$"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
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
		Description: "Resource schema for AWS::DataSync::LocationObjectStorage.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataSync::LocationObjectStorage").WithTerraformTypeName("awscc_datasync_location_object_storage")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_key":         "AccessKey",
		"agent_arns":         "AgentArns",
		"bucket_name":        "BucketName",
		"key":                "Key",
		"location_arn":       "LocationArn",
		"location_uri":       "LocationUri",
		"secret_key":         "SecretKey",
		"server_certificate": "ServerCertificate",
		"server_hostname":    "ServerHostname",
		"server_port":        "ServerPort",
		"server_protocol":    "ServerProtocol",
		"subdirectory":       "Subdirectory",
		"tags":               "Tags",
		"value":              "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/SecretKey",
		"/properties/BucketName",
		"/properties/ServerHostname",
		"/properties/Subdirectory",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
