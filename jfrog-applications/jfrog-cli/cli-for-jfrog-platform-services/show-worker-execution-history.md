# Show Worker Execution History

### Overview

Display in a json format the execution history of a specific worker.

|                        |                                                                                                                |
| ---------------------- | -------------------------------------------------------------------------------------------------------------- |
| Command name           | worker execution-history                                                                                       |
| Abbreviation           | worker exec-hist, worker eh                                                                                    |
| **Command options:**   |                                                                                                                |
| `--server-id`          | \[Optional] Server ID configured using the config command.                                                     |
| `--timeout-ms`         | \[Default: 5000] The request timeout in milliseconds.                                                          |
| `--project-key`        | \[Optional] List events available to a specific project.                                                       |
| `--with-test-runs`     | \[Default: false] Whether to include test-runs entries.                                                        |
| **Command arguments:** |                                                                                                                |
| worker-key             | \[Optional] The worker key. If not provided it will be read from the `manifest.json` in the current directory. |

### Example

Retrieves the execution history of a worker named `my-worker` including test runs.

```
jf worker execution-history --with-test-runs my-worker
```
