// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_appstream_app_block", appBlockDataSource)
}

// appBlockDataSource returns the Terraform awscc_appstream_app_block data source.
// This Terraform data source corresponds to the CloudFormation AWS::AppStream::AppBlock resource.
func appBlockDataSource(ctx context.Context) (datasource.DataSource, error) {
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
		// Property: CreatedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: DisplayName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"display_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PackagingType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"packaging_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PostSetupScriptDetails
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "ExecutableParameters": {
		//	      "type": "string"
		//	    },
		//	    "ExecutablePath": {
		//	      "type": "string"
		//	    },
		//	    "ScriptS3Location": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "S3Bucket": {
		//	          "type": "string"
		//	        },
		//	        "S3Key": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "S3Bucket"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "TimeoutInSeconds": {
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "required": [
		//	    "ScriptS3Location",
		//	    "ExecutablePath",
		//	    "TimeoutInSeconds"
		//	  ],
		//	  "type": "object"
		//	}
		"post_setup_script_details": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ExecutableParameters
				"executable_parameters": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: ExecutablePath
				"executable_path": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: ScriptS3Location
				"script_s3_location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: S3Bucket
						"s3_bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: S3Key
						"s3_key": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: TimeoutInSeconds
				"timeout_in_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SetupScriptDetails
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "ExecutableParameters": {
		//	      "type": "string"
		//	    },
		//	    "ExecutablePath": {
		//	      "type": "string"
		//	    },
		//	    "ScriptS3Location": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "S3Bucket": {
		//	          "type": "string"
		//	        },
		//	        "S3Key": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "S3Bucket"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "TimeoutInSeconds": {
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "required": [
		//	    "ScriptS3Location",
		//	    "ExecutablePath",
		//	    "TimeoutInSeconds"
		//	  ],
		//	  "type": "object"
		//	}
		"setup_script_details": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ExecutableParameters
				"executable_parameters": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: ExecutablePath
				"executable_path": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: ScriptS3Location
				"script_s3_location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: S3Bucket
						"s3_bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: S3Key
						"s3_key": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: TimeoutInSeconds
				"timeout_in_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: SourceS3Location
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "S3Bucket": {
		//	      "type": "string"
		//	    },
		//	    "S3Key": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "S3Bucket"
		//	  ],
		//	  "type": "object"
		//	}
		"source_s3_location": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: S3Bucket
				"s3_bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: S3Key
				"s3_key": schema.StringAttribute{ /*START ATTRIBUTE*/
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
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "TagKey": {
		//	        "type": "string"
		//	      },
		//	      "TagValue": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
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
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: TagKey
					"tag_key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: TagValue
					"tag_value": schema.StringAttribute{ /*START ATTRIBUTE*/
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
		Description: "Data Source schema for AWS::AppStream::AppBlock",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::AppStream::AppBlock").WithTerraformTypeName("awscc_appstream_app_block")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                       "Arn",
		"created_time":              "CreatedTime",
		"description":               "Description",
		"display_name":              "DisplayName",
		"executable_parameters":     "ExecutableParameters",
		"executable_path":           "ExecutablePath",
		"key":                       "Key",
		"name":                      "Name",
		"packaging_type":            "PackagingType",
		"post_setup_script_details": "PostSetupScriptDetails",
		"s3_bucket":                 "S3Bucket",
		"s3_key":                    "S3Key",
		"script_s3_location":        "ScriptS3Location",
		"setup_script_details":      "SetupScriptDetails",
		"source_s3_location":        "SourceS3Location",
		"tag_key":                   "TagKey",
		"tag_value":                 "TagValue",
		"tags":                      "Tags",
		"timeout_in_seconds":        "TimeoutInSeconds",
		"value":                     "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
