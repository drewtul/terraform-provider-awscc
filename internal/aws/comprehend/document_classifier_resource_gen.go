// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package comprehend

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_comprehend_document_classifier", documentClassifierResource)
}

// documentClassifierResource returns the Terraform awscc_comprehend_document_classifier resource.
// This Terraform resource corresponds to the CloudFormation AWS::Comprehend::DocumentClassifier resource.
func documentClassifierResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "arn:aws(-[^:]+)?:comprehend:[a-zA-Z0-9-]*:[0-9]{12}:document-classifier/[a-zA-Z0-9](-*[a-zA-Z0-9])*(/version/[a-zA-Z0-9](-*[a-zA-Z0-9])*)?",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DataAccessRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 20,
		//	  "pattern": "arn:aws(-[^:]+)?:iam::[0-9]{12}:role/.+",
		//	  "type": "string"
		//	}
		"data_access_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(20, 2048),
				stringvalidator.RegexMatches(regexp.MustCompile("arn:aws(-[^:]+)?:iam::[0-9]{12}:role/.+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DocumentClassifierName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*$",
		//	  "type": "string"
		//	}
		"document_classifier_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 63),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9](-*[a-zA-Z0-9])*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: InputDataConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "AugmentedManifests": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "AttributeNames": {
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*",
		//	              "type": "string"
		//	            },
		//	            "maxItems": 63,
		//	            "minItems": 1,
		//	            "type": "array",
		//	            "uniqueItems": true
		//	          },
		//	          "S3Uri": {
		//	            "maxLength": 1024,
		//	            "pattern": "s3://[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9](/.*)?",
		//	            "type": "string"
		//	          },
		//	          "Split": {
		//	            "enum": [
		//	              "TRAIN",
		//	              "TEST"
		//	            ],
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "AttributeNames",
		//	          "S3Uri"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "DataFormat": {
		//	      "enum": [
		//	        "COMPREHEND_CSV",
		//	        "AUGMENTED_MANIFEST"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "DocumentReaderConfig": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "DocumentReadAction": {
		//	          "enum": [
		//	            "TEXTRACT_DETECT_DOCUMENT_TEXT",
		//	            "TEXTRACT_ANALYZE_DOCUMENT"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "DocumentReadMode": {
		//	          "enum": [
		//	            "SERVICE_DEFAULT",
		//	            "FORCE_DOCUMENT_READ_ACTION"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "FeatureTypes": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "enum": [
		//	              "TABLES",
		//	              "FORMS"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "maxItems": 2,
		//	          "minItems": 1,
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        }
		//	      },
		//	      "required": [
		//	        "DocumentReadAction"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "DocumentType": {
		//	      "enum": [
		//	        "PLAIN_TEXT_DOCUMENT",
		//	        "SEMI_STRUCTURED_DOCUMENT"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "Documents": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "S3Uri": {
		//	          "maxLength": 1024,
		//	          "pattern": "s3://[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9](/.*)?",
		//	          "type": "string"
		//	        },
		//	        "TestS3Uri": {
		//	          "maxLength": 1024,
		//	          "pattern": "s3://[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9](/.*)?",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "S3Uri"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "LabelDelimiter": {
		//	      "maxLength": 1,
		//	      "minLength": 1,
		//	      "pattern": "^[ ~!@#$%^*\\-_+=|\\\\:;\\t\u003e?/]$",
		//	      "type": "string"
		//	    },
		//	    "S3Uri": {
		//	      "maxLength": 1024,
		//	      "pattern": "s3://[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9](/.*)?",
		//	      "type": "string"
		//	    },
		//	    "TestS3Uri": {
		//	      "maxLength": 1024,
		//	      "pattern": "s3://[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9](/.*)?",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"input_data_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AugmentedManifests
				"augmented_manifests": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: AttributeNames
							"attribute_names": schema.SetAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Required:    true,
								Validators: []validator.Set{ /*START VALIDATORS*/
									setvalidator.SizeBetween(1, 63),
									setvalidator.ValueStringsAre(
										stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9](-*[a-zA-Z0-9])*"), ""),
									),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: S3Uri
							"s3_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
								Required: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.LengthAtMost(1024),
									stringvalidator.RegexMatches(regexp.MustCompile("s3://[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9](/.*)?"), ""),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: Split
							"split": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"TRAIN",
										"TEST",
									),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
						setplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: DataFormat
				"data_format": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"COMPREHEND_CSV",
							"AUGMENTED_MANIFEST",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: DocumentReaderConfig
				"document_reader_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: DocumentReadAction
						"document_read_action": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"TEXTRACT_DETECT_DOCUMENT_TEXT",
									"TEXTRACT_ANALYZE_DOCUMENT",
								),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: DocumentReadMode
						"document_read_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"SERVICE_DEFAULT",
									"FORCE_DOCUMENT_READ_ACTION",
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: FeatureTypes
						"feature_types": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Optional:    true,
							Computed:    true,
							Validators: []validator.Set{ /*START VALIDATORS*/
								setvalidator.SizeBetween(1, 2),
								setvalidator.ValueStringsAre(
									stringvalidator.OneOf(
										"TABLES",
										"FORMS",
									),
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
								setplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: DocumentType
				"document_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"PLAIN_TEXT_DOCUMENT",
							"SEMI_STRUCTURED_DOCUMENT",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Documents
				"documents": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: S3Uri
						"s3_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthAtMost(1024),
								stringvalidator.RegexMatches(regexp.MustCompile("s3://[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9](/.*)?"), ""),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
						// Property: TestS3Uri
						"test_s3_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthAtMost(1024),
								stringvalidator.RegexMatches(regexp.MustCompile("s3://[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9](/.*)?"), ""),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: LabelDelimiter
				"label_delimiter": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 1),
						stringvalidator.RegexMatches(regexp.MustCompile("^[ ~!@#$%^*\\-_+=|\\\\:;\\t>?/]$"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: S3Uri
				"s3_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(1024),
						stringvalidator.RegexMatches(regexp.MustCompile("s3://[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9](/.*)?"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: TestS3Uri
				"test_s3_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(1024),
						stringvalidator.RegexMatches(regexp.MustCompile("s3://[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9](/.*)?"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Required: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LanguageCode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "en",
		//	    "es",
		//	    "fr",
		//	    "it",
		//	    "de",
		//	    "pt"
		//	  ],
		//	  "type": "string"
		//	}
		"language_code": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"en",
					"es",
					"fr",
					"it",
					"de",
					"pt",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Mode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "MULTI_CLASS",
		//	    "MULTI_LABEL"
		//	  ],
		//	  "type": "string"
		//	}
		"mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"MULTI_CLASS",
					"MULTI_LABEL",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ModelKmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"model_kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 2048),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ModelPolicy
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 20000,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"model_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 20000),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: OutputDataConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "KmsKeyId": {
		//	      "maxLength": 2048,
		//	      "minLength": 1,
		//	      "type": "string"
		//	    },
		//	    "S3Uri": {
		//	      "maxLength": 1024,
		//	      "pattern": "s3://[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9](/.*)?",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"output_data_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: KmsKeyId
				"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 2048),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: S3Uri
				"s3_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(1024),
						stringvalidator.RegexMatches(regexp.MustCompile("s3://[a-z0-9][\\.\\-a-z0-9]{1,61}[a-z0-9](/.*)?"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
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
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VersionName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*$",
		//	  "type": "string"
		//	}
		"version_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 63),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9](-*[a-zA-Z0-9])*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VolumeKmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"volume_kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 2048),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpcConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "SecurityGroupIds": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "maxLength": 32,
		//	        "minLength": 1,
		//	        "pattern": "[-0-9a-zA-Z]+",
		//	        "type": "string"
		//	      },
		//	      "maxItems": 5,
		//	      "minItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "Subnets": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "maxLength": 32,
		//	        "minLength": 1,
		//	        "pattern": "[-0-9a-zA-Z]+",
		//	        "type": "string"
		//	      },
		//	      "maxItems": 16,
		//	      "minItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    }
		//	  },
		//	  "required": [
		//	    "SecurityGroupIds",
		//	    "Subnets"
		//	  ],
		//	  "type": "object"
		//	}
		"vpc_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: SecurityGroupIds
				"security_group_ids": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Required:    true,
					Validators: []validator.Set{ /*START VALIDATORS*/
						setvalidator.SizeBetween(1, 5),
						setvalidator.ValueStringsAre(
							stringvalidator.LengthBetween(1, 32),
							stringvalidator.RegexMatches(regexp.MustCompile("[-0-9a-zA-Z]+"), ""),
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: Subnets
				"subnets": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Required:    true,
					Validators: []validator.Set{ /*START VALIDATORS*/
						setvalidator.SizeBetween(1, 16),
						setvalidator.ValueStringsAre(
							stringvalidator.LengthBetween(1, 32),
							stringvalidator.RegexMatches(regexp.MustCompile("[-0-9a-zA-Z]+"), ""),
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplaceIfConfigured(),
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
		Description: "Document Classifier enables training document classifier models.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Comprehend::DocumentClassifier").WithTerraformTypeName("awscc_comprehend_document_classifier")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                      "Arn",
		"attribute_names":          "AttributeNames",
		"augmented_manifests":      "AugmentedManifests",
		"data_access_role_arn":     "DataAccessRoleArn",
		"data_format":              "DataFormat",
		"document_classifier_name": "DocumentClassifierName",
		"document_read_action":     "DocumentReadAction",
		"document_read_mode":       "DocumentReadMode",
		"document_reader_config":   "DocumentReaderConfig",
		"document_type":            "DocumentType",
		"documents":                "Documents",
		"feature_types":            "FeatureTypes",
		"input_data_config":        "InputDataConfig",
		"key":                      "Key",
		"kms_key_id":               "KmsKeyId",
		"label_delimiter":          "LabelDelimiter",
		"language_code":            "LanguageCode",
		"mode":                     "Mode",
		"model_kms_key_id":         "ModelKmsKeyId",
		"model_policy":             "ModelPolicy",
		"output_data_config":       "OutputDataConfig",
		"s3_uri":                   "S3Uri",
		"security_group_ids":       "SecurityGroupIds",
		"split":                    "Split",
		"subnets":                  "Subnets",
		"tags":                     "Tags",
		"test_s3_uri":              "TestS3Uri",
		"value":                    "Value",
		"version_name":             "VersionName",
		"volume_kms_key_id":        "VolumeKmsKeyId",
		"vpc_config":               "VpcConfig",
	})

	opts = opts.WithCreateTimeoutInMinutes(2160).WithDeleteTimeoutInMinutes(120)

	opts = opts.WithUpdateTimeoutInMinutes(10)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
