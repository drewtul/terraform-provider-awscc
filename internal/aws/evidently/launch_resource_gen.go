// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package evidently

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
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
	registry.AddResourceFactory("awscc_evidently_launch", launchResource)
}

// launchResource returns the Terraform awscc_evidently_launch resource.
// This Terraform resource corresponds to the CloudFormation AWS::Evidently::Launch resource.
func launchResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "arn:[^:]*:[^:]*:[^:]*:[^:]*:project/[-a-zA-Z0-9._]*/launch/[-a-zA-Z0-9._]*",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 160,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 160),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ExecutionStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Start or Stop Launch Launch. Default is not started.",
		//	  "properties": {
		//	    "DesiredState": {
		//	      "description": "Provide CANCELLED or COMPLETED as the launch desired state. Defaults to Completed if not provided.",
		//	      "type": "string"
		//	    },
		//	    "Reason": {
		//	      "description": "Provide a reason for stopping the launch. Defaults to empty if not provided.",
		//	      "type": "string"
		//	    },
		//	    "Status": {
		//	      "description": "Provide START or STOP action to apply on a launch",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Status"
		//	  ],
		//	  "type": "object"
		//	}
		"execution_status": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DesiredState
				"desired_state": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Provide CANCELLED or COMPLETED as the launch desired state. Defaults to Completed if not provided.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Reason
				"reason": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Provide a reason for stopping the launch. Defaults to empty if not provided.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Status
				"status": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Provide START or STOP action to apply on a launch",
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
			Description: "Start or Stop Launch Launch. Default is not started.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Groups
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Description": {
		//	        "maxLength": 160,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      },
		//	      "Feature": {
		//	        "type": "string"
		//	      },
		//	      "GroupName": {
		//	        "maxLength": 127,
		//	        "minLength": 1,
		//	        "pattern": "[-a-zA-Z0-9._]*",
		//	        "type": "string"
		//	      },
		//	      "Variation": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "GroupName",
		//	      "Feature",
		//	      "Variation"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 5,
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"groups": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Description
					"description": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 160),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Feature
					"feature": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
					}, /*END ATTRIBUTE*/
					// Property: GroupName
					"group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 127),
							stringvalidator.RegexMatches(regexp.MustCompile("[-a-zA-Z0-9._]*"), ""),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Variation
					"variation": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Required: true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 5),
				listvalidator.UniqueValues(),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: MetricMonitors
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "EntityIdKey": {
		//	        "description": "The JSON path to reference the entity id in the event.",
		//	        "type": "string"
		//	      },
		//	      "EventPattern": {
		//	        "description": "Event patterns have the same structure as the events they match. Rules use event patterns to select events. An event pattern either matches an event or it doesn't.",
		//	        "type": "string"
		//	      },
		//	      "MetricName": {
		//	        "maxLength": 255,
		//	        "minLength": 1,
		//	        "pattern": "^[\\S]+$",
		//	        "type": "string"
		//	      },
		//	      "UnitLabel": {
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": ".*",
		//	        "type": "string"
		//	      },
		//	      "ValueKey": {
		//	        "description": "The JSON path to reference the numerical metric value in the event.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "MetricName",
		//	      "EntityIdKey",
		//	      "ValueKey"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 3,
		//	  "minItems": 0,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"metric_monitors": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: EntityIdKey
					"entity_id_key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The JSON path to reference the entity id in the event.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: EventPattern
					"event_pattern": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Event patterns have the same structure as the events they match. Rules use event patterns to select events. An event pattern either matches an event or it doesn't.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: MetricName
					"metric_name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 255),
							stringvalidator.RegexMatches(regexp.MustCompile("^[\\S]+$"), ""),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: UnitLabel
					"unit_label": schema.StringAttribute{ /*START ATTRIBUTE*/
						Optional: true,
						Computed: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
							stringvalidator.RegexMatches(regexp.MustCompile(".*"), ""),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: ValueKey
					"value_key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The JSON path to reference the numerical metric value in the event.",
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
			Optional: true,
			Computed: true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(0, 3),
				listvalidator.UniqueValues(),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 127,
		//	  "minLength": 1,
		//	  "pattern": "[-a-zA-Z0-9._]*",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 127),
				stringvalidator.RegexMatches(regexp.MustCompile("[-a-zA-Z0-9._]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Project
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 2048,
		//	  "minLength": 0,
		//	  "pattern": "([-a-zA-Z0-9._]*)|(arn:[^:]*:[^:]*:[^:]*:[^:]*:project/[-a-zA-Z0-9._]*)",
		//	  "type": "string"
		//	}
		"project": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 2048),
				stringvalidator.RegexMatches(regexp.MustCompile("([-a-zA-Z0-9._]*)|(arn:[^:]*:[^:]*:[^:]*:[^:]*:project/[-a-zA-Z0-9._]*)"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RandomizationSalt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 127,
		//	  "minLength": 0,
		//	  "pattern": ".*",
		//	  "type": "string"
		//	}
		"randomization_salt": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 127),
				stringvalidator.RegexMatches(regexp.MustCompile(".*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ScheduledSplitsConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "GroupWeights": {
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "GroupName": {
		//	              "maxLength": 127,
		//	              "minLength": 1,
		//	              "pattern": "[-a-zA-Z0-9._]*",
		//	              "type": "string"
		//	            },
		//	            "SplitWeight": {
		//	              "type": "integer"
		//	            }
		//	          },
		//	          "required": [
		//	            "GroupName",
		//	            "SplitWeight"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "type": "array",
		//	        "uniqueItems": true
		//	      },
		//	      "SegmentOverrides": {
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "EvaluationOrder": {
		//	              "type": "integer"
		//	            },
		//	            "Segment": {
		//	              "maxLength": 2048,
		//	              "minLength": 1,
		//	              "pattern": "([-a-zA-Z0-9._]*)|(arn:[^:]*:[^:]*:[^:]*:[^:]*:segment/[-a-zA-Z0-9._]*)",
		//	              "type": "string"
		//	            },
		//	            "Weights": {
		//	              "insertionOrder": false,
		//	              "items": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "GroupName": {
		//	                    "maxLength": 127,
		//	                    "minLength": 1,
		//	                    "pattern": "[-a-zA-Z0-9._]*",
		//	                    "type": "string"
		//	                  },
		//	                  "SplitWeight": {
		//	                    "type": "integer"
		//	                  }
		//	                },
		//	                "required": [
		//	                  "GroupName",
		//	                  "SplitWeight"
		//	                ],
		//	                "type": "object"
		//	              },
		//	              "type": "array",
		//	              "uniqueItems": true
		//	            }
		//	          },
		//	          "required": [
		//	            "Segment",
		//	            "EvaluationOrder",
		//	            "Weights"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "type": "array",
		//	        "uniqueItems": true
		//	      },
		//	      "StartTime": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "StartTime",
		//	      "GroupWeights"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 6,
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"scheduled_splits_config": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: GroupWeights
					"group_weights": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: GroupName
								"group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
									Required: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(1, 127),
										stringvalidator.RegexMatches(regexp.MustCompile("[-a-zA-Z0-9._]*"), ""),
									}, /*END VALIDATORS*/
								}, /*END ATTRIBUTE*/
								// Property: SplitWeight
								"split_weight": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Required: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Required: true,
					}, /*END ATTRIBUTE*/
					// Property: SegmentOverrides
					"segment_overrides": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: EvaluationOrder
								"evaluation_order": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									Validators: []validator.Int64{ /*START VALIDATORS*/
										fwvalidators.NotNullInt64(),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
										int64planmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Segment
								"segment": schema.StringAttribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(1, 2048),
										stringvalidator.RegexMatches(regexp.MustCompile("([-a-zA-Z0-9._]*)|(arn:[^:]*:[^:]*:[^:]*:[^:]*:segment/[-a-zA-Z0-9._]*)"), ""),
										fwvalidators.NotNullString(),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Weights
								"weights": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
									NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: GroupName
											"group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
												Optional: true,
												Computed: true,
												Validators: []validator.String{ /*START VALIDATORS*/
													stringvalidator.LengthBetween(1, 127),
													stringvalidator.RegexMatches(regexp.MustCompile("[-a-zA-Z0-9._]*"), ""),
													fwvalidators.NotNullString(),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
													stringplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
											// Property: SplitWeight
											"split_weight": schema.Int64Attribute{ /*START ATTRIBUTE*/
												Optional: true,
												Computed: true,
												Validators: []validator.Int64{ /*START VALIDATORS*/
													fwvalidators.NotNullInt64(),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
													int64planmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
									}, /*END NESTED OBJECT*/
									Optional: true,
									Computed: true,
									Validators: []validator.Set{ /*START VALIDATORS*/
										fwvalidators.NotNullSet(),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
										setplanmodifier.UseStateForUnknown(),
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
					// Property: StartTime
					"start_time": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Required: true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 6),
				listvalidator.UniqueValues(),
			}, /*END VALIDATORS*/
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
		//	        "pattern": "",
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
		Description: "Resource Type definition for AWS::Evidently::Launch.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Evidently::Launch").WithTerraformTypeName("awscc_evidently_launch")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                     "Arn",
		"description":             "Description",
		"desired_state":           "DesiredState",
		"entity_id_key":           "EntityIdKey",
		"evaluation_order":        "EvaluationOrder",
		"event_pattern":           "EventPattern",
		"execution_status":        "ExecutionStatus",
		"feature":                 "Feature",
		"group_name":              "GroupName",
		"group_weights":           "GroupWeights",
		"groups":                  "Groups",
		"key":                     "Key",
		"metric_monitors":         "MetricMonitors",
		"metric_name":             "MetricName",
		"name":                    "Name",
		"project":                 "Project",
		"randomization_salt":      "RandomizationSalt",
		"reason":                  "Reason",
		"scheduled_splits_config": "ScheduledSplitsConfig",
		"segment":                 "Segment",
		"segment_overrides":       "SegmentOverrides",
		"split_weight":            "SplitWeight",
		"start_time":              "StartTime",
		"status":                  "Status",
		"tags":                    "Tags",
		"unit_label":              "UnitLabel",
		"value":                   "Value",
		"value_key":               "ValueKey",
		"variation":               "Variation",
		"weights":                 "Weights",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
