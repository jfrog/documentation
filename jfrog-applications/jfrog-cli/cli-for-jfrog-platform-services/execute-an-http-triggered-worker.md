# Execute an HTTP-Triggered Worker

### Overview

Execute an HTTP-triggered worker.

|                        |                                                                                                                                               |
| ---------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| Command name           | worker execute                                                                                                                                |
| Abbreviation           | worker exec, worker e                                                                                                                         |
| **Command options:**   |                                                                                                                                               |
| `--server-id`          | \[Optional] Server ID configured using the config command.                                                                                    |
| `--timeout-ms`         | \[Default: 5000] The request timeout in milliseconds.                                                                                         |
| `--project-key`        | \[Optional] The key of the project that the worker belongs to.                                                                                |
| **Command arguments:** |                                                                                                                                               |
| worker-key             | The worker key. If not provided it will be read from the `manifest.json` in the current directory.                                            |
| json-payload           | The json payload expected by the worker. Use `-` to read the payload from standard input. Use `@<file-path>` to read from a file located at . |

### Example

Execute an HTTP-triggered worker initialized in the current directory, with a payload located in a file named `payload.json` from the same directory.

```
jf worker execute @payload.json
```

Execute an HTTP-triggered worker with a payload from the standard input.

```
jf worker execute - <<EOF
{
  “a”: “key”,
  “an-integer”: 14
}
EOF
```

Execute an HTTP-triggered worker by providing the payload as an argument.

```
jf worker execute ‘{“my”:”payload”}’
```
