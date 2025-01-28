// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package codebuild

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_codebuild_fleet", fleetDataSource)
}

// fleetDataSource returns the Terraform awscc_codebuild_fleet data source.
// This Terraform data source corresponds to the CloudFormation AWS::CodeBuild::Fleet resource.
func fleetDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: BaseCapacity
		// CloudFormation resource type schema:
		//
		//	{
		//	  "minimum": 1,
		//	  "type": "integer"
		//	}
		"base_capacity": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ComputeConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "disk": {
		//	      "type": "integer"
		//	    },
		//	    "machineType": {
		//	      "enum": [
		//	        "GENERAL",
		//	        "NVME"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "memory": {
		//	      "type": "integer"
		//	    },
		//	    "vCpu": {
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"compute_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: disk
				"disk": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: machineType
				"machine_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: memory
				"memory": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: vCpu
				"v_cpu": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ComputeType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "BUILD_GENERAL1_SMALL",
		//	    "BUILD_GENERAL1_MEDIUM",
		//	    "BUILD_GENERAL1_LARGE",
		//	    "BUILD_GENERAL1_XLARGE",
		//	    "BUILD_GENERAL1_2XLARGE",
		//	    "ATTRIBUTE_BASED_COMPUTE"
		//	  ],
		//	  "type": "string"
		//	}
		"compute_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: EnvironmentType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "WINDOWS_SERVER_2019_CONTAINER",
		//	    "WINDOWS_SERVER_2022_CONTAINER",
		//	    "LINUX_CONTAINER",
		//	    "LINUX_GPU_CONTAINER",
		//	    "ARM_CONTAINER",
		//	    "MAC_ARM",
		//	    "LINUX_EC2",
		//	    "ARM_EC2",
		//	    "WINDOWS_EC2"
		//	  ],
		//	  "type": "string"
		//	}
		"environment_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: FleetProxyConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "DefaultBehavior": {
		//	      "enum": [
		//	        "ALLOW_ALL",
		//	        "DENY_ALL"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "OrderedProxyRules": {
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Effect": {
		//	            "enum": [
		//	              "ALLOW",
		//	              "DENY"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "Entities": {
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "type": "array"
		//	          },
		//	          "Type": {
		//	            "enum": [
		//	              "DOMAIN",
		//	              "IP"
		//	            ],
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"fleet_proxy_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DefaultBehavior
				"default_behavior": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: OrderedProxyRules
				"ordered_proxy_rules": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Effect
							"effect": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Entities
							"entities": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Type
							"type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: FleetServiceRole
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^(?:arn:)[a-zA-Z+-=,._:/@]+$",
		//	  "type": "string"
		//	}
		"fleet_service_role": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: FleetVpcConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "SecurityGroupIds": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "Subnets": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "VpcId": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"fleet_vpc_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: SecurityGroupIds
				"security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Subnets
				"subnets": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: VpcId
				"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ImageId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"image_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 128,
		//	  "minLength": 2,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: OverflowBehavior
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "QUEUE",
		//	    "ON_DEMAND"
		//	  ],
		//	  "type": "string"
		//	}
		"overflow_behavior": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ScalingConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "MaxCapacity": {
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    },
		//	    "ScalingType": {
		//	      "enum": [
		//	        "TARGET_TRACKING_SCALING"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "TargetTrackingScalingConfigs": {
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "MetricType": {
		//	            "enum": [
		//	              "FLEET_UTILIZATION_RATE"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "TargetValue": {
		//	            "type": "number"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"scaling_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: MaxCapacity
				"max_capacity": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: ScalingType
				"scaling_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: TargetTrackingScalingConfigs
				"target_tracking_scaling_configs": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: MetricType
							"metric_type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: TargetValue
							"target_value": schema.Float64Attribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
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
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "pattern": "[a-zA-Z+-=._:/]+$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CodeBuild::Fleet",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CodeBuild::Fleet").WithTerraformTypeName("awscc_codebuild_fleet")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                             "Arn",
		"base_capacity":                   "BaseCapacity",
		"compute_configuration":           "ComputeConfiguration",
		"compute_type":                    "ComputeType",
		"default_behavior":                "DefaultBehavior",
		"disk":                            "disk",
		"effect":                          "Effect",
		"entities":                        "Entities",
		"environment_type":                "EnvironmentType",
		"fleet_proxy_configuration":       "FleetProxyConfiguration",
		"fleet_service_role":              "FleetServiceRole",
		"fleet_vpc_config":                "FleetVpcConfig",
		"image_id":                        "ImageId",
		"key":                             "Key",
		"machine_type":                    "machineType",
		"max_capacity":                    "MaxCapacity",
		"memory":                          "memory",
		"metric_type":                     "MetricType",
		"name":                            "Name",
		"ordered_proxy_rules":             "OrderedProxyRules",
		"overflow_behavior":               "OverflowBehavior",
		"scaling_configuration":           "ScalingConfiguration",
		"scaling_type":                    "ScalingType",
		"security_group_ids":              "SecurityGroupIds",
		"subnets":                         "Subnets",
		"tags":                            "Tags",
		"target_tracking_scaling_configs": "TargetTrackingScalingConfigs",
		"target_value":                    "TargetValue",
		"type":                            "Type",
		"v_cpu":                           "vCpu",
		"value":                           "Value",
		"vpc_id":                          "VpcId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
