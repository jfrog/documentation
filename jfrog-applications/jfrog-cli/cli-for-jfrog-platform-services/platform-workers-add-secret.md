# Add a secret to a worker manifest

## Overview

This command is used to edit a worker manifest in order to add or edit secret that can be used for deployment or/and execution.

Secrets are store encrypted with a master password that will be requested by the command.

Once secrets are added to the manifest the master password will be required by the `deploy` and `test-run` commands.

|                        |                       |
|------------------------|---------------------------------------------------------|
| Command name           | worker add-secret     |
| Abbreviation           | worker as             |
| **Command options:**   |                                                         |
| `--edit`               | \[Default: false] Whether to update an existing secret. |
| **Command arguments:** |                                                         |
| secret-name            | The secret name                                         |

## Example

Add a secret name `my-secret` to a worker initialized in the current directory.

```
jf worker add-secret my-secret
```
