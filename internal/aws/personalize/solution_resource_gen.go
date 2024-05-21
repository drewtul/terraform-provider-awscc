// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package personalize

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_personalize_solution", solutionResource)
}

// solutionResource returns the Terraform awscc_personalize_solution resource.
// This Terraform resource corresponds to the CloudFormation AWS::Personalize::Solution resource.
func solutionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DatasetGroupArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the dataset group that provides the training data.",
		//	  "maxLength": 256,
		//	  "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
		//	  "type": "string"
		//	}
		"dataset_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the dataset group that provides the training data.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(256),
				stringvalidator.RegexMatches(regexp.MustCompile("arn:([a-z\\d-]+):personalize:.*:.*:.+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EventType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "When your have multiple event types (using an EVENT_TYPE schema field), this parameter specifies which event type (for example, 'click' or 'like') is used for training the model. If you do not provide an eventType, Amazon Personalize will use all interactions for training with equal weight regardless of type.",
		//	  "maxLength": 256,
		//	  "type": "string"
		//	}
		"event_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "When your have multiple event types (using an EVENT_TYPE schema field), this parameter specifies which event type (for example, 'click' or 'like') is used for training the model. If you do not provide an eventType, Amazon Personalize will use all interactions for training with equal weight regardless of type.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(256),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name for the solution",
		//	  "maxLength": 63,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9][a-zA-Z0-9\\-_]*",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name for the solution",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 63),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9][a-zA-Z0-9\\-_]*"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PerformAutoML
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Whether to perform automated machine learning (AutoML). The default is false. For this case, you must specify recipeArn.",
		//	  "type": "boolean"
		//	}
		"perform_auto_ml": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Whether to perform automated machine learning (AutoML). The default is false. For this case, you must specify recipeArn.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
				boolplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PerformHPO
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Whether to perform hyperparameter optimization (HPO) on the specified or selected recipe. The default is false. When performing AutoML, this parameter is always true and you should not set it to false.",
		//	  "type": "boolean"
		//	}
		"perform_hpo": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Whether to perform hyperparameter optimization (HPO) on the specified or selected recipe. The default is false. When performing AutoML, this parameter is always true and you should not set it to false.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
				boolplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RecipeArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the recipe to use for model training. Only specified when performAutoML is false.",
		//	  "maxLength": 256,
		//	  "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
		//	  "type": "string"
		//	}
		"recipe_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the recipe to use for model training. Only specified when performAutoML is false.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(256),
				stringvalidator.RegexMatches(regexp.MustCompile("arn:([a-z\\d-]+):personalize:.*:.*:.+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SolutionArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the solution",
		//	  "maxLength": 256,
		//	  "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
		//	  "type": "string"
		//	}
		"solution_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the solution",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SolutionConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The configuration to use with the solution. When performAutoML is set to true, Amazon Personalize only evaluates the autoMLConfig section of the solution configuration.",
		//	  "properties": {
		//	    "AlgorithmHyperParameters": {
		//	      "additionalProperties": false,
		//	      "description": "Lists the hyperparameter names and ranges.",
		//	      "patternProperties": {
		//	        "": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "AutoMLConfig": {
		//	      "additionalProperties": false,
		//	      "description": "The AutoMLConfig object containing a list of recipes to search when AutoML is performed.",
		//	      "properties": {
		//	        "MetricName": {
		//	          "description": "The metric to optimize.",
		//	          "maxLength": 256,
		//	          "type": "string"
		//	        },
		//	        "RecipeList": {
		//	          "description": "The list of candidate recipes.",
		//	          "insertionOrder": true,
		//	          "items": {
		//	            "maxLength": 256,
		//	            "pattern": "arn:([a-z\\d-]+):personalize:.*:.*:.+",
		//	            "type": "string"
		//	          },
		//	          "maxItems": 100,
		//	          "type": "array"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "EventValueThreshold": {
		//	      "description": "Only events with a value greater than or equal to this threshold are used for training a model.",
		//	      "maxLength": 256,
		//	      "type": "string"
		//	    },
		//	    "FeatureTransformationParameters": {
		//	      "additionalProperties": false,
		//	      "description": "Lists the feature transformation parameters.",
		//	      "patternProperties": {
		//	        "": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "HpoConfig": {
		//	      "additionalProperties": false,
		//	      "description": "Describes the properties for hyperparameter optimization (HPO)",
		//	      "properties": {
		//	        "AlgorithmHyperParameterRanges": {
		//	          "additionalProperties": false,
		//	          "description": "The hyperparameters and their allowable ranges",
		//	          "properties": {
		//	            "CategoricalHyperParameterRanges": {
		//	              "description": "The categorical hyperparameters and their ranges.",
		//	              "insertionOrder": true,
		//	              "items": {
		//	                "additionalProperties": false,
		//	                "description": "Provides the name and values of a Categorical hyperparameter.",
		//	                "properties": {
		//	                  "Name": {
		//	                    "description": "The name of the hyperparameter.",
		//	                    "maxLength": 256,
		//	                    "type": "string"
		//	                  },
		//	                  "Values": {
		//	                    "description": "A list of the categories for the hyperparameter.",
		//	                    "insertionOrder": true,
		//	                    "items": {
		//	                      "maxLength": 1000,
		//	                      "type": "string"
		//	                    },
		//	                    "maxItems": 100,
		//	                    "type": "array"
		//	                  }
		//	                },
		//	                "type": "object"
		//	              },
		//	              "maxItems": 100,
		//	              "type": "array"
		//	            },
		//	            "ContinuousHyperParameterRanges": {
		//	              "description": "The continuous hyperparameters and their ranges.",
		//	              "insertionOrder": true,
		//	              "items": {
		//	                "additionalProperties": false,
		//	                "description": "Provides the name and range of a continuous hyperparameter.",
		//	                "properties": {
		//	                  "MaxValue": {
		//	                    "description": "The maximum allowable value for the hyperparameter.",
		//	                    "minimum": -1000000,
		//	                    "type": "number"
		//	                  },
		//	                  "MinValue": {
		//	                    "description": "The minimum allowable value for the hyperparameter.",
		//	                    "minimum": -1000000,
		//	                    "type": "number"
		//	                  },
		//	                  "Name": {
		//	                    "description": "The name of the hyperparameter.",
		//	                    "maxLength": 256,
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "type": "object"
		//	              },
		//	              "maxItems": 100,
		//	              "type": "array"
		//	            },
		//	            "IntegerHyperParameterRanges": {
		//	              "description": "The integer hyperparameters and their ranges.",
		//	              "insertionOrder": true,
		//	              "items": {
		//	                "additionalProperties": false,
		//	                "description": "Provides the name and range of an integer-valued hyperparameter.",
		//	                "properties": {
		//	                  "MaxValue": {
		//	                    "description": "The maximum allowable value for the hyperparameter.",
		//	                    "maximum": 1000000,
		//	                    "type": "integer"
		//	                  },
		//	                  "MinValue": {
		//	                    "description": "The minimum allowable value for the hyperparameter.",
		//	                    "minimum": -1000000,
		//	                    "type": "integer"
		//	                  },
		//	                  "Name": {
		//	                    "description": "The name of the hyperparameter.",
		//	                    "maxLength": 256,
		//	                    "type": "string"
		//	                  }
		//	                },
		//	                "type": "object"
		//	              },
		//	              "maxItems": 100,
		//	              "type": "array"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "HpoObjective": {
		//	          "additionalProperties": false,
		//	          "description": "The metric to optimize during HPO.",
		//	          "properties": {
		//	            "MetricName": {
		//	              "description": "The name of the metric",
		//	              "maxLength": 256,
		//	              "type": "string"
		//	            },
		//	            "MetricRegex": {
		//	              "description": "A regular expression for finding the metric in the training job logs.",
		//	              "maxLength": 256,
		//	              "type": "string"
		//	            },
		//	            "Type": {
		//	              "description": "The type of the metric. Valid values are Maximize and Minimize.",
		//	              "enum": [
		//	                "Maximize",
		//	                "Minimize"
		//	              ],
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "HpoResourceConfig": {
		//	          "additionalProperties": false,
		//	          "description": "Describes the resource configuration for hyperparameter optimization (HPO).",
		//	          "properties": {
		//	            "MaxNumberOfTrainingJobs": {
		//	              "description": "The maximum number of training jobs when you create a solution version. The maximum value for maxNumberOfTrainingJobs is 40.",
		//	              "maxLength": 256,
		//	              "type": "string"
		//	            },
		//	            "MaxParallelTrainingJobs": {
		//	              "description": "The maximum number of parallel training jobs when you create a solution version. The maximum value for maxParallelTrainingJobs is 10.",
		//	              "maxLength": 256,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"solution_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AlgorithmHyperParameters
				"algorithm_hyper_parameters": // Pattern: ""
				schema.MapAttribute{          /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Lists the hyperparameter names and ranges.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
						mapplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: AutoMLConfig
				"auto_ml_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: MetricName
						"metric_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The metric to optimize.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthAtMost(256),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: RecipeList
						"recipe_list": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "The list of candidate recipes.",
							Optional:    true,
							Computed:    true,
							Validators: []validator.List{ /*START VALIDATORS*/
								listvalidator.SizeAtMost(100),
								listvalidator.ValueStringsAre(
									stringvalidator.LengthAtMost(256),
									stringvalidator.RegexMatches(regexp.MustCompile("arn:([a-z\\d-]+):personalize:.*:.*:.+"), ""),
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The AutoMLConfig object containing a list of recipes to search when AutoML is performed.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: EventValueThreshold
				"event_value_threshold": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Only events with a value greater than or equal to this threshold are used for training a model.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthAtMost(256),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: FeatureTransformationParameters
				"feature_transformation_parameters": // Pattern: ""
				schema.MapAttribute{                 /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Lists the feature transformation parameters.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
						mapplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: HpoConfig
				"hpo_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: AlgorithmHyperParameterRanges
						"algorithm_hyper_parameter_ranges": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: CategoricalHyperParameterRanges
								"categorical_hyper_parameter_ranges": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
									NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: Name
											"name": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "The name of the hyperparameter.",
												Optional:    true,
												Computed:    true,
												Validators: []validator.String{ /*START VALIDATORS*/
													stringvalidator.LengthAtMost(256),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
													stringplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
											// Property: Values
											"values": schema.ListAttribute{ /*START ATTRIBUTE*/
												ElementType: types.StringType,
												Description: "A list of the categories for the hyperparameter.",
												Optional:    true,
												Computed:    true,
												Validators: []validator.List{ /*START VALIDATORS*/
													listvalidator.SizeAtMost(100),
													listvalidator.ValueStringsAre(
														stringvalidator.LengthAtMost(1000),
													),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
													listplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
									}, /*END NESTED OBJECT*/
									Description: "The categorical hyperparameters and their ranges.",
									Optional:    true,
									Computed:    true,
									Validators: []validator.List{ /*START VALIDATORS*/
										listvalidator.SizeAtMost(100),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
										listplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: ContinuousHyperParameterRanges
								"continuous_hyper_parameter_ranges": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
									NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: MaxValue
											"max_value": schema.Float64Attribute{ /*START ATTRIBUTE*/
												Description: "The maximum allowable value for the hyperparameter.",
												Optional:    true,
												Computed:    true,
												Validators: []validator.Float64{ /*START VALIDATORS*/
													float64validator.AtLeast(-1000000.000000),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
													float64planmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
											// Property: MinValue
											"min_value": schema.Float64Attribute{ /*START ATTRIBUTE*/
												Description: "The minimum allowable value for the hyperparameter.",
												Optional:    true,
												Computed:    true,
												Validators: []validator.Float64{ /*START VALIDATORS*/
													float64validator.AtLeast(-1000000.000000),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
													float64planmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
											// Property: Name
											"name": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "The name of the hyperparameter.",
												Optional:    true,
												Computed:    true,
												Validators: []validator.String{ /*START VALIDATORS*/
													stringvalidator.LengthAtMost(256),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
													stringplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
									}, /*END NESTED OBJECT*/
									Description: "The continuous hyperparameters and their ranges.",
									Optional:    true,
									Computed:    true,
									Validators: []validator.List{ /*START VALIDATORS*/
										listvalidator.SizeAtMost(100),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
										listplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: IntegerHyperParameterRanges
								"integer_hyper_parameter_ranges": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
									NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: MaxValue
											"max_value": schema.Int64Attribute{ /*START ATTRIBUTE*/
												Description: "The maximum allowable value for the hyperparameter.",
												Optional:    true,
												Computed:    true,
												Validators: []validator.Int64{ /*START VALIDATORS*/
													int64validator.AtMost(1000000),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
													int64planmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
											// Property: MinValue
											"min_value": schema.Int64Attribute{ /*START ATTRIBUTE*/
												Description: "The minimum allowable value for the hyperparameter.",
												Optional:    true,
												Computed:    true,
												Validators: []validator.Int64{ /*START VALIDATORS*/
													int64validator.AtLeast(-1000000),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
													int64planmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
											// Property: Name
											"name": schema.StringAttribute{ /*START ATTRIBUTE*/
												Description: "The name of the hyperparameter.",
												Optional:    true,
												Computed:    true,
												Validators: []validator.String{ /*START VALIDATORS*/
													stringvalidator.LengthAtMost(256),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
													stringplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
									}, /*END NESTED OBJECT*/
									Description: "The integer hyperparameters and their ranges.",
									Optional:    true,
									Computed:    true,
									Validators: []validator.List{ /*START VALIDATORS*/
										listvalidator.SizeAtMost(100),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
										listplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "The hyperparameters and their allowable ranges",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: HpoObjective
						"hpo_objective": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: MetricName
								"metric_name": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The name of the metric",
									Optional:    true,
									Computed:    true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthAtMost(256),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: MetricRegex
								"metric_regex": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "A regular expression for finding the metric in the training job logs.",
									Optional:    true,
									Computed:    true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthAtMost(256),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Type
								"type": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The type of the metric. Valid values are Maximize and Minimize.",
									Optional:    true,
									Computed:    true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.OneOf(
											"Maximize",
											"Minimize",
										),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "The metric to optimize during HPO.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: HpoResourceConfig
						"hpo_resource_config": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: MaxNumberOfTrainingJobs
								"max_number_of_training_jobs": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The maximum number of training jobs when you create a solution version. The maximum value for maxNumberOfTrainingJobs is 40.",
									Optional:    true,
									Computed:    true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthAtMost(256),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: MaxParallelTrainingJobs
								"max_parallel_training_jobs": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The maximum number of parallel training jobs when you create a solution version. The maximum value for maxParallelTrainingJobs is 10.",
									Optional:    true,
									Computed:    true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthAtMost(256),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Description: "Describes the resource configuration for hyperparameter optimization (HPO).",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Describes the properties for hyperparameter optimization (HPO)",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The configuration to use with the solution. When performAutoML is set to true, Amazon Personalize only evaluates the autoMLConfig section of the solution configuration.",
			Optional:    true,
			Computed:    true,
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
		Description: "Resource schema for AWS::Personalize::Solution.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Personalize::Solution").WithTerraformTypeName("awscc_personalize_solution")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"algorithm_hyper_parameter_ranges":   "AlgorithmHyperParameterRanges",
		"algorithm_hyper_parameters":         "AlgorithmHyperParameters",
		"auto_ml_config":                     "AutoMLConfig",
		"categorical_hyper_parameter_ranges": "CategoricalHyperParameterRanges",
		"continuous_hyper_parameter_ranges":  "ContinuousHyperParameterRanges",
		"dataset_group_arn":                  "DatasetGroupArn",
		"event_type":                         "EventType",
		"event_value_threshold":              "EventValueThreshold",
		"feature_transformation_parameters":  "FeatureTransformationParameters",
		"hpo_config":                         "HpoConfig",
		"hpo_objective":                      "HpoObjective",
		"hpo_resource_config":                "HpoResourceConfig",
		"integer_hyper_parameter_ranges":     "IntegerHyperParameterRanges",
		"max_number_of_training_jobs":        "MaxNumberOfTrainingJobs",
		"max_parallel_training_jobs":         "MaxParallelTrainingJobs",
		"max_value":                          "MaxValue",
		"metric_name":                        "MetricName",
		"metric_regex":                       "MetricRegex",
		"min_value":                          "MinValue",
		"name":                               "Name",
		"perform_auto_ml":                    "PerformAutoML",
		"perform_hpo":                        "PerformHPO",
		"recipe_arn":                         "RecipeArn",
		"recipe_list":                        "RecipeList",
		"solution_arn":                       "SolutionArn",
		"solution_config":                    "SolutionConfig",
		"type":                               "Type",
		"values":                             "Values",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
