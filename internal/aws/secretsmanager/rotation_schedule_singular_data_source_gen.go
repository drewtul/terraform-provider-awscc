// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package secretsmanager

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_secretsmanager_rotation_schedule", rotationScheduleDataSource)
}

// rotationScheduleDataSource returns the Terraform awscc_secretsmanager_rotation_schedule data source.
// This Terraform data source corresponds to the CloudFormation AWS::SecretsManager::RotationSchedule resource.
func rotationScheduleDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: HostedRotationLambda
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Creates a new Lambda rotation function based on one of the Secrets Manager rotation function templates. To use a rotation function that already exists, specify RotationLambdaARN instead.",
		//	  "properties": {
		//	    "ExcludeCharacters": {
		//	      "description": "A string of the characters that you don't want in the password.",
		//	      "type": "string"
		//	    },
		//	    "KmsKeyArn": {
		//	      "description": "The ARN of the KMS key that Secrets Manager uses to encrypt the secret. If you don't specify this value, then Secrets Manager uses the key aws/secretsmanager. If aws/secretsmanager doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.",
		//	      "type": "string"
		//	    },
		//	    "MasterSecretArn": {
		//	      "description": "The ARN of the secret that contains superuser credentials, if you use the alternating users rotation strategy. CloudFormation grants the execution role for the Lambda rotation function GetSecretValue permission to the secret in this property.",
		//	      "type": "string"
		//	    },
		//	    "MasterSecretKmsKeyArn": {
		//	      "description": "The ARN of the KMS key that Secrets Manager used to encrypt the superuser secret, if you use the alternating users strategy and the superuser secret is encrypted with a customer managed key. You don't need to specify this property if the superuser secret is encrypted using the key aws/secretsmanager. CloudFormation grants the execution role for the Lambda rotation function Decrypt, DescribeKey, and GenerateDataKey permission to the key in this property.",
		//	      "type": "string"
		//	    },
		//	    "RotationLambdaName": {
		//	      "description": "The name of the Lambda rotation function.",
		//	      "type": "string"
		//	    },
		//	    "RotationType": {
		//	      "description": "The type of rotation template to use",
		//	      "type": "string"
		//	    },
		//	    "Runtime": {
		//	      "description": "The python runtime associated with the Lambda function",
		//	      "type": "string"
		//	    },
		//	    "SuperuserSecretArn": {
		//	      "description": "The ARN of the secret that contains superuser credentials, if you use the alternating users rotation strategy. CloudFormation grants the execution role for the Lambda rotation function GetSecretValue permission to the secret in this property.",
		//	      "type": "string"
		//	    },
		//	    "SuperuserSecretKmsKeyArn": {
		//	      "description": "The ARN of the KMS key that Secrets Manager used to encrypt the superuser secret, if you use the alternating users strategy and the superuser secret is encrypted with a customer managed key. You don't need to specify this property if the superuser secret is encrypted using the key aws/secretsmanager. CloudFormation grants the execution role for the Lambda rotation function Decrypt, DescribeKey, and GenerateDataKey permission to the key in this property.",
		//	      "type": "string"
		//	    },
		//	    "VpcSecurityGroupIds": {
		//	      "description": "A comma-separated list of security group IDs applied to the target database.",
		//	      "type": "string"
		//	    },
		//	    "VpcSubnetIds": {
		//	      "description": "A comma separated list of VPC subnet IDs of the target database network. The Lambda rotation function is in the same subnet group.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "RotationType"
		//	  ],
		//	  "type": "object"
		//	}
		"hosted_rotation_lambda": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ExcludeCharacters
				"exclude_characters": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A string of the characters that you don't want in the password.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: KmsKeyArn
				"kms_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the KMS key that Secrets Manager uses to encrypt the secret. If you don't specify this value, then Secrets Manager uses the key aws/secretsmanager. If aws/secretsmanager doesn't yet exist, then Secrets Manager creates it for you automatically the first time it encrypts the secret value.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MasterSecretArn
				"master_secret_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the secret that contains superuser credentials, if you use the alternating users rotation strategy. CloudFormation grants the execution role for the Lambda rotation function GetSecretValue permission to the secret in this property.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: MasterSecretKmsKeyArn
				"master_secret_kms_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the KMS key that Secrets Manager used to encrypt the superuser secret, if you use the alternating users strategy and the superuser secret is encrypted with a customer managed key. You don't need to specify this property if the superuser secret is encrypted using the key aws/secretsmanager. CloudFormation grants the execution role for the Lambda rotation function Decrypt, DescribeKey, and GenerateDataKey permission to the key in this property.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RotationLambdaName
				"rotation_lambda_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the Lambda rotation function.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: RotationType
				"rotation_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The type of rotation template to use",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Runtime
				"runtime": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The python runtime associated with the Lambda function",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SuperuserSecretArn
				"superuser_secret_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the secret that contains superuser credentials, if you use the alternating users rotation strategy. CloudFormation grants the execution role for the Lambda rotation function GetSecretValue permission to the secret in this property.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SuperuserSecretKmsKeyArn
				"superuser_secret_kms_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the KMS key that Secrets Manager used to encrypt the superuser secret, if you use the alternating users strategy and the superuser secret is encrypted with a customer managed key. You don't need to specify this property if the superuser secret is encrypted using the key aws/secretsmanager. CloudFormation grants the execution role for the Lambda rotation function Decrypt, DescribeKey, and GenerateDataKey permission to the key in this property.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: VpcSecurityGroupIds
				"vpc_security_group_ids": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A comma-separated list of security group IDs applied to the target database.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: VpcSubnetIds
				"vpc_subnet_ids": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A comma separated list of VPC subnet IDs of the target database network. The Lambda rotation function is in the same subnet group.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Creates a new Lambda rotation function based on one of the Secrets Manager rotation function templates. To use a rotation function that already exists, specify RotationLambdaARN instead.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the secret.",
		//	  "type": "string"
		//	}
		"rotation_schedule_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the secret.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RotateImmediatelyOnUpdate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies whether to rotate the secret immediately or wait until the next scheduled rotation window.",
		//	  "type": "boolean"
		//	}
		"rotate_immediately_on_update": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether to rotate the secret immediately or wait until the next scheduled rotation window.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RotationLambdaARN
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of an existing Lambda rotation function. To specify a rotation function that is also defined in this template, use the Ref function.",
		//	  "type": "string"
		//	}
		"rotation_lambda_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of an existing Lambda rotation function. To specify a rotation function that is also defined in this template, use the Ref function.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: RotationRules
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A structure that defines the rotation configuration for this secret.",
		//	  "properties": {
		//	    "AutomaticallyAfterDays": {
		//	      "description": "The number of days between automatic scheduled rotations of the secret. You can use this value to check that your secret meets your compliance guidelines for how often secrets must be rotated.",
		//	      "type": "integer"
		//	    },
		//	    "Duration": {
		//	      "description": "The length of the rotation window in hours, for example 3h for a three hour window. Secrets Manager rotates your secret at any time during this window. The window must not extend into the next rotation window or the next UTC day. The window starts according to the ScheduleExpression. If you don't specify a Duration, for a ScheduleExpression in hours, the window automatically closes after one hour. For a ScheduleExpression in days, the window automatically closes at the end of the UTC day.",
		//	      "type": "string"
		//	    },
		//	    "ScheduleExpression": {
		//	      "description": "A cron() or rate() expression that defines the schedule for rotating your secret. Secrets Manager rotation schedules use UTC time zone.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"rotation_rules": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AutomaticallyAfterDays
				"automatically_after_days": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The number of days between automatic scheduled rotations of the secret. You can use this value to check that your secret meets your compliance guidelines for how often secrets must be rotated.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Duration
				"duration": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The length of the rotation window in hours, for example 3h for a three hour window. Secrets Manager rotates your secret at any time during this window. The window must not extend into the next rotation window or the next UTC day. The window starts according to the ScheduleExpression. If you don't specify a Duration, for a ScheduleExpression in hours, the window automatically closes after one hour. For a ScheduleExpression in days, the window automatically closes at the end of the UTC day.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ScheduleExpression
				"schedule_expression": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "A cron() or rate() expression that defines the schedule for rotating your secret. Secrets Manager rotation schedules use UTC time zone.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "A structure that defines the rotation configuration for this secret.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SecretId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN or name of the secret to rotate.",
		//	  "type": "string"
		//	}
		"secret_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN or name of the secret to rotate.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SecretsManager::RotationSchedule",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SecretsManager::RotationSchedule").WithTerraformTypeName("awscc_secretsmanager_rotation_schedule")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"automatically_after_days":     "AutomaticallyAfterDays",
		"duration":                     "Duration",
		"exclude_characters":           "ExcludeCharacters",
		"hosted_rotation_lambda":       "HostedRotationLambda",
		"kms_key_arn":                  "KmsKeyArn",
		"master_secret_arn":            "MasterSecretArn",
		"master_secret_kms_key_arn":    "MasterSecretKmsKeyArn",
		"rotate_immediately_on_update": "RotateImmediatelyOnUpdate",
		"rotation_lambda_arn":          "RotationLambdaARN",
		"rotation_lambda_name":         "RotationLambdaName",
		"rotation_rules":               "RotationRules",
		"rotation_schedule_id":         "Id",
		"rotation_type":                "RotationType",
		"runtime":                      "Runtime",
		"schedule_expression":          "ScheduleExpression",
		"secret_id":                    "SecretId",
		"superuser_secret_arn":         "SuperuserSecretArn",
		"superuser_secret_kms_key_arn": "SuperuserSecretKmsKeyArn",
		"vpc_security_group_ids":       "VpcSecurityGroupIds",
		"vpc_subnet_ids":               "VpcSubnetIds",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}