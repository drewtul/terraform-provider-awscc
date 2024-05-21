// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_sagemaker_inference_experiment", inferenceExperimentResource)
}

// inferenceExperimentResource returns the Terraform awscc_sagemaker_inference_experiment resource.
// This Terraform resource corresponds to the CloudFormation AWS::SageMaker::InferenceExperiment resource.
func inferenceExperimentResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the inference experiment.",
		//	  "maxLength": 256,
		//	  "minLength": 20,
		//	  "pattern": "^arn:aws[a-z\\-]*:sagemaker:[a-z0-9\\-]*:[0-9]{12}:inference-experiment/[a-zA-Z_0-9+=,.@\\-_/]+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the inference experiment.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp at which you created the inference experiment.",
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The timestamp at which you created the inference experiment.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DataStorageConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The Amazon S3 location and configuration for storing inference request and response data.",
		//	  "properties": {
		//	    "ContentType": {
		//	      "additionalProperties": false,
		//	      "description": "Configuration specifying how to treat different headers. If no headers are specified SageMaker will by default base64 encode when capturing the data.",
		//	      "properties": {
		//	        "CsvContentTypes": {
		//	          "description": "The list of all content type headers that SageMaker will treat as CSV and capture accordingly.",
		//	          "items": {
		//	            "maxLength": 256,
		//	            "minLength": 1,
		//	            "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*/[a-zA-Z0-9](-*[a-zA-Z0-9.])*",
		//	            "type": "string"
		//	          },
		//	          "maxItems": 10,
		//	          "minItems": 1,
		//	          "type": "array"
		//	        },
		//	        "JsonContentTypes": {
		//	          "description": "The list of all content type headers that SageMaker will treat as JSON and capture accordingly.",
		//	          "items": {
		//	            "maxLength": 256,
		//	            "minLength": 1,
		//	            "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*/[a-zA-Z0-9](-*[a-zA-Z0-9.])*",
		//	            "type": "string"
		//	          },
		//	          "maxItems": 10,
		//	          "minItems": 1,
		//	          "type": "array"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "Destination": {
		//	      "description": "The Amazon S3 bucket where the inference request and response data is stored.",
		//	      "maxLength": 512,
		//	      "pattern": "^(https|s3)://([^/])/?(.*)$",
		//	      "type": "string"
		//	    },
		//	    "KmsKey": {
		//	      "description": "The AWS Key Management Service key that Amazon SageMaker uses to encrypt captured data at rest using Amazon S3 server-side encryption.",
		//	      "maxLength": 2048,
		//	      "pattern": ".*",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Destination"
		//	  ],
		//	  "type": "object"
		//	}
		"data_storage_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ContentType
				"content_type": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CsvContentTypes
						"csv_content_types": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "The list of all content type headers that SageMaker will treat as CSV and capture accordingly.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.SizeBetween(1, 10),
								listvalidator.ValueStringsAre(
									stringvalidator.LengthBetween(1, 256),
									stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9](-*[a-zA-Z0-9])*/[a-zA-Z0-9](-*[a-zA-Z0-9.])*"), ""),
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: JsonContentTypes
						"json_content_types": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "The list of all content type headers that SageMaker will treat as JSON and capture accordingly.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.SizeBetween(1, 10),
								listvalidator.ValueStringsAre(
									stringvalidator.LengthBetween(1, 256),
									stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9](-*[a-zA-Z0-9])*/[a-zA-Z0-9](-*[a-zA-Z0-9.])*"), ""),
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Configuration specifying how to treat different headers. If no headers are specified SageMaker will by default base64 encode when capturing the data.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Destination
				"destination": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The Amazon S3 bucket where the inference request and response data is stored.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(512),
						stringvalidator.RegexMatches(regexp.MustCompile("^(https|s3)://([^/])/?(.*)$"), ""),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: KmsKey
				"kms_key": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The AWS Key Management Service key that Amazon SageMaker uses to encrypt captured data at rest using Amazon S3 server-side encryption.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(2048),
						stringvalidator.RegexMatches(regexp.MustCompile(".*"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The Amazon S3 location and configuration for storing inference request and response data.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the inference experiment.",
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "pattern": ".*",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the inference experiment.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 1024),
				stringvalidator.RegexMatches(regexp.MustCompile(".*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DesiredState
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The desired state of the experiment after starting or stopping operation.",
		//	  "enum": [
		//	    "Running",
		//	    "Completed",
		//	    "Cancelled"
		//	  ],
		//	  "type": "string"
		//	}
		"desired_state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The desired state of the experiment after starting or stopping operation.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"Running",
					"Completed",
					"Cancelled",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EndpointMetadata
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The metadata of the endpoint on which the inference experiment ran.",
		//	  "properties": {
		//	    "EndpointConfigName": {
		//	      "description": "The name of the endpoint configuration.",
		//	      "maxLength": 63,
		//	      "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*",
		//	      "type": "string"
		//	    },
		//	    "EndpointName": {
		//	      "description": "The name of the endpoint used to run the inference experiment.",
		//	      "maxLength": 63,
		//	      "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*",
		//	      "type": "string"
		//	    },
		//	    "EndpointStatus": {
		//	      "description": "The status of the endpoint. For possible values of the status of an endpoint.",
		//	      "enum": [
		//	        "Creating",
		//	        "Updating",
		//	        "SystemUpdating",
		//	        "RollingBack",
		//	        "InService",
		//	        "OutOfService",
		//	        "Deleting",
		//	        "Failed"
		//	      ],
		//	      "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "EndpointName"
		//	  ],
		//	  "type": "object"
		//	}
		"endpoint_metadata": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EndpointConfigName
				"endpoint_config_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the endpoint configuration.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: EndpointName
				"endpoint_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the endpoint used to run the inference experiment.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: EndpointStatus
				"endpoint_status": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The status of the endpoint. For possible values of the status of an endpoint.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The metadata of the endpoint on which the inference experiment ran.",
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EndpointName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the endpoint used to run the inference experiment.",
		//	  "maxLength": 63,
		//	  "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*",
		//	  "type": "string"
		//	}
		"endpoint_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the endpoint used to run the inference experiment.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(63),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9](-*[a-zA-Z0-9])*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KmsKey
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.",
		//	  "maxLength": 2048,
		//	  "pattern": ".*",
		//	  "type": "string"
		//	}
		"kms_key": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(2048),
				stringvalidator.RegexMatches(regexp.MustCompile(".*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LastModifiedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The timestamp at which you last modified the inference experiment.",
		//	  "type": "string"
		//	}
		"last_modified_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The timestamp at which you last modified the inference experiment.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ModelVariants
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of ModelVariantConfig objects. Each ModelVariantConfig object in the array describes the infrastructure configuration for the corresponding variant.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Contains information about the deployment options of a model.",
		//	    "properties": {
		//	      "InfrastructureConfig": {
		//	        "additionalProperties": false,
		//	        "description": "The configuration for the infrastructure that the model will be deployed to.",
		//	        "properties": {
		//	          "InfrastructureType": {
		//	            "description": "The type of the inference experiment that you want to run.",
		//	            "enum": [
		//	              "RealTimeInference"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "RealTimeInferenceConfig": {
		//	            "additionalProperties": false,
		//	            "description": "The infrastructure configuration for deploying the model to a real-time inference endpoint.",
		//	            "properties": {
		//	              "InstanceCount": {
		//	                "description": "The number of instances of the type specified by InstanceType.",
		//	                "type": "integer"
		//	              },
		//	              "InstanceType": {
		//	                "description": "The instance type the model is deployed to.",
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "InstanceType",
		//	              "InstanceCount"
		//	            ],
		//	            "type": "object"
		//	          }
		//	        },
		//	        "required": [
		//	          "InfrastructureType",
		//	          "RealTimeInferenceConfig"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "ModelName": {
		//	        "description": "The name of the Amazon SageMaker Model entity.",
		//	        "maxLength": 63,
		//	        "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*",
		//	        "type": "string"
		//	      },
		//	      "VariantName": {
		//	        "description": "The name of the variant.",
		//	        "maxLength": 63,
		//	        "pattern": "^[a-zA-Z0-9]([\\-a-zA-Z0-9]*[a-zA-Z0-9])?",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "ModelName",
		//	      "VariantName",
		//	      "InfrastructureConfig"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 2,
		//	  "type": "array"
		//	}
		"model_variants": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: InfrastructureConfig
					"infrastructure_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: InfrastructureType
							"infrastructure_type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The type of the inference experiment that you want to run.",
								Required:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"RealTimeInference",
									),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: RealTimeInferenceConfig
							"real_time_inference_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: InstanceCount
									"instance_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
										Description: "The number of instances of the type specified by InstanceType.",
										Required:    true,
									}, /*END ATTRIBUTE*/
									// Property: InstanceType
									"instance_type": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The instance type the model is deployed to.",
										Required:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "The infrastructure configuration for deploying the model to a real-time inference endpoint.",
								Required:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "The configuration for the infrastructure that the model will be deployed to.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: ModelName
					"model_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The name of the Amazon SageMaker Model entity.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthAtMost(63),
							stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9](-*[a-zA-Z0-9])*"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: VariantName
					"variant_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The name of the variant.",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthAtMost(63),
							stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9]([\\-a-zA-Z0-9]*[a-zA-Z0-9])?"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of ModelVariantConfig objects. Each ModelVariantConfig object in the array describes the infrastructure configuration for the corresponding variant.",
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtMost(2),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name for the inference experiment.",
		//	  "maxLength": 120,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name for the inference experiment.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 120),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to access model artifacts and container images, and manage Amazon SageMaker Inference endpoints for model deployment.",
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "^arn:aws[a-z\\-]*:iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+$",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to access model artifacts and container images, and manage Amazon SageMaker Inference endpoints for model deployment.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(20, 2048),
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws[a-z\\-]*:iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Schedule
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The duration for which you want the inference experiment to run.",
		//	  "properties": {
		//	    "EndTime": {
		//	      "description": "The timestamp at which the inference experiment ended or will end.",
		//	      "type": "string"
		//	    },
		//	    "StartTime": {
		//	      "description": "The timestamp at which the inference experiment started or will start.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"schedule": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EndTime
				"end_time": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The timestamp at which the inference experiment ended or will end.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: StartTime
				"start_time": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The timestamp at which the inference experiment started or will start.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The duration for which you want the inference experiment to run.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ShadowModeConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The configuration of ShadowMode inference experiment type. Use this field to specify a production variant which takes all the inference requests, and a shadow variant to which Amazon SageMaker replicates a percentage of the inference requests. For the shadow variant also specify the percentage of requests that Amazon SageMaker replicates.",
		//	  "properties": {
		//	    "ShadowModelVariants": {
		//	      "description": "List of shadow variant configurations.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "The name and sampling percentage of a shadow variant.",
		//	        "properties": {
		//	          "SamplingPercentage": {
		//	            "description": "The percentage of inference requests that Amazon SageMaker replicates from the production variant to the shadow variant.",
		//	            "maximum": 100,
		//	            "type": "integer"
		//	          },
		//	          "ShadowModelVariantName": {
		//	            "description": "The name of the shadow variant.",
		//	            "maxLength": 63,
		//	            "pattern": "^[a-zA-Z0-9]([\\-a-zA-Z0-9]*[a-zA-Z0-9])?",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "ShadowModelVariantName",
		//	          "SamplingPercentage"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 1,
		//	      "minItems": 1,
		//	      "type": "array"
		//	    },
		//	    "SourceModelVariantName": {
		//	      "description": "The name of the production variant, which takes all the inference requests.",
		//	      "maxLength": 63,
		//	      "pattern": "^[a-zA-Z0-9]([\\-a-zA-Z0-9]*[a-zA-Z0-9])?",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "SourceModelVariantName",
		//	    "ShadowModelVariants"
		//	  ],
		//	  "type": "object"
		//	}
		"shadow_mode_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ShadowModelVariants
				"shadow_model_variants": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: SamplingPercentage
							"sampling_percentage": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "The percentage of inference requests that Amazon SageMaker replicates from the production variant to the shadow variant.",
								Required:    true,
								Validators: []validator.Int64{ /*START VALIDATORS*/
									int64validator.AtMost(100),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: ShadowModelVariantName
							"shadow_model_variant_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The name of the shadow variant.",
								Required:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9]([\\-a-zA-Z0-9]*[a-zA-Z0-9])?"), ""),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "List of shadow variant configurations.",
					Required:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeBetween(1, 1),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: SourceModelVariantName
				"source_model_variant_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the production variant, which takes all the inference requests.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(63),
						stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9]([\\-a-zA-Z0-9]*[a-zA-Z0-9])?"), ""),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The configuration of ShadowMode inference experiment type. Use this field to specify a production variant which takes all the inference requests, and a shadow variant to which Amazon SageMaker replicates a percentage of the inference requests. For the shadow variant also specify the percentage of requests that Amazon SageMaker replicates.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of the inference experiment.",
		//	  "enum": [
		//	    "Creating",
		//	    "Created",
		//	    "Updating",
		//	    "Starting",
		//	    "Stopping",
		//	    "Running",
		//	    "Completed",
		//	    "Cancelled"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of the inference experiment.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StatusReason
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The error message or client-specified reason from the StopInferenceExperiment API, that explains the status of the inference experiment.",
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "pattern": ".*",
		//	  "type": "string"
		//	}
		"status_reason": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The error message or client-specified reason from the StopInferenceExperiment API, that explains the status of the inference experiment.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 1024),
				stringvalidator.RegexMatches(regexp.MustCompile(".*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
		//	        "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
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
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							stringvalidator.RegexMatches(regexp.MustCompile("^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthAtMost(256),
							stringvalidator.RegexMatches(regexp.MustCompile("^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of the inference experiment that you want to run.",
		//	  "enum": [
		//	    "ShadowMode"
		//	  ],
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of the inference experiment that you want to run.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"ShadowMode",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
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
		Description: "Resource Type definition for AWS::SageMaker::InferenceExperiment",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::InferenceExperiment").WithTerraformTypeName("awscc_sagemaker_inference_experiment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                        "Arn",
		"content_type":               "ContentType",
		"creation_time":              "CreationTime",
		"csv_content_types":          "CsvContentTypes",
		"data_storage_config":        "DataStorageConfig",
		"description":                "Description",
		"desired_state":              "DesiredState",
		"destination":                "Destination",
		"end_time":                   "EndTime",
		"endpoint_config_name":       "EndpointConfigName",
		"endpoint_metadata":          "EndpointMetadata",
		"endpoint_name":              "EndpointName",
		"endpoint_status":            "EndpointStatus",
		"infrastructure_config":      "InfrastructureConfig",
		"infrastructure_type":        "InfrastructureType",
		"instance_count":             "InstanceCount",
		"instance_type":              "InstanceType",
		"json_content_types":         "JsonContentTypes",
		"key":                        "Key",
		"kms_key":                    "KmsKey",
		"last_modified_time":         "LastModifiedTime",
		"model_name":                 "ModelName",
		"model_variants":             "ModelVariants",
		"name":                       "Name",
		"real_time_inference_config": "RealTimeInferenceConfig",
		"role_arn":                   "RoleArn",
		"sampling_percentage":        "SamplingPercentage",
		"schedule":                   "Schedule",
		"shadow_mode_config":         "ShadowModeConfig",
		"shadow_model_variant_name":  "ShadowModelVariantName",
		"shadow_model_variants":      "ShadowModelVariants",
		"source_model_variant_name":  "SourceModelVariantName",
		"start_time":                 "StartTime",
		"status":                     "Status",
		"status_reason":              "StatusReason",
		"tags":                       "Tags",
		"type":                       "Type",
		"value":                      "Value",
		"variant_name":               "VariantName",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
