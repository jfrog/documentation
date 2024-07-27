# Initialize a Worker

## Overview

This command is used to initialize a new platform worker.

|                        |                                                                                                                                                                                    |
|------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command name           | worker undeploy                                                                                                                                                                    |
| Abbreviation           | worker rm                                                                                                                                                                          |
| **Command options:**   |                                                                                                                                                                                    |
| `--server-id`          | \[Optional] Server ID configured using the config command.                                                                                                                         |
| `--timeout-ms`         | \[Default: 5000] The request timeout in milliseconds.                                                                                                                              |
| **Command arguments:** |                                                                                                                                                                                    |
| worker-key             | \[Optional]  The worker key. If not provided it will be read from the `manifest.json` in the current directory.                                                                                                                                                            |

## Example

Undeploy a worker name `my-worker` from a server identified by `my-server`.

```
jfrog worker undeploy --server-id my-server my-worker
```