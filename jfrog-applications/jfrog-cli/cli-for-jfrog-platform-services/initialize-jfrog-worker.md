# Initialize JFrog Worker

### Overview

This command is used to initialize a new JFrog worker.

|                        |                                                                                                                                                 |
| ---------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| Command name           | worker init                                                                                                                                     |
| Abbreviation           | worker i                                                                                                                                        |
| **Command options:**   |                                                                                                                                                 |
| `--server-id`          | \[Optional] Server ID configured using the config command.                                                                                      |
| `--timeout-ms`         | \[Default: 5000] The request timeout in milliseconds.                                                                                           |
| `--force`              | \[Default: false] Whether to overwrite existing files.                                                                                          |
| `--no-test`            | \[Default: false] Whether to skip test generation.                                                                                              |
| `--application`        | \[Optional] The application that provides the event. If omitted the service will try to guess it and raise an error if no application is found. |
| `--project-key`        | \[Optional] The key of the project that the worker should belong to.                                                                            |
| **Command arguments:** |                                                                                                                                                 |
| action                 | The name of the action to init (eg: BEFORE\_DOWNLOAD). To have the list of all available actions use `jf worker list-event`.                    |
| worker-name            | The name of the worker.                                                                                                                         |

This command generates the following files:

* `manifest.json` – Contains the Worker specification, including its name, code location, secrets, and other data useful to the Worker.
* `package.json` – Describes the development dependencies of the Worker. This file is not used when executing your Worker in the runtime.
* `worker.ts` – The Worker source code, populated with sample code for the event.
* `worker.spec.ts` - The source code for the Worker's unit tests.
* `tsconfig.json` - The TypeScript configuration file.
* `types.ts` - A file containing the event's specific types that can be used in the Worker code.

#### Example: Initialize a `BEFORE_DOWNLOAD` Worker

This example initializes a new `BEFORE_DOWNLOAD` Worker named `my-worker`.

```
jf worker init BEFORE_DOWNLOAD my-worker
```
