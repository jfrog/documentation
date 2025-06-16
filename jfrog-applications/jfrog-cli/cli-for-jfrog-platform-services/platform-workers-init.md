# Initialize a Worker

## Overview

This command is used to initialize a new JFrog worker.

|                        |                         |
|---------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command name           | worker init |
| Abbreviation           | worker i |
| **Command options:**   ||
| `--server-id`          | \[Optional] Server ID configured using the config command.|
| `--timeout-ms`         | \[Default: 5000] The request timeout in milliseconds.|
| `--force`              | \[Default: false] Whether to overwrite existing files. |
| `--no-test`            | \[Default: false] Whether to skip test generation.|
| `--application`            | \[Optional] The application that provides the event. If omitted the service will try to guess it and raise an error if no application is found.|
| `--project-key`            | \[Optional] The key of the project that the worker should belong to.|
| **Command arguments:** ||
| action                 | The name of the action to init (eg: BEFORE_DOWNLOAD). To have the list of all available actions use `jf worker list-event`. |
| worker-name            | The name of the worker.|

This command will generate the following files:

- `manifest.json` – The worker specification which includes its name, the code location, the secrets and all the data that is useful to the worker.
- `package.json` – The file describing the development dependencies of the worker, this will not be used when executing your worker in our runtime.
- `worker.ts` – The worker code source, here it will be our sample code for the event.
- `worker.spec.ts` - The source code of the worker unit tests.
- `tsconfig.json` - The Typescript configuration file.
- `types.ts` - A file containing the event's specific types that can be used in the worker code.

## Example

Initialize a new BEFORE_DOWNLOAD worker named `my-worker`.

```
jf worker init BEFORE_DOWNLOAD my-worker
```
