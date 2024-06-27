# Add worker secret

## Overview

This command is used to edit a worker manifest to add secret that can be used for deployment or/and execution.

|                        |                                                         |
|------------------------|---------------------------------------------------------|
| Command name           | worker add-secret                                       |
| Abbreviation           | worker as                                               |
| **Command options:**   |                                                         |
| `--edit`               | \[Default: false] Whether to update an existing secret. |
| **Command arguments:** |                                                         |
| secret-name            | The secret name                                         |

## Example

Add a secret name `my-secret` to a worker initialized in the current directory.

```
jfrog worker add-secret my-secret
```