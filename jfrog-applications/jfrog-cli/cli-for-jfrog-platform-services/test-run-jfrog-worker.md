# Test-Run JFrog Worker

### Overview

Use this command to test-run a Worker. You must initialize the Worker before running this command. The command executes the Worker with its local content, so you can use it to test the Worker's execution before pushing local changes to the server.

|                        |                                                                                                                                               |
| ---------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| **Command name**       | `worker test-run`                                                                                                                             |
| **Abbreviation**       | worker dry-run, worker dr, worker tr                                                                                                          |
| **Command options:**   |                                                                                                                                               |
| `--server-id`          | \[Optional] Server ID configured using the config command.                                                                                    |
| `--timeout-ms`         | \[Default: 5000] The request timeout in milliseconds.                                                                                         |
| `--no-secrets`         | \[Default: false] Do not use registered secrets.                                                                                              |
| **Command arguments:** |                                                                                                                                               |
| json-payload           | The json payload expected by the worker. Use `-` to read the payload from standard input. Use `@<file-path>` to read from a file located at . |

#### Example

This example test-runs a Worker that has been initialized in the current directory, using a payload from a file named `payload.json` in the same directory.

```
jf worker dry-run @payload.json
```
