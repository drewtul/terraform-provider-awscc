// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_security_group", securityGroupDataSource)
}

// securityGroupDataSource returns the Terraform awscc_ec2_security_group data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::SecurityGroup resource.
func securityGroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: GroupDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description for the security group.",
		//	  "type": "string"
		//	}
		"group_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description for the security group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GroupId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The group ID of the specified security group.",
		//	  "type": "string"
		//	}
		"group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The group ID of the specified security group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the security group.",
		//	  "type": "string"
		//	}
		"group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the security group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The group name or group ID depending on whether the SG is created in default or specific VPC",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The group name or group ID depending on whether the SG is created in default or specific VPC",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SecurityGroupEgress
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "[VPC only] The outbound rules associated with the security group. There is a short interruption during which you cannot connect to the security group.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "CidrIp": {
		//	        "type": "string"
		//	      },
		//	      "CidrIpv6": {
		//	        "type": "string"
		//	      },
		//	      "Description": {
		//	        "type": "string"
		//	      },
		//	      "DestinationPrefixListId": {
		//	        "type": "string"
		//	      },
		//	      "DestinationSecurityGroupId": {
		//	        "type": "string"
		//	      },
		//	      "FromPort": {
		//	        "type": "integer"
		//	      },
		//	      "IpProtocol": {
		//	        "type": "string"
		//	      },
		//	      "SourceSecurityGroupId": {
		//	        "type": "string"
		//	      },
		//	      "ToPort": {
		//	        "type": "integer"
		//	      }
		//	    },
		//	    "required": [
		//	      "IpProtocol"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"security_group_egress": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: CidrIp
					"cidr_ip": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: CidrIpv6
					"cidr_ipv_6": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Description
					"description": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: DestinationPrefixListId
					"destination_prefix_list_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: DestinationSecurityGroupId
					"destination_security_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: FromPort
					"from_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: IpProtocol
					"ip_protocol": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: SourceSecurityGroupId
					"source_security_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: ToPort
					"to_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "[VPC only] The outbound rules associated with the security group. There is a short interruption during which you cannot connect to the security group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SecurityGroupIngress
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The inbound rules associated with the security group. There is a short interruption during which you cannot connect to the security group.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "CidrIp": {
		//	        "type": "string"
		//	      },
		//	      "CidrIpv6": {
		//	        "type": "string"
		//	      },
		//	      "Description": {
		//	        "type": "string"
		//	      },
		//	      "FromPort": {
		//	        "type": "integer"
		//	      },
		//	      "IpProtocol": {
		//	        "type": "string"
		//	      },
		//	      "SourcePrefixListId": {
		//	        "type": "string"
		//	      },
		//	      "SourceSecurityGroupId": {
		//	        "type": "string"
		//	      },
		//	      "SourceSecurityGroupName": {
		//	        "type": "string"
		//	      },
		//	      "SourceSecurityGroupOwnerId": {
		//	        "type": "string"
		//	      },
		//	      "ToPort": {
		//	        "type": "integer"
		//	      }
		//	    },
		//	    "required": [
		//	      "IpProtocol"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"security_group_ingress": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: CidrIp
					"cidr_ip": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: CidrIpv6
					"cidr_ipv_6": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Description
					"description": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: FromPort
					"from_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: IpProtocol
					"ip_protocol": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: SourcePrefixListId
					"source_prefix_list_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: SourceSecurityGroupId
					"source_security_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: SourceSecurityGroupName
					"source_security_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: SourceSecurityGroupOwnerId
					"source_security_group_owner_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: ToPort
					"to_port": schema.Int64Attribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The inbound rules associated with the security group. There is a short interruption during which you cannot connect to the security group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Any tags assigned to the security group.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Any tags assigned to the security group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpcId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the VPC for the security group.",
		//	  "type": "string"
		//	}
		"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the VPC for the security group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::SecurityGroup",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::SecurityGroup").WithTerraformTypeName("awscc_ec2_security_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"cidr_ip":                        "CidrIp",
		"cidr_ipv_6":                     "CidrIpv6",
		"description":                    "Description",
		"destination_prefix_list_id":     "DestinationPrefixListId",
		"destination_security_group_id":  "DestinationSecurityGroupId",
		"from_port":                      "FromPort",
		"group_description":              "GroupDescription",
		"group_id":                       "GroupId",
		"group_name":                     "GroupName",
		"ip_protocol":                    "IpProtocol",
		"key":                            "Key",
		"security_group_egress":          "SecurityGroupEgress",
		"security_group_id":              "Id",
		"security_group_ingress":         "SecurityGroupIngress",
		"source_prefix_list_id":          "SourcePrefixListId",
		"source_security_group_id":       "SourceSecurityGroupId",
		"source_security_group_name":     "SourceSecurityGroupName",
		"source_security_group_owner_id": "SourceSecurityGroupOwnerId",
		"tags":                           "Tags",
		"to_port":                        "ToPort",
		"value":                          "Value",
		"vpc_id":                         "VpcId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
