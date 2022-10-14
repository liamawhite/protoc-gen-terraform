// Copyright 2022 Liam White
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-terraform. DO NOT EDIT.
package test

import (
	"context"

	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	tfsdk "github.com/hashicorp/terraform-plugin-framework/tfsdk"
	types "github.com/hashicorp/terraform-plugin-framework/types"
)

// GenSchemaTest returns tfsdk.Schema definition for Test
func GenSchemaTest(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{Attributes: map[string]tfsdk.Attribute{
		"bool": {
			Description: "Bool bool field",
			Optional:    true,
			Type:        types.BoolType,
		},
		"branch1": {
			Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{"str": {
				Description: "Str string field",
				Optional:    true,
				Type:        types.StringType,
			}}),
			Description: "Branch1 is the first oneOf branch",
			Optional:    true,
		},
		"branch2": {
			Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{"int32": {
				Description: "Int32 int field",
				Optional:    true,
				Type:        types.Int64Type,
			}}),
			Description: "Branch2 is the second oneOf branch",
			Optional:    true,
		},
		"branch3": {
			Description: "Branch3 is the third branch which is simple string",
			Optional:    true,
			Type:        types.StringType,
		},
		"bytes": {
			Description: "Bytes byte[] field",
			Optional:    true,
			Type:        types.StringType,
		},
		"double": {
			Description: "Double double field",
			Optional:    true,
			Type:        types.Float64Type,
		},
		"float": {
			Description: "Float float field",
			Optional:    true,
			Type:        types.Float64Type,
		},
		"inject_computed": {
			Computed: true,
			Optional: false,
			Required: false,
			Type:     types.StringType,
		},
		"inject_optional": {
			Computed: false,
			Optional: true,
			Required: false,
			Type:     types.BoolType,
		},
		"inject_required": {
			Computed: false,
			Optional: false,
			Required: true,
			Type:     types.Int64Type,
		},
		"int32": {
			Description: "Int32 int32 field",
			Optional:    true,
			Type:        types.Int64Type,
		},
		"int64": {
			Description: "Int64 int64 field",
			Optional:    true,
			Type:        types.Int64Type,
		},
		"map": {
			Description: "Map normal map",
			Optional:    true,
			Type:        types.MapType{ElemType: types.StringType},
		},
		"mode": {
			Description: "Mode is the enum value",
			Optional:    true,
			Type:        types.Int64Type,
		},
		"nested": {
			Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
				"map": {
					Description: "Nested map repeated nested messages",
					Optional:    true,
					Type:        types.MapType{ElemType: types.StringType},
				},
				"map_object_nested": {
					Attributes: tfsdk.MapNestedAttributes(map[string]tfsdk.Attribute{"str": {
						Description: "Str string field",
						Optional:    true,
						Type:        types.StringType,
					}}),
					Description: "MapObjectNested nested object map",
					Optional:    true,
				},
				"other_nested_list": {
					Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{"str": {
						Description: "Str string field",
						Optional:    true,
						Type:        types.StringType,
					}}),
					Description: "Nested repeated nested messages",
					Optional:    true,
				},
				"str": {
					Description: "Str string field",
					Optional:    true,
					Type:        types.StringType,
				},
			}),
			Description: "Nested nested message field, non-nullable",
			Optional:    true,
		},
		"nested_list": {
			Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
				"map": {
					Description: "Nested map repeated nested messages",
					Optional:    true,
					Type:        types.MapType{ElemType: types.StringType},
				},
				"map_object_nested": {
					Attributes: tfsdk.MapNestedAttributes(map[string]tfsdk.Attribute{"str": {
						Description: "Str string field",
						Optional:    true,
						Type:        types.StringType,
					}}),
					Description: "MapObjectNested nested object map",
					Optional:    true,
				},
				"other_nested_list": {
					Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{"str": {
						Description: "Str string field",
						Optional:    true,
						Type:        types.StringType,
					}}),
					Description: "Nested repeated nested messages",
					Optional:    true,
				},
				"str": {
					Description: "Str string field",
					Optional:    true,
					Type:        types.StringType,
				},
			}),
			Description: "NestedList nested message array",
			Optional:    true,
		},
		"nested_map": {
			Attributes: tfsdk.MapNestedAttributes(map[string]tfsdk.Attribute{
				"map": {
					Description: "Nested map repeated nested messages",
					Optional:    true,
					Type:        types.MapType{ElemType: types.StringType},
				},
				"map_object_nested": {
					Attributes: tfsdk.MapNestedAttributes(map[string]tfsdk.Attribute{"str": {
						Description: "Str string field",
						Optional:    true,
						Type:        types.StringType,
					}}),
					Description: "MapObjectNested nested object map",
					Optional:    true,
				},
				"other_nested_list": {
					Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{"str": {
						Description: "Str string field",
						Optional:    true,
						Type:        types.StringType,
					}}),
					Description: "Nested repeated nested messages",
					Optional:    true,
				},
				"str": {
					Description: "Str string field",
					Optional:    true,
					Type:        types.StringType,
				},
			}),
			Description: "MapObject is the object map",
			Optional:    true,
		},
		"required": {
			Description: "Required string field",
			Required:    true,
			Type:        types.StringType,
		},
		"str": {
			Description: "Str string field",
			Optional:    true,
			Type:        types.StringType,
		},
		"string_list": {
			Description: "StringList []string field",
			Optional:    true,
			Type:        types.ListType{ElemType: types.StringType},
		},
	}}, nil
}

// GenSchemaEmptyMessageBranch returns tfsdk.Schema definition for EmptyMessageBranch
func GenSchemaEmptyMessageBranch(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{Attributes: map[string]tfsdk.Attribute{}}, nil
}

// GenSchemaNested returns tfsdk.Schema definition for Nested
func GenSchemaNested(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{Attributes: map[string]tfsdk.Attribute{
		"map": {
			Description: "Nested map repeated nested messages",
			Optional:    true,
			Type:        types.MapType{ElemType: types.StringType},
		},
		"map_object_nested": {
			Attributes: tfsdk.MapNestedAttributes(map[string]tfsdk.Attribute{"str": {
				Description: "Str string field",
				Optional:    true,
				Type:        types.StringType,
			}}),
			Description: "MapObjectNested nested object map",
			Optional:    true,
		},
		"other_nested_list": {
			Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{"str": {
				Description: "Str string field",
				Optional:    true,
				Type:        types.StringType,
			}}),
			Description: "Nested repeated nested messages",
			Optional:    true,
		},
		"str": {
			Description: "Str string field",
			Optional:    true,
			Type:        types.StringType,
		},
	}}, nil
}

// GenSchemaOtherNested returns tfsdk.Schema definition for OtherNested
func GenSchemaOtherNested(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{Attributes: map[string]tfsdk.Attribute{"str": {
		Description: "Str string field",
		Optional:    true,
		Type:        types.StringType,
	}}}, nil
}

// GenSchemaBranch1 returns tfsdk.Schema definition for Branch1
func GenSchemaBranch1(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{Attributes: map[string]tfsdk.Attribute{"str": {
		Description: "Str string field",
		Optional:    true,
		Type:        types.StringType,
	}}}, nil
}

// GenSchemaBranch2 returns tfsdk.Schema definition for Branch2
func GenSchemaBranch2(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{Attributes: map[string]tfsdk.Attribute{"int32": {
		Description: "Int32 int field",
		Optional:    true,
		Type:        types.Int64Type,
	}}}, nil
}
