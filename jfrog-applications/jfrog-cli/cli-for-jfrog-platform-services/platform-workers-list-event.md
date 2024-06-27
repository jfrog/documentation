# List events

## Overview

This command list all the available events on your Artifactory.

|                        |                                                                                                                                                                                     |
|------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command name           | worker list-event                                                                                                                                                                   |
| Abbreviation           | worker le                                                                                                                                                                           |
| **Command options:**   |                                                                                                                                                                                     |
| `--server-id`          | \[Optional] Server ID configured using the config command.                                                                                                                          |
| `--timeout-ms`         | \[Default: 5000] The request timeout in milliseconds.                                                                                                                               |

## Example

List event supported by a server identified by `my-server`.

```
jfrog worker list-event --server-id my-server
```