# Deploy a worker

## Overview

This command is used to update the worker definition (code, description , filter, secret ...) on your Artifactory instance.

|                        |                                                                                                                                                                                     |
|------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command name           | worker deploy                                                                                                                                                                       |
| Abbreviation           | worker d                                                                                                                                                                            |
| **Command options:**   |                                                                                                                                                                                     |
| `--server-id`          | \[Optional] Server ID configured using the config command.                                                                                                                          |
| `--timeout-ms`         | \[Default: 5000] The request timeout in milliseconds.                                                                                                                               |
| `--no-secrets`         | \[Default: false] Do not use registered secrets.                                                                                                                                 |

## Example

Deploy a worker to server with id `my-server`.

```
jfrog worker server deploy --server-id my-server
```