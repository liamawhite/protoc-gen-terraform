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

// GenSchemaTest2 returns tfsdk.Schema definition for Test2
func GenSchemaTest2(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{Attributes: map[string]tfsdk.Attribute{"str": {
		Description: "Str string field",
		Optional:    true,
		Type:        types.StringType,
	}}}, nil
}
