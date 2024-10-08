// Copyright 2016-2023, Pulumi Corporation.
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

package onepassword

import (
	"fmt"
	"path"

	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	pf "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"

	"github.com/1Password/pulumi-onepassword/provider/pkg/version"
	onepassword "github.com/1Password/terraform-provider-onepassword/v2/providerlink"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "onepassword"
	// modules:
	mainMod = "index" // the onepassword module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(resource.PropertyMap, shim.ResourceConfig) error {
	return nil
}

//go:embed cmd/pulumi-resource-onepassword/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		// Instantiate the Terraform provider
		P:                       pf.ShimProvider(onepassword.New(version.Version)()),
		Name:                    "onepassword",
		TFProviderModuleVersion: "v2",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "1Password",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "1Password",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "https://1password.com/logo-images/1password-logo-dark@2x.png",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g github://api.github.com/org/pulumi-provider-name
		PluginDownloadURL: "github://api.github.com/1Password/pulumi-onepassword",
		Description:       "Use the 1Password Pulumi provider to access and manage items in your 1Password vaults.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "onepassword", "1Password", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.developer.1password.com",
		Repository: "https://github.com/1Password/pulumi-onepassword",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg:    "1Password",
		Version:      version.Version,
		MetadataInfo: tfbridge.NewProviderMetadata(metadata),
		// Map the environment variables used in the 1Password Terraform provider schema
		// (see link below) to the Pulumi provider configuration.
		// https://github.com/1Password/terraform-provider-onepassword/blob/main/onepassword/provider.go#L40-L71
		Config: map[string]*tfbridge.SchemaInfo{
			"url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OP_CONNECT_HOST"},
				},
			},
			"token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OP_CONNECT_TOKEN"},
				},
			},
			"service_account_token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OP_SERVICE_ACCOUNT_TOKEN"},
				},
			},
			"account": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OP_ACCOUNT"},
				},
			},
			"op_cli_path": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"OP_CLI_PATH"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			// Excerpt from https://github.com/pulumi/pulumi-tf-provider-boilerplate
			// > The name of the provider is omitted from the mapped name due to
			// > the presence of namespaces in all supported Pulumi languages.
			"onepassword_item": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Item")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Excerpt from https://github.com/pulumi/pulumi-tf-provider-boilerplate
			// > Note the 'get' prefix for data sources.
			"onepassword_vault": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getVault")},
			"onepassword_item":  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getItem")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@1password/pulumi-onepassword",
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "pulumi_onepassword",
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				fmt.Sprintf("github.com/1Password/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	// These are new API's that you may opt to use to automatically compute resource
	// tokens, and apply auto aliasing for full backwards compatibility.  For more
	// information, please reference:
	// https://pkg.go.dev/github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge#ProviderInfo.ComputeTokens
	prov.MustComputeTokens(tokens.SingleModule("onepassword_", mainMod,
		tokens.MakeStandard(mainPkg)))
	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
