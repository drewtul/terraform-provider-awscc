// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_customer_gateway", customerGatewayResource)
}

// customerGatewayResource returns the Terraform awscc_ec2_customer_gateway resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::CustomerGateway resource.
func customerGatewayResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: BgpAsn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": 65000,
		//	  "description": "For customer gateway devices that support BGP, specify the device's ASN. You must specify either ``BgpAsn`` or ``BgpAsnExtended`` when creating the customer gateway. If the ASN is larger than ``2,147,483,647``, you must use ``BgpAsnExtended``.\n Default: 65000\n Valid values: ``1`` to ``2,147,483,647``",
		//	  "type": "integer"
		//	}
		"bgp_asn": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "For customer gateway devices that support BGP, specify the device's ASN. You must specify either ``BgpAsn`` or ``BgpAsnExtended`` when creating the customer gateway. If the ASN is larger than ``2,147,483,647``, you must use ``BgpAsnExtended``.\n Default: 65000\n Valid values: ``1`` to ``2,147,483,647``",
			Optional:    true,
			Computed:    true,
			Default:     int64default.StaticInt64(65000),
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
				int64planmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BgpAsnExtended
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "For customer gateway devices that support BGP, specify the device's ASN. You must specify either ``BgpAsn`` or ``BgpAsnExtended`` when creating the customer gateway. If the ASN is larger than ``2,147,483,647``, you must use ``BgpAsnExtended``.\n Valid values: ``2,147,483,648`` to ``4,294,967,295``",
		//	  "maximum": 4294967294,
		//	  "minimum": 2147483648,
		//	  "type": "number"
		//	}
		"bgp_asn_extended": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Description: "For customer gateway devices that support BGP, specify the device's ASN. You must specify either ``BgpAsn`` or ``BgpAsnExtended`` when creating the customer gateway. If the ASN is larger than ``2,147,483,647``, you must use ``BgpAsnExtended``.\n Valid values: ``2,147,483,648`` to ``4,294,967,295``",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Float64{ /*START VALIDATORS*/
				float64validator.Between(2147483648.000000, 4294967294.000000),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
				float64planmodifier.UseStateForUnknown(),
				float64planmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CertificateArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) for the customer gateway certificate.",
		//	  "pattern": "^arn:(aws[a-zA-Z-]*)?:acm:[a-z]{2}((-gov)|(-iso(b?)))?-[a-z]+-\\d{1}:\\d{12}:certificate\\/[a-zA-Z0-9-_]+$",
		//	  "type": "string"
		//	}
		"certificate_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) for the customer gateway certificate.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:(aws[a-zA-Z-]*)?:acm:[a-z]{2}((-gov)|(-iso(b?)))?-[a-z]+-\\d{1}:\\d{12}:certificate\\/[a-zA-Z0-9-_]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CustomerGatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"customer_gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DeviceName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of customer gateway device.",
		//	  "type": "string"
		//	}
		"device_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of customer gateway device.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IpAddress
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "IPv4 address for the customer gateway device's outside interface. The address must be static. If ``OutsideIpAddressType`` in your VPN connection options is set to ``PrivateIpv4``, you can use an RFC6598 or RFC1918 private IPv4 address. If ``OutsideIpAddressType`` is set to ``PublicIpv4``, you can use a public IPv4 address.",
		//	  "type": "string"
		//	}
		"ip_address": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "IPv4 address for the customer gateway device's outside interface. The address must be static. If ``OutsideIpAddressType`` in your VPN connection options is set to ``PrivateIpv4``, you can use an RFC6598 or RFC1918 private IPv4 address. If ``OutsideIpAddressType`` is set to ``PublicIpv4``, you can use a public IPv4 address.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "One or more tags for the customer gateway.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Specifies a tag. For more information, see [Resource tags](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The tag key.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The tag value.",
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
						Description: "The tag key.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag value.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "One or more tags for the customer gateway.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of VPN connection that this customer gateway supports (``ipsec.1``).",
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of VPN connection that this customer gateway supports (``ipsec.1``).",
			Required:    true,
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
		Description: "Specifies a customer gateway.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::CustomerGateway").WithTerraformTypeName("awscc_ec2_customer_gateway")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bgp_asn":             "BgpAsn",
		"bgp_asn_extended":    "BgpAsnExtended",
		"certificate_arn":     "CertificateArn",
		"customer_gateway_id": "CustomerGatewayId",
		"device_name":         "DeviceName",
		"ip_address":          "IpAddress",
		"key":                 "Key",
		"tags":                "Tags",
		"type":                "Type",
		"value":               "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
