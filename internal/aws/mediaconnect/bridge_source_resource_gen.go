// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package mediaconnect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_mediaconnect_bridge_source", bridgeSourceResource)
}

// bridgeSourceResource returns the Terraform awscc_mediaconnect_bridge_source resource.
// This Terraform resource corresponds to the CloudFormation AWS::MediaConnect::BridgeSource resource.
func bridgeSourceResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: BridgeArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Number (ARN) of the bridge.",
		//	  "type": "string"
		//	}
		"bridge_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Number (ARN) of the bridge.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: FlowSource
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The source of the bridge. A flow source originates in MediaConnect as an existing cloud flow.",
		//	  "properties": {
		//	    "FlowArn": {
		//	      "description": "The ARN of the cloud flow used as a source of this bridge.",
		//	      "type": "string"
		//	    },
		//	    "FlowVpcInterfaceAttachment": {
		//	      "additionalProperties": false,
		//	      "description": "The name of the VPC interface attachment to use for this source.",
		//	      "properties": {
		//	        "VpcInterfaceName": {
		//	          "description": "The name of the VPC interface to use for this resource.",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "FlowArn"
		//	  ],
		//	  "type": "object"
		//	}
		"flow_source": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: FlowArn
				"flow_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the cloud flow used as a source of this bridge.",
					Required:    true,
				}, /*END ATTRIBUTE*/
				// Property: FlowVpcInterfaceAttachment
				"flow_vpc_interface_attachment": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: VpcInterfaceName
						"vpc_interface_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The name of the VPC interface to use for this resource.",
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The name of the VPC interface attachment to use for this source.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The source of the bridge. A flow source originates in MediaConnect as an existing cloud flow.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the source.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the source.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NetworkSource
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The source of the bridge. A network source originates at your premises.",
		//	  "properties": {
		//	    "MulticastIp": {
		//	      "description": "The network source multicast IP.",
		//	      "type": "string"
		//	    },
		//	    "NetworkName": {
		//	      "description": "The network source's gateway network name.",
		//	      "type": "string"
		//	    },
		//	    "Port": {
		//	      "description": "The network source port.",
		//	      "type": "integer"
		//	    },
		//	    "Protocol": {
		//	      "description": "The network source protocol.",
		//	      "enum": [
		//	        "rtp-fec",
		//	        "rtp",
		//	        "udp"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Protocol",
		//	    "MulticastIp",
		//	    "Port",
		//	    "NetworkName"
		//	  ],
		//	  "type": "object"
		//	}
		"network_source": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: MulticastIp
				"multicast_ip": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The network source multicast IP.",
					Required:    true,
				}, /*END ATTRIBUTE*/
				// Property: NetworkName
				"network_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The network source's gateway network name.",
					Required:    true,
				}, /*END ATTRIBUTE*/
				// Property: Port
				"port": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "The network source port.",
					Required:    true,
				}, /*END ATTRIBUTE*/
				// Property: Protocol
				"protocol": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The network source protocol.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"rtp-fec",
							"rtp",
							"udp",
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The source of the bridge. A network source originates at your premises.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::MediaConnect::BridgeSource",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaConnect::BridgeSource").WithTerraformTypeName("awscc_mediaconnect_bridge_source")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bridge_arn":                    "BridgeArn",
		"flow_arn":                      "FlowArn",
		"flow_source":                   "FlowSource",
		"flow_vpc_interface_attachment": "FlowVpcInterfaceAttachment",
		"multicast_ip":                  "MulticastIp",
		"name":                          "Name",
		"network_name":                  "NetworkName",
		"network_source":                "NetworkSource",
		"port":                          "Port",
		"protocol":                      "Protocol",
		"vpc_interface_name":            "VpcInterfaceName",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
