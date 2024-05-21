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
pip install pulumi-onepassword
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

<!-- TODO: Confirm that this path to the Go SDK is correct. -->
<!-- This should be straightforward to confirm once we build out everything to the sdk/ directory. -->

```bash
go get github.com/1Password/pulumi-onepassword/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

<!-- TODO: Confirm that this namespace and package name are correct for .NET. -->

```bash
dotnet add package 1Password.Pulumi-OnePassword
```

## Configuration

<!-- TODO: Add configuration options specific to the provider. -->

The following configuration points are available for the `1Password` provider:

- `foo:apiKey` (environment: `FOO_API_KEY`) - the API key for `foo`
- `foo:region` (environment: `FOO_REGION`) - the region in which to deploy resources

## Reference

<!-- TODO: Confirm that this URL to the API documentation is correct. -->

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/1Password/api-docs/).
