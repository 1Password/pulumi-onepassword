import * as pulumi from "@pulumi/pulumi";
import * as onepassword from "@pulumi/onepassword";

// Fetching login details
onepassword
  .getItem({
    uuid: "kg5v6kiqn6tv3izfufxtbcxvhe",
    vault: "65vh6pscsij7jx54uwgpmfqr5m",
  })
  .then((c) => {
    pulumi.log.info(c.password);
  });
