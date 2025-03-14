---
sidebar_position: 2
tags:
  - contributing
  - clients
---

# Contribute to the Clients

## TypeScript/JavaScript Clients

Without some scripting making sweeping changes in the clients would be hard because of how they depend on each other: `js/types` (`@singulatron/types`) is a dependency of `js/client` (`@singulatron/client`).

To fix this a tiny script `link_local.sh` was introduced.

Your local workflow when editing the `@singulatron/types` should be is to issue the `bash link_local.sh` in the `clients/js` folder. The script links up and builds the packages in the correct order for local testing.

### Publishing

Just bump the version number in the `package.json`s and the clients will be automatically published when merged to main.
