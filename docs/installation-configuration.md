---
title: 1Password: Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the 1Password Pulumi Provider.
layout: package
---

The 1Password Pulumi provider enables accessing and managing items in your 1Password vaults.

## Installation

The 1Password Pulumi provider is available as a package in most Pulumi languages:

- JavaScript/TypeScript: [@1password/pulumi-onepassword](https://www.npmjs.com/package/@1password/pulumi-onepassword)
- Python: [pulumi_onepassword](https://pypi.org/project/pulumi-onepassword/)
- Go: [github.com/1Password/pulumi-onepassword/sdk/go/onepassword](https://pkg.go.dev/github.com/1Password/pulumi-onepassword/sdk/go/onepassword)
- .NET: _coming soon_

## Configuring Credentials

The 1Password pulumi provider needs to be configured with 1Password credentials before it can be used to access and manage items in your 1Password vault. You can configure it to use one of the following:

- Service account
  - `pulumi-onepassword:service_account_token` (environment: `OP_SERVICE_ACCOUNT_TOKEN`) - The 1Password service account token to use with 1Password CLI.
- Connect
  - `pulumi-onepassword:url` (environment: `OP_CONNECT_HOST`) - the URL where your 1Password Connect API can be found.
  - `pulumi-onepassword:token` (environment: `OP_CONNECT_TOKEN`) - the token for your Connect API.
- Personal account
  - `pulumi-onepassword:account` (environment: `OP_ACCOUNT`) - A valid account's sign-in address or ID to use with 1Password CLI and biometric unlock.

Once the credentials obtained, there are two ways to communicate your configuration to Pulumi:

1. Set the environment variables for the preferred method:

   ```sh
   $ export OP_SERVICE_ACCOUNT_TOKEN=XXXXXXXXXXXXXX
   ```

2. Set them using configuration, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:

   ```sh
   $ pulumi config set pulumi-onepassword:service_account_token XXXXXXXXXXXXXX --secret
   ```

Remember to pass `--secret` when setting any sensitive data (in this example `pulumi-onepassword:service_account_token`) so that it is properly encrypted. The complete list of configuration parameters is in the [1Password Pulumi provider README](https://github.com/1Password/pulumi-onepassword/blob/main/README.md).
