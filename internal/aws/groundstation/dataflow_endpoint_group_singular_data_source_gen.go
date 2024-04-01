// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package groundstation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_groundstation_dataflow_endpoint_group", dataflowEndpointGroupDataSource)
}

// dataflowEndpointGroupDataSource returns the Terraform awscc_groundstation_dataflow_endpoint_group data source.
// This Terraform data source corresponds to the CloudFormation AWS::GroundStation::DataflowEndpointGroup resource.
func dataflowEndpointGroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ContactPostPassDurationSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amount of time, in seconds, after a contact ends that the Ground Station Dataflow Endpoint Group will be in a POSTPASS state. A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the POSTPASS state.",
		//	  "type": "integer"
		//	}
		"contact_post_pass_duration_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Amount of time, in seconds, after a contact ends that the Ground Station Dataflow Endpoint Group will be in a POSTPASS state. A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the POSTPASS state.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ContactPrePassDurationSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Amount of time, in seconds, before a contact starts that the Ground Station Dataflow Endpoint Group will be in a PREPASS state. A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the PREPASS state.",
		//	  "type": "integer"
		//	}
		"contact_pre_pass_duration_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Amount of time, in seconds, before a contact starts that the Ground Station Dataflow Endpoint Group will be in a PREPASS state. A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the PREPASS state.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EndpointDetails
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "oneOf": [
		//	      {
		//	        "required": [
		//	          "Endpoint",
		//	          "SecurityDetails"
		//	        ]
		//	      },
		//	      {
		//	        "required": [
		//	          "AwsGroundStationAgentEndpoint"
		//	        ]
		//	      }
		//	    ],
		//	    "properties": {
		//	      "AwsGroundStationAgentEndpoint": {
		//	        "additionalProperties": false,
		//	        "description": "Information about AwsGroundStationAgentEndpoint.",
		//	        "properties": {
		//	          "AgentStatus": {
		//	            "description": "The status of AgentEndpoint.",
		//	            "enum": [
		//	              "SUCCESS",
		//	              "FAILED",
		//	              "ACTIVE",
		//	              "INACTIVE"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "AuditResults": {
		//	            "description": "The results of the audit.",
		//	            "enum": [
		//	              "HEALTHY",
		//	              "UNHEALTHY"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "EgressAddress": {
		//	            "additionalProperties": false,
		//	            "description": "Egress address of AgentEndpoint with an optional mtu.",
		//	            "properties": {
		//	              "Mtu": {
		//	                "description": "Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.",
		//	                "type": "integer"
		//	              },
		//	              "SocketAddress": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "Name": {
		//	                    "type": "string"
		//	                  },
		//	                  "Port": {
		//	                    "type": "integer"
		//	                  }
		//	                },
		//	                "type": "object"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "IngressAddress": {
		//	            "additionalProperties": false,
		//	            "description": "Ingress address of AgentEndpoint with a port range and an optional mtu.",
		//	            "properties": {
		//	              "Mtu": {
		//	                "description": "Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.",
		//	                "type": "integer"
		//	              },
		//	              "SocketAddress": {
		//	                "additionalProperties": false,
		//	                "description": "A socket address with a port range.",
		//	                "properties": {
		//	                  "Name": {
		//	                    "description": "IPv4 socket address.",
		//	                    "type": "string"
		//	                  },
		//	                  "PortRange": {
		//	                    "additionalProperties": false,
		//	                    "description": "Port range of a socket address.",
		//	                    "properties": {
		//	                      "Maximum": {
		//	                        "description": "A maximum value.",
		//	                        "type": "integer"
		//	                      },
		//	                      "Minimum": {
		//	                        "description": "A minimum value.",
		//	                        "type": "integer"
		//	                      }
		//	                    },
		//	                    "type": "object"
		//	                  }
		//	                },
		//	                "type": "object"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "Name": {
		//	            "pattern": "^[ a-zA-Z0-9_:-]{1,256}$",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "Endpoint": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Address": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "Name": {
		//	                "type": "string"
		//	              },
		//	              "Port": {
		//	                "type": "integer"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "Mtu": {
		//	            "type": "integer"
		//	          },
		//	          "Name": {
		//	            "pattern": "^[ a-zA-Z0-9_:-]{1,256}$",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "SecurityDetails": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "RoleArn": {
		//	            "type": "string"
		//	          },
		//	          "SecurityGroupIds": {
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "type": "array"
		//	          },
		//	          "SubnetIds": {
		//	            "items": {
		//	              "type": "string"
		//	            },
		//	            "type": "array"
		//	          }
		//	        },
		//	        "type": "object"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"endpoint_details": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: AwsGroundStationAgentEndpoint
					"aws_ground_station_agent_endpoint": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: AgentStatus
							"agent_status": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The status of AgentEndpoint.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: AuditResults
							"audit_results": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The results of the audit.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: EgressAddress
							"egress_address": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Mtu
									"mtu": schema.Int64Attribute{ /*START ATTRIBUTE*/
										Description: "Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: SocketAddress
									"socket_address": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: Name
											"name": schema.StringAttribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
											// Property: Port
											"port": schema.Int64Attribute{ /*START ATTRIBUTE*/
												Computed: true,
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
										Computed: true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "Egress address of AgentEndpoint with an optional mtu.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: IngressAddress
							"ingress_address": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Mtu
									"mtu": schema.Int64Attribute{ /*START ATTRIBUTE*/
										Description: "Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: SocketAddress
									"socket_address": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: Name
											"name": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "IPv4 socket address.",
												Computed:    true,
											}, /*END ATTRIBUTE*/
											// Property: PortRange
											"port_range": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
												Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
													// Property: Maximum
													"maximum": schema.Int64Attribute{ /*START ATTRIBUTE*/
														Description: "A maximum value.",
														Computed:    true,
													}, /*END ATTRIBUTE*/
													// Property: Minimum
													"minimum": schema.Int64Attribute{ /*START ATTRIBUTE*/
														Description: "A minimum value.",
														Computed:    true,
													}, /*END ATTRIBUTE*/
												}, /*END SCHEMA*/
												Description: "Port range of a socket address.",
												Computed:    true,
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
										Description: "A socket address with a port range.",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "Ingress address of AgentEndpoint with a port range and an optional mtu.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Information about AwsGroundStationAgentEndpoint.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Endpoint
					"endpoint": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Address
							"address": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Name
									"name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
									// Property: Port
									"port": schema.Int64Attribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Mtu
							"mtu": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Name
							"name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: SecurityDetails
					"security_details": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: RoleArn
							"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
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
						}, /*END SCHEMA*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "pattern": "^[ a-zA-Z0-9\\+\\-=._:/@]{1,128}$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "pattern": "^[ a-zA-Z0-9\\+\\-=._:/@]{1,256}$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
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
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::GroundStation::DataflowEndpointGroup",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::GroundStation::DataflowEndpointGroup").WithTerraformTypeName("awscc_groundstation_dataflow_endpoint_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"address":                            "Address",
		"agent_status":                       "AgentStatus",
		"arn":                                "Arn",
		"audit_results":                      "AuditResults",
		"aws_ground_station_agent_endpoint":  "AwsGroundStationAgentEndpoint",
		"contact_post_pass_duration_seconds": "ContactPostPassDurationSeconds",
		"contact_pre_pass_duration_seconds":  "ContactPrePassDurationSeconds",
		"dataflow_endpoint_group_id":         "Id",
		"egress_address":                     "EgressAddress",
		"endpoint":                           "Endpoint",
		"endpoint_details":                   "EndpointDetails",
		"ingress_address":                    "IngressAddress",
		"key":                                "Key",
		"maximum":                            "Maximum",
		"minimum":                            "Minimum",
		"mtu":                                "Mtu",
		"name":                               "Name",
		"port":                               "Port",
		"port_range":                         "PortRange",
		"role_arn":                           "RoleArn",
		"security_details":                   "SecurityDetails",
		"security_group_ids":                 "SecurityGroupIds",
		"socket_address":                     "SocketAddress",
		"subnet_ids":                         "SubnetIds",
		"tags":                               "Tags",
		"value":                              "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
