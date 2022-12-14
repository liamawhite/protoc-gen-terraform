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

package generate

const (
	// SDK represents the path to Terraform SDK package
	SDK = "github.com/hashicorp/terraform-plugin-framework/tfsdk"
	// Types represents the path to Terraform types package
	Types = "github.com/hashicorp/terraform-plugin-framework/types"
	// Diag represents the path to Terraform diag package
	Diag = "github.com/hashicorp/terraform-plugin-framework/diag"
	// Attr represents the name of Terraform attr package
	Attr = "github.com/hashicorp/terraform-plugin-framework/attr"
	// TFTypes represents the name of Terraform SDK TFTypes package
	TFTypes = "github.com/hashicorp/terraform-plugin-go/tftypes"
)
