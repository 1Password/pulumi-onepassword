// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package onepassword

import (
	"context"
	"reflect"

	"github.com/1Password/pulumi-onepassword/sdk/go/onepassword/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the onepassword package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// A valid account's sign-in address or ID to use biometrics unlock. Can also be sourced from `OP_ACCOUNT` environment
	// variable. Provider will use the 1Password CLI if set.
	Account pulumi.StringPtrOutput `pulumi:"account"`
	// The path to the 1Password CLI binary. Can also be sourced from `OP_CLI_PATH` environment variable. Defaults to `op`.
	OpCliPath pulumi.StringPtrOutput `pulumi:"opCliPath"`
	// A valid 1Password service account token. Can also be sourced from `OP_SERVICE_ACCOUNT_TOKEN` environment variable.
	// Provider will use the 1Password CLI if set.
	ServiceAccountToken pulumi.StringPtrOutput `pulumi:"serviceAccountToken"`
	// A valid token for your 1Password Connect server. Can also be sourced from `OP_CONNECT_TOKEN` environment variable.
	// Provider will use 1Password Connect server if set.
	Token pulumi.StringPtrOutput `pulumi:"token"`
	// The HTTP(S) URL where your 1Password Connect server can be found. Can also be sourced `OP_CONNECT_HOST` environment
	// variable. Provider will use 1Password Connect server if set.
	Url pulumi.StringPtrOutput `pulumi:"url"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if args.Account == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "OP_ACCOUNT"); d != nil {
			args.Account = pulumi.StringPtr(d.(string))
		}
	}
	if args.OpCliPath == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "OP_CLI_PATH"); d != nil {
			args.OpCliPath = pulumi.StringPtr(d.(string))
		}
	}
	if args.ServiceAccountToken == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "OP_SERVICE_ACCOUNT_TOKEN"); d != nil {
			args.ServiceAccountToken = pulumi.StringPtr(d.(string))
		}
	}
	if args.Token == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "OP_CONNECT_TOKEN"); d != nil {
			args.Token = pulumi.StringPtr(d.(string))
		}
	}
	if args.Url == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "OP_CONNECT_HOST"); d != nil {
			args.Url = pulumi.StringPtr(d.(string))
		}
	}
	if args.ServiceAccountToken != nil {
		args.ServiceAccountToken = pulumi.ToSecret(args.ServiceAccountToken).(pulumi.StringPtrInput)
	}
	if args.Token != nil {
		args.Token = pulumi.ToSecret(args.Token).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"serviceAccountToken",
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:onepassword", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// A valid account's sign-in address or ID to use biometrics unlock. Can also be sourced from `OP_ACCOUNT` environment
	// variable. Provider will use the 1Password CLI if set.
	Account *string `pulumi:"account"`
	// The path to the 1Password CLI binary. Can also be sourced from `OP_CLI_PATH` environment variable. Defaults to `op`.
	OpCliPath *string `pulumi:"opCliPath"`
	// A valid 1Password service account token. Can also be sourced from `OP_SERVICE_ACCOUNT_TOKEN` environment variable.
	// Provider will use the 1Password CLI if set.
	ServiceAccountToken *string `pulumi:"serviceAccountToken"`
	// A valid token for your 1Password Connect server. Can also be sourced from `OP_CONNECT_TOKEN` environment variable.
	// Provider will use 1Password Connect server if set.
	Token *string `pulumi:"token"`
	// The HTTP(S) URL where your 1Password Connect server can be found. Can also be sourced `OP_CONNECT_HOST` environment
	// variable. Provider will use 1Password Connect server if set.
	Url *string `pulumi:"url"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// A valid account's sign-in address or ID to use biometrics unlock. Can also be sourced from `OP_ACCOUNT` environment
	// variable. Provider will use the 1Password CLI if set.
	Account pulumi.StringPtrInput
	// The path to the 1Password CLI binary. Can also be sourced from `OP_CLI_PATH` environment variable. Defaults to `op`.
	OpCliPath pulumi.StringPtrInput
	// A valid 1Password service account token. Can also be sourced from `OP_SERVICE_ACCOUNT_TOKEN` environment variable.
	// Provider will use the 1Password CLI if set.
	ServiceAccountToken pulumi.StringPtrInput
	// A valid token for your 1Password Connect server. Can also be sourced from `OP_CONNECT_TOKEN` environment variable.
	// Provider will use 1Password Connect server if set.
	Token pulumi.StringPtrInput
	// The HTTP(S) URL where your 1Password Connect server can be found. Can also be sourced `OP_CONNECT_HOST` environment
	// variable. Provider will use 1Password Connect server if set.
	Url pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

// A valid account's sign-in address or ID to use biometrics unlock. Can also be sourced from `OP_ACCOUNT` environment
// variable. Provider will use the 1Password CLI if set.
func (o ProviderOutput) Account() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Account }).(pulumi.StringPtrOutput)
}

// The path to the 1Password CLI binary. Can also be sourced from `OP_CLI_PATH` environment variable. Defaults to `op`.
func (o ProviderOutput) OpCliPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.OpCliPath }).(pulumi.StringPtrOutput)
}

// A valid 1Password service account token. Can also be sourced from `OP_SERVICE_ACCOUNT_TOKEN` environment variable.
// Provider will use the 1Password CLI if set.
func (o ProviderOutput) ServiceAccountToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ServiceAccountToken }).(pulumi.StringPtrOutput)
}

// A valid token for your 1Password Connect server. Can also be sourced from `OP_CONNECT_TOKEN` environment variable.
// Provider will use 1Password Connect server if set.
func (o ProviderOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Token }).(pulumi.StringPtrOutput)
}

// The HTTP(S) URL where your 1Password Connect server can be found. Can also be sourced `OP_CONNECT_HOST` environment
// variable. Provider will use 1Password Connect server if set.
func (o ProviderOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Url }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
