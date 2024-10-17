// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package transfer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_transfer_server", serverDataSource)
}

// serverDataSource returns the Terraform awscc_transfer_server data source.
// This Terraform data source corresponds to the CloudFormation AWS::Transfer::Server resource.
func serverDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1600,
		//	  "minLength": 20,
		//	  "pattern": "^arn:\\S+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: As2ServiceManagedEgressIpAddresses
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of egress IP addresses of this server. These IP addresses are only relevant for servers that use the AS2 protocol. They are used for sending asynchronous MDNs. These IP addresses are assigned automatically when you create an AS2 server. Additionally, if you update an existing server and add the AS2 protocol, static IP addresses are assigned as well.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "pattern": "^\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}$",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"as_2_service_managed_egress_ip_addresses": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The list of egress IP addresses of this server. These IP addresses are only relevant for servers that use the AS2 protocol. They are used for sending asynchronous MDNs. These IP addresses are assigned automatically when you create an AS2 server. Additionally, if you update an existing server and add the AS2 protocol, static IP addresses are assigned as well.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Certificate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 1600,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"certificate": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Domain
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "S3",
		//	    "EFS"
		//	  ],
		//	  "type": "string"
		//	}
		"domain": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: EndpointDetails
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "AddressAllocationIds": {
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "SecurityGroupIds": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "maxLength": 20,
		//	        "minLength": 11,
		//	        "pattern": "^sg-[0-9a-f]{8,17}$",
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "SubnetIds": {
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "VpcEndpointId": {
		//	      "maxLength": 22,
		//	      "minLength": 22,
		//	      "pattern": "^vpce-[0-9a-f]{17}$",
		//	      "type": "string"
		//	    },
		//	    "VpcId": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"endpoint_details": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AddressAllocationIds
				"address_allocation_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SecurityGroupIds
				"security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: SubnetIds
				"subnet_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: VpcEndpointId
				"vpc_endpoint_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: VpcId
				"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: EndpointType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "PUBLIC",
		//	    "VPC",
		//	    "VPC_ENDPOINT"
		//	  ],
		//	  "type": "string"
		//	}
		"endpoint_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: IdentityProviderDetails
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "DirectoryId": {
		//	      "maxLength": 12,
		//	      "minLength": 12,
		//	      "pattern": "^d-[0-9a-f]{10}$",
		//	      "type": "string"
		//	    },
		//	    "Function": {
		//	      "maxLength": 170,
		//	      "minLength": 1,
		//	      "pattern": "^arn:[a-z-]+:lambda:.*$",
		//	      "type": "string"
		//	    },
		//	    "InvocationRole": {
		//	      "maxLength": 2048,
		//	      "minLength": 20,
		//	      "pattern": "^arn:.*role/\\S+$",
		//	      "type": "string"
		//	    },
		//	    "SftpAuthenticationMethods": {
		//	      "enum": [
		//	        "PASSWORD",
		//	        "PUBLIC_KEY",
		//	        "PUBLIC_KEY_OR_PASSWORD",
		//	        "PUBLIC_KEY_AND_PASSWORD"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "Url": {
		//	      "maxLength": 255,
		//	      "minLength": 0,
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"identity_provider_details": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DirectoryId
				"directory_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Function
				"function": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: InvocationRole
				"invocation_role": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: SftpAuthenticationMethods
				"sftp_authentication_methods": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Url
				"url": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: IdentityProviderType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "SERVICE_MANAGED",
		//	    "API_GATEWAY",
		//	    "AWS_DIRECTORY_SERVICE",
		//	    "AWS_LAMBDA"
		//	  ],
		//	  "type": "string"
		//	}
		"identity_provider_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: LoggingRole
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 0,
		//	  "pattern": "^(|arn:.*role/\\S+)$",
		//	  "type": "string"
		//	}
		"logging_role": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PostAuthenticationLoginBanner
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 4096,
		//	  "minLength": 0,
		//	  "pattern": "^[\\x09-\\x0D\\x20-\\x7E]*$",
		//	  "type": "string"
		//	}
		"post_authentication_login_banner": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PreAuthenticationLoginBanner
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 4096,
		//	  "minLength": 0,
		//	  "pattern": "^[\\x09-\\x0D\\x20-\\x7E]*$",
		//	  "type": "string"
		//	}
		"pre_authentication_login_banner": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ProtocolDetails
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "As2Transports": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "enum": [
		//	          "HTTP"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "maxItems": 1,
		//	      "minItems": 1,
		//	      "type": "array"
		//	    },
		//	    "PassiveIp": {
		//	      "maxLength": 15,
		//	      "minLength": 0,
		//	      "type": "string"
		//	    },
		//	    "SetStatOption": {
		//	      "enum": [
		//	        "DEFAULT",
		//	        "ENABLE_NO_OP"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "TlsSessionResumptionMode": {
		//	      "enum": [
		//	        "DISABLED",
		//	        "ENABLED",
		//	        "ENFORCED"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"protocol_details": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: As2Transports
				"as_2_transports": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: PassiveIp
				"passive_ip": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: SetStatOption
				"set_stat_option": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: TlsSessionResumptionMode
				"tls_session_resumption_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Protocols
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "enum": [
		//	      "SFTP",
		//	      "FTP",
		//	      "FTPS",
		//	      "AS2"
		//	    ],
		//	    "type": "string"
		//	  },
		//	  "maxItems": 4,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"protocols": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: S3StorageOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "DirectoryListingOptimization": {
		//	      "description": "Indicates whether optimization to directory listing on S3 servers is used. Disabled by default for compatibility.",
		//	      "enum": [
		//	        "ENABLED",
		//	        "DISABLED"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"s3_storage_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DirectoryListingOptimization
				"directory_listing_optimization": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates whether optimization to directory listing on S3 servers is used. Disabled by default for compatibility.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SecurityPolicyName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 100,
		//	  "minLength": 0,
		//	  "pattern": "^TransferSecurityPolicy-.+$",
		//	  "type": "string"
		//	}
		"security_policy_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ServerId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 19,
		//	  "minLength": 19,
		//	  "pattern": "^s-([0-9a-f]{17})$",
		//	  "type": "string"
		//	}
		"server_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: StructuredLogDestinations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 1600,
		//	    "minLength": 20,
		//	    "pattern": "^arn:\\S+$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 1,
		//	  "minItems": 0,
		//	  "type": "array"
		//	}
		"structured_log_destinations": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Computed:    true,
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
		//	        "minLength": 0,
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
		//	  "maxItems": 50,
		//	  "minItems": 1,
		//	  "type": "array"
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
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: WorkflowDetails
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "OnPartialUpload": {
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "ExecutionRole": {
		//	            "maxLength": 2048,
		//	            "minLength": 20,
		//	            "pattern": "^arn:.*role/\\S+$",
		//	            "type": "string"
		//	          },
		//	          "WorkflowId": {
		//	            "maxLength": 19,
		//	            "minLength": 19,
		//	            "pattern": "^w-([a-z0-9]{17})$",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "ExecutionRole",
		//	          "WorkflowId"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 1,
		//	      "minItems": 0,
		//	      "type": "array"
		//	    },
		//	    "OnUpload": {
		//	      "insertionOrder": true,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "ExecutionRole": {
		//	            "maxLength": 2048,
		//	            "minLength": 20,
		//	            "pattern": "^arn:.*role/\\S+$",
		//	            "type": "string"
		//	          },
		//	          "WorkflowId": {
		//	            "maxLength": 19,
		//	            "minLength": 19,
		//	            "pattern": "^w-([a-z0-9]{17})$",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "ExecutionRole",
		//	          "WorkflowId"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 1,
		//	      "minItems": 0,
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"workflow_details": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: OnPartialUpload
				"on_partial_upload": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ExecutionRole
							"execution_role": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: WorkflowId
							"workflow_id": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: OnUpload
				"on_upload": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: ExecutionRole
							"execution_role": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: WorkflowId
							"workflow_id": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Transfer::Server",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Transfer::Server").WithTerraformTypeName("awscc_transfer_server")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"address_allocation_ids": "AddressAllocationIds",
		"arn":                    "Arn",
		"as_2_service_managed_egress_ip_addresses": "As2ServiceManagedEgressIpAddresses",
		"as_2_transports":                          "As2Transports",
		"certificate":                              "Certificate",
		"directory_id":                             "DirectoryId",
		"directory_listing_optimization":           "DirectoryListingOptimization",
		"domain":                                   "Domain",
		"endpoint_details":                         "EndpointDetails",
		"endpoint_type":                            "EndpointType",
		"execution_role":                           "ExecutionRole",
		"function":                                 "Function",
		"identity_provider_details":                "IdentityProviderDetails",
		"identity_provider_type":                   "IdentityProviderType",
		"invocation_role":                          "InvocationRole",
		"key":                                      "Key",
		"logging_role":                             "LoggingRole",
		"on_partial_upload":                        "OnPartialUpload",
		"on_upload":                                "OnUpload",
		"passive_ip":                               "PassiveIp",
		"post_authentication_login_banner":         "PostAuthenticationLoginBanner",
		"pre_authentication_login_banner":          "PreAuthenticationLoginBanner",
		"protocol_details":                         "ProtocolDetails",
		"protocols":                                "Protocols",
		"s3_storage_options":                       "S3StorageOptions",
		"security_group_ids":                       "SecurityGroupIds",
		"security_policy_name":                     "SecurityPolicyName",
		"server_id":                                "ServerId",
		"set_stat_option":                          "SetStatOption",
		"sftp_authentication_methods":              "SftpAuthenticationMethods",
		"structured_log_destinations":              "StructuredLogDestinations",
		"subnet_ids":                               "SubnetIds",
		"tags":                                     "Tags",
		"tls_session_resumption_mode":              "TlsSessionResumptionMode",
		"url":                                      "Url",
		"value":                                    "Value",
		"vpc_endpoint_id":                          "VpcEndpointId",
		"vpc_id":                                   "VpcId",
		"workflow_details":                         "WorkflowDetails",
		"workflow_id":                              "WorkflowId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
