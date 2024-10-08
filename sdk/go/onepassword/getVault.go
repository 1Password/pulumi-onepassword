// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package onepassword

import (
	"context"
	"reflect"

	"github.com/1Password/pulumi-onepassword/sdk/go/onepassword/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get details of a vault by either its name or uuid.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/1Password/pulumi-onepassword/sdk/go/onepassword"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := onepassword.GetVault(ctx, &onepassword.GetVaultArgs{
//				Name: pulumi.StringRef("your-vault-name"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetVault(ctx *pulumi.Context, args *GetVaultArgs, opts ...pulumi.InvokeOption) (*GetVaultResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVaultResult
	err := ctx.Invoke("onepassword:index/getVault:getVault", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVault.
type GetVaultArgs struct {
	// The name of the vault to retrieve. This field will be populated with the name of the vault if the vault it looked up by its UUID.
	Name *string `pulumi:"name"`
	// The UUID of the vault to retrieve. This field will be populated with the UUID of the vault if the vault it looked up by its name.
	Uuid *string `pulumi:"uuid"`
}

// A collection of values returned by getVault.
type GetVaultResult struct {
	// The description of the vault.
	Description string `pulumi:"description"`
	Id          string `pulumi:"id"`
	// The name of the vault to retrieve. This field will be populated with the name of the vault if the vault it looked up by its UUID.
	Name string `pulumi:"name"`
	// The UUID of the vault to retrieve. This field will be populated with the UUID of the vault if the vault it looked up by its name.
	Uuid string `pulumi:"uuid"`
}

func GetVaultOutput(ctx *pulumi.Context, args GetVaultOutputArgs, opts ...pulumi.InvokeOption) GetVaultResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVaultResult, error) {
			args := v.(GetVaultArgs)
			r, err := GetVault(ctx, &args, opts...)
			var s GetVaultResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVaultResultOutput)
}

// A collection of arguments for invoking getVault.
type GetVaultOutputArgs struct {
	// The name of the vault to retrieve. This field will be populated with the name of the vault if the vault it looked up by its UUID.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The UUID of the vault to retrieve. This field will be populated with the UUID of the vault if the vault it looked up by its name.
	Uuid pulumi.StringPtrInput `pulumi:"uuid"`
}

func (GetVaultOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVaultArgs)(nil)).Elem()
}

// A collection of values returned by getVault.
type GetVaultResultOutput struct{ *pulumi.OutputState }

func (GetVaultResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVaultResult)(nil)).Elem()
}

func (o GetVaultResultOutput) ToGetVaultResultOutput() GetVaultResultOutput {
	return o
}

func (o GetVaultResultOutput) ToGetVaultResultOutputWithContext(ctx context.Context) GetVaultResultOutput {
	return o
}

// The description of the vault.
func (o GetVaultResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetVaultResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetVaultResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVaultResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the vault to retrieve. This field will be populated with the name of the vault if the vault it looked up by its UUID.
func (o GetVaultResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetVaultResult) string { return v.Name }).(pulumi.StringOutput)
}

// The UUID of the vault to retrieve. This field will be populated with the UUID of the vault if the vault it looked up by its name.
func (o GetVaultResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v GetVaultResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVaultResultOutput{})
}
