package shim

import (
	"github.com/1Password/pulumi-onepassword/provider/pkg/version"
	"github.com/1Password/terraform-provider-onepassword/v2/providerlink"
	tfpf "github.com/hashicorp/terraform-plugin-framework/provider"
)

func NewProvider() tfpf.Provider {
	return providerlink.New(version.Version)()
}
