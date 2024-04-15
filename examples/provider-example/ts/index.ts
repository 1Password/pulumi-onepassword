import * as pulumi from "@pulumi/pulumi";
import * as onepassword from "@1password/pulumi-onepassword";

const config = new pulumi.Config();
const itemUuid = config.require("OP_ITEM_UUID");
const vaultUuid = config.require("OP_VAULT_UUID");

// Fetching login details
onepassword
  .getItem({
    uuid: itemUuid,
    vault: vaultUuid,
  })
  .then((c) => {
    pulumi.log.info(c.password);
  });
