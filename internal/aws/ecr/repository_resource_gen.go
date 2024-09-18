// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ecr

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_ecr_repository", repositoryResource)
}

// repositoryResource returns the Terraform awscc_ecr_repository resource.
// This Terraform resource corresponds to the CloudFormation AWS::ECR::Repository resource.
func repositoryResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EmptyOnDelete
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "If true, deleting the repository force deletes the contents of the repository. If false, the repository must be empty before attempting to delete it.",
		//	  "type": "boolean"
		//	}
		"empty_on_delete": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "If true, deleting the repository force deletes the contents of the repository. If false, the repository must be empty before attempting to delete it.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// EmptyOnDelete is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: EncryptionConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest.",
		//	  "properties": {
		//	    "EncryptionType": {
		//	      "description": "The encryption type to use.\n If you use the ``KMS`` encryption type, the contents of the repository will be encrypted using server-side encryption with KMSlong key stored in KMS. When you use KMS to encrypt your data, you can either use the default AWS managed KMS key for Amazon ECR, or specify your own KMS key, which you already created. For more information, see [Protecting data using server-side encryption with an key stored in (SSE-KMS)](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html) in the *Amazon Simple Storage Service Console Developer Guide*.\n If you use the ``AES256`` encryption type, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts the images in the repository using an AES-256 encryption algorithm. For more information, see [Protecting data using server-side encryption with Amazon S3-managed encryption keys (SSE-S3)](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html) in the *Amazon Simple Storage Service Console Developer Guide*.",
		//	      "enum": [
		//	        "AES256",
		//	        "KMS",
		//	        "KMS_DSSE"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "KmsKey": {
		//	      "description": "If you use the ``KMS`` encryption type, specify the KMS key to use for encryption. The alias, key ID, or full ARN of the KMS key can be specified. The key must exist in the same Region as the repository. If no key is specified, the default AWS managed KMS key for Amazon ECR will be used.",
		//	      "maxLength": 2048,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "EncryptionType"
		//	  ],
		//	  "type": "object"
		//	}
		"encryption_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EncryptionType
				"encryption_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The encryption type to use.\n If you use the ``KMS`` encryption type, the contents of the repository will be encrypted using server-side encryption with KMSlong key stored in KMS. When you use KMS to encrypt your data, you can either use the default AWS managed KMS key for Amazon ECR, or specify your own KMS key, which you already created. For more information, see [Protecting data using server-side encryption with an key stored in (SSE-KMS)](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingKMSEncryption.html) in the *Amazon Simple Storage Service Console Developer Guide*.\n If you use the ``AES256`` encryption type, Amazon ECR uses server-side encryption with Amazon S3-managed encryption keys which encrypts the images in the repository using an AES-256 encryption algorithm. For more information, see [Protecting data using server-side encryption with Amazon S3-managed encryption keys (SSE-S3)](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html) in the *Amazon Simple Storage Service Console Developer Guide*.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"AES256",
							"KMS",
							"KMS_DSSE",
						),
						fwvalidators.NotNullString(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
						stringplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: KmsKey
				"kms_key": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "If you use the ``KMS`` encryption type, specify the KMS key to use for encryption. The alias, key ID, or full ARN of the KMS key can be specified. The key must exist in the same Region as the repository. If no key is specified, the default AWS managed KMS key for Amazon ECR will be used.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 2048),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
						stringplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ImageScanningConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The image scanning configuration for the repository. This determines whether images are scanned for known vulnerabilities after being pushed to the repository.",
		//	  "properties": {
		//	    "ScanOnPush": {
		//	      "description": "The setting that determines whether images are scanned after being pushed to a repository. If set to ``true``, images will be scanned after being pushed. If this parameter is not specified, it will default to ``false`` and images will not be scanned unless a scan is manually started.",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"image_scanning_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ScanOnPush
				"scan_on_push": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "The setting that determines whether images are scanned after being pushed to a repository. If set to ``true``, images will be scanned after being pushed. If this parameter is not specified, it will default to ``false`` and images will not be scanned unless a scan is manually started.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The image scanning configuration for the repository. This determines whether images are scanned for known vulnerabilities after being pushed to the repository.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ImageTagMutability
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tag mutability setting for the repository. If this parameter is omitted, the default setting of ``MUTABLE`` will be used which will allow image tags to be overwritten. If ``IMMUTABLE`` is specified, all image tags within the repository will be immutable which will prevent them from being overwritten.",
		//	  "enum": [
		//	    "MUTABLE",
		//	    "IMMUTABLE"
		//	  ],
		//	  "type": "string"
		//	}
		"image_tag_mutability": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The tag mutability setting for the repository. If this parameter is omitted, the default setting of ``MUTABLE`` will be used which will allow image tags to be overwritten. If ``IMMUTABLE`` is specified, all image tags within the repository will be immutable which will prevent them from being overwritten.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"MUTABLE",
					"IMMUTABLE",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LifecyclePolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Creates or updates a lifecycle policy. For information about lifecycle policy syntax, see [Lifecycle policy template](https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html).",
		//	  "properties": {
		//	    "LifecyclePolicyText": {
		//	      "description": "The JSON repository policy text to apply to the repository.",
		//	      "maxLength": 30720,
		//	      "minLength": 100,
		//	      "type": "string"
		//	    },
		//	    "RegistryId": {
		//	      "description": "The AWS account ID associated with the registry that contains the repository. If you do? not specify a registry, the default registry is assumed.",
		//	      "maxLength": 12,
		//	      "minLength": 12,
		//	      "pattern": "^[0-9]{12}$",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"lifecycle_policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: LifecyclePolicyText
				"lifecycle_policy_text": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The JSON repository policy text to apply to the repository.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(100, 30720),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: RegistryId
				"registry_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The AWS account ID associated with the registry that contains the repository. If you do? not specify a registry, the default registry is assumed.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(12, 12),
						stringvalidator.RegexMatches(regexp.MustCompile("^[0-9]{12}$"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Creates or updates a lifecycle policy. For information about lifecycle policy syntax, see [Lifecycle policy template](https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html).",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RepositoryName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name to use for the repository. The repository name may be specified on its own (such as ``nginx-web-app``) or it can be prepended with a namespace to group the repository into a category (such as ``project-a/nginx-web-app``). If you don't specify a name, CFNlong generates a unique physical ID and uses that ID for the repository name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).\n The repository name must start with a letter and can only contain lowercase letters, numbers, hyphens, underscores, and forward slashes.\n  If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.",
		//	  "maxLength": 256,
		//	  "minLength": 2,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"repository_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name to use for the repository. The repository name may be specified on its own (such as ``nginx-web-app``) or it can be prepended with a namespace to group the repository into a category (such as ``project-a/nginx-web-app``). If you don't specify a name, CFNlong generates a unique physical ID and uses that ID for the repository name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).\n The repository name must start with a letter and can only contain lowercase letters, numbers, hyphens, underscores, and forward slashes.\n  If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(2, 256),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RepositoryPolicyText
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The JSON repository policy text to apply to the repository. For more information, see [Amazon ECR repository policies](https://docs.aws.amazon.com/AmazonECR/latest/userguide/repository-policy-examples.html) in the *Amazon Elastic Container Registry User Guide*.",
		//	  "type": "string"
		//	}
		"repository_policy_text": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The JSON repository policy text to apply to the repository. For more information, see [Amazon ECR repository policies](https://docs.aws.amazon.com/AmazonECR/latest/userguide/repository-policy-examples.html) in the *Amazon Elastic Container Registry User Guide*.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RepositoryUri
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"repository_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
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
		//	    "description": "The metadata to apply to a resource to help you categorize and organize them. Each tag consists of a key and a value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "One part of a key-value pair that make up a tag. A ``key`` is a general label that acts like a category for more specific tag values.",
		//	        "maxLength": 127,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "A ``value`` acts as a descriptor within a tag category (key).",
		//	        "maxLength": 255,
		//	        "minLength": 1,
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
						Description: "One part of a key-value pair that make up a tag. A ``key`` is a general label that acts like a category for more specific tag values.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 127),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A ``value`` acts as a descriptor within a tag category (key).",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 255),
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
		Description: "The ``AWS::ECR::Repository`` resource specifies an Amazon Elastic Container Registry (Amazon ECR) repository, where users can push and pull Docker images, Open Container Initiative (OCI) images, and OCI compatible artifacts. For more information, see [Amazon ECR private repositories](https://docs.aws.amazon.com/AmazonECR/latest/userguide/Repositories.html) in the *Amazon ECR User Guide*.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ECR::Repository").WithTerraformTypeName("awscc_ecr_repository")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                          "Arn",
		"empty_on_delete":              "EmptyOnDelete",
		"encryption_configuration":     "EncryptionConfiguration",
		"encryption_type":              "EncryptionType",
		"image_scanning_configuration": "ImageScanningConfiguration",
		"image_tag_mutability":         "ImageTagMutability",
		"key":                          "Key",
		"kms_key":                      "KmsKey",
		"lifecycle_policy":             "LifecyclePolicy",
		"lifecycle_policy_text":        "LifecyclePolicyText",
		"registry_id":                  "RegistryId",
		"repository_name":              "RepositoryName",
		"repository_policy_text":       "RepositoryPolicyText",
		"repository_uri":               "RepositoryUri",
		"scan_on_push":                 "ScanOnPush",
		"tags":                         "Tags",
		"value":                        "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/EmptyOnDelete",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
