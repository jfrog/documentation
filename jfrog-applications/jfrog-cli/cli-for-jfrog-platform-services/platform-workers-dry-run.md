# Test Run a worker

## Overview

Test-run a worker. The worker needs to be initialized before running this command.
The command will execute the worker with its local content, so it can be used to test the worker execution before pushing the local changes to the server.

|                        |                       |
|------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command name           | worker test-run       |
| Abbreviation           | worker dry-run, worker dr, worker tr             |
| **Command options:**   |                       |
| `--server-id`          | \[Optional] Server ID configured using the config command.|
| `--timeout-ms`         | \[Default: 5000] The request timeout in milliseconds. |
| `--no-secrets`         | \[Default: false] Do not use registered secrets. |
| **Command arguments:** |                       |
| json-payload           | The json payload expected by the worker. Use `-` to read the payload from standard input. Use `@<file-path>` to read from a file located at <file-path>. |

## Example

Test-run a worker initialized in the current directory, with a payload located in a file named `payload.json` from the same directory.

```
jf worker dry-run @payload.json
```
