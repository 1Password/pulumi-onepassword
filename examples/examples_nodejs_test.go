//go:build nodejs || all
// +build nodejs all

package examples

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@1password/pulumi-onepassword",
		},
	})

	return baseJS
}

func TestProviderTs(t *testing.T) {
	itemUUID := os.Getenv("OP_ITEM_UUID")
	vaultUUID := os.Getenv("OP_VAULT_UUID")

    test := getJSBaseOptions(t).
        With(integration.ProgramTestOptions{
            Dir: filepath.Join(getCwd(t), "provider-example", "ts"),
			Config: map[string]string {
				"OP_ITEM_UUID": itemUUID,
				"OP_VAULT_UUID": vaultUUID,
			},
        })
    integration.ProgramTest(t, &test)
}
