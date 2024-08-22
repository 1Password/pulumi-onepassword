# 1Password Pulumi provider

Use the 1Password Pulumi provider to access and manage items in your 1Password vaults.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @1password/pulumi-onepassword
```

or `yarn`:

```bash
yarn add @1password/pulumi-onepassword
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_onepassword
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

<!-- TODO: Confirm that this path to the Go SDK is correct. -->
<!-- This should be straightforward to confirm once we build out everything to the sdk/ directory. -->

```bash
go get github.com/1Password/pulumi-onepassword/sdk/go/...
```

<!--

### .NET

To use from .NET, install using `dotnet add package`:

TODO: Confirm that this namespace and package name are correct for .NET.

```bash
dotnet add package 1Password.Pulumi-OnePassword
```

-->

## Configuration

The following configuration points are available for the `1Password` provider:

- `pulumi-onepassword:url` (environment: `OP_CONNECT_HOST`) - the URL where your 1Password Connect API can be found
- `pulumi-onepassword:token` (environment: `OP_CONNECT_TOKEN`) - the token for your Connect API.
- `pulumi-onepassword:service_account_token` (environment: `OP_SERVICE_ACCOUNT_TOKEN`) - The 1Password service account token to use with 1Password CLI.
- `pulumi-onepassword:account` (environment: `OP_ACCOUNT`) - A valid account's sign-in address or ID to use with 1Password CLI and biometric unlock.
- `pulumi-onepassword:op_cli_path` (environment: `OP_CLI_PATH`) - The path to the 1Password CLI binary.

## Reference

<!-- TODO: Confirm that this URL to the API documentation is correct. -->

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/1Password/api-docs/).
