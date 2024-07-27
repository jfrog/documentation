# Test a worker

## Overview

Dry run a worker. The worker needs to be initialized before running this command.

|                        |                                                                                                                                                         |
|------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command name           | worker dry-run                                                                                                                                          |
| Abbreviation           | worker dr                                                                                                                                               |
| **Command options:**   |                                                                                                                                                         |
| `--server-id`          | \[Optional] Server ID configured using the config command.                                                                                              |
| `--timeout-ms`         | \[Default: 5000] The request timeout in milliseconds.                                                                                                   |
| `--no-secrets`         | \[Default: false] Do not use registered secrets.                                                                                                        |
| **Command arguments:** |                                                                                                                                                         |
| json-payload           | The json payload expected by the worker. Use `-` to read the payload from standard input. Use `@<file-path>` to read from a file located at <file-path>. |

## Example

Dry run a worker initialized in the current directory, with a payload located in a file named `payload.json` from the same directory.

```
jfrog worker dry-run @payload.json
```