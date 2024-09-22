# List registered workers

## Overview

List workers saved on your Artifactory instance. The default output is a CSV format with columns `name`,`action`,`description`,`enabled`.

|                        |                                                                                                                                                                                     |
|------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command name           | worker list                                                                                                                                                                         |
| Abbreviation           | worker ls                                                                                                                                                                           |
| **Command options:**   |                                                                                                                                                                                     |
| `--server-id`          | \[Optional] Server ID configured using the config command.                                                                                                                          |
| `--json`               | \[Default: false] Whether to use JSON instead of CSV as output.                                                                                                                     |
| `--timeout-ms`         | \[Default: 5000] The request timeout in milliseconds.                                                                                                                               |


# Example

List all workers registered in a platform named `my-platform` with detailed data.

```
jfrog worker list --server-id my-platform --json
```