// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_verified_access_instance", verifiedAccessInstanceResource)
}

// verifiedAccessInstanceResource returns the Terraform awscc_ec2_verified_access_instance resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::VerifiedAccessInstance resource.
func verifiedAccessInstanceResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Time this Verified Access Instance was created.",
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Time this Verified Access Instance was created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description for the AWS Verified Access instance.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description for the AWS Verified Access instance.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FipsEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether FIPS is enabled",
		//	  "type": "boolean"
		//	}
		"fips_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether FIPS is enabled",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LastUpdatedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Time this Verified Access Instance was last updated.",
		//	  "type": "string"
		//	}
		"last_updated_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Time this Verified Access Instance was last updated.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LoggingConfigurations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The configuration options for AWS Verified Access instances.",
		//	  "properties": {
		//	    "CloudWatchLogs": {
		//	      "additionalProperties": false,
		//	      "description": "Sends Verified Access logs to CloudWatch Logs.",
		//	      "properties": {
		//	        "Enabled": {
		//	          "description": "Indicates whether logging is enabled.",
		//	          "type": "boolean"
		//	        },
		//	        "LogGroup": {
		//	          "description": "The ID of the CloudWatch Logs log group.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "IncludeTrustContext": {
		//	      "description": "Include claims from trust providers in Verified Access logs.",
		//	      "type": "boolean"
		//	    },
		//	    "KinesisDataFirehose": {
		//	      "additionalProperties": false,
		//	      "description": "Sends Verified Access logs to Kinesis.",
		//	      "properties": {
		//	        "DeliveryStream": {
		//	          "description": "The ID of the delivery stream.",
		//	          "type": "string"
		//	        },
		//	        "Enabled": {
		//	          "description": "Indicates whether logging is enabled.",
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "LogVersion": {
		//	      "description": "Select log version for Verified Access logs.",
		//	      "type": "string"
		//	    },
		//	    "S3": {
		//	      "additionalProperties": false,
		//	      "description": "Sends Verified Access logs to Amazon S3.",
		//	      "properties": {
		//	        "BucketName": {
		//	          "description": "The bucket name.",
		//	          "type": "string"
		//	        },
		//	        "BucketOwner": {
		//	          "description": "The ID of the AWS account that owns the Amazon S3 bucket.",
		//	          "type": "string"
		//	        },
		//	        "Enabled": {
		//	          "description": "Indicates whether logging is enabled.",
		//	          "type": "boolean"
		//	        },
		//	        "Prefix": {
		//	          "description": "The bucket prefix.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"logging_configurations": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CloudWatchLogs
				"cloudwatch_logs": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "Indicates whether logging is enabled.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
								boolplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: LogGroup
						"log_group": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ID of the CloudWatch Logs log group.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Sends Verified Access logs to CloudWatch Logs.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: IncludeTrustContext
				"include_trust_context": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Include claims from trust providers in Verified Access logs.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: KinesisDataFirehose
				"kinesis_data_firehose": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: DeliveryStream
						"delivery_stream": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ID of the delivery stream.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "Indicates whether logging is enabled.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
								boolplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Sends Verified Access logs to Kinesis.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: LogVersion
				"log_version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Select log version for Verified Access logs.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: S3
				"s3": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: BucketName
						"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The bucket name.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: BucketOwner
						"bucket_owner": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The ID of the AWS account that owns the Amazon S3 bucket.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "Indicates whether logging is enabled.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
								boolplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Prefix
						"prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The bucket prefix.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Sends Verified Access logs to Amazon S3.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The configuration options for AWS Verified Access instances.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
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
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VerifiedAccessInstanceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the AWS Verified Access instance.",
		//	  "type": "string"
		//	}
		"verified_access_instance_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the AWS Verified Access instance.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VerifiedAccessTrustProviderIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IDs of the AWS Verified Access trust providers.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "The ID of the AWS Verified Access trust provider.",
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"verified_access_trust_provider_ids": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The IDs of the AWS Verified Access trust providers.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VerifiedAccessTrustProviders
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "AWS Verified Access trust providers.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A Verified Access Trust Provider.",
		//	    "properties": {
		//	      "Description": {
		//	        "description": "The description of trust provider.",
		//	        "type": "string"
		//	      },
		//	      "DeviceTrustProviderType": {
		//	        "description": "The type of device-based trust provider.",
		//	        "type": "string"
		//	      },
		//	      "TrustProviderType": {
		//	        "description": "The type of trust provider (user- or device-based).",
		//	        "type": "string"
		//	      },
		//	      "UserTrustProviderType": {
		//	        "description": "The type of user-based trust provider.",
		//	        "type": "string"
		//	      },
		//	      "VerifiedAccessTrustProviderId": {
		//	        "description": "The ID of the trust provider.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"verified_access_trust_providers": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Description
					"description": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The description of trust provider.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: DeviceTrustProviderType
					"device_trust_provider_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The type of device-based trust provider.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: TrustProviderType
					"trust_provider_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The type of trust provider (user- or device-based).",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: UserTrustProviderType
					"user_trust_provider_type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The type of user-based trust provider.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: VerifiedAccessTrustProviderId
					"verified_access_trust_provider_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The ID of the trust provider.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "AWS Verified Access trust providers.",
			Optional:    true,
			Computed:    true,
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
		Description: "The AWS::EC2::VerifiedAccessInstance resource creates an AWS EC2 Verified Access Instance.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::VerifiedAccessInstance").WithTerraformTypeName("awscc_ec2_verified_access_instance")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bucket_name":                        "BucketName",
		"bucket_owner":                       "BucketOwner",
		"cloudwatch_logs":                    "CloudWatchLogs",
		"creation_time":                      "CreationTime",
		"delivery_stream":                    "DeliveryStream",
		"description":                        "Description",
		"device_trust_provider_type":         "DeviceTrustProviderType",
		"enabled":                            "Enabled",
		"fips_enabled":                       "FipsEnabled",
		"include_trust_context":              "IncludeTrustContext",
		"key":                                "Key",
		"kinesis_data_firehose":              "KinesisDataFirehose",
		"last_updated_time":                  "LastUpdatedTime",
		"log_group":                          "LogGroup",
		"log_version":                        "LogVersion",
		"logging_configurations":             "LoggingConfigurations",
		"prefix":                             "Prefix",
		"s3":                                 "S3",
		"tags":                               "Tags",
		"trust_provider_type":                "TrustProviderType",
		"user_trust_provider_type":           "UserTrustProviderType",
		"value":                              "Value",
		"verified_access_instance_id":        "VerifiedAccessInstanceId",
		"verified_access_trust_provider_id":  "VerifiedAccessTrustProviderId",
		"verified_access_trust_provider_ids": "VerifiedAccessTrustProviderIds",
		"verified_access_trust_providers":    "VerifiedAccessTrustProviders",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
