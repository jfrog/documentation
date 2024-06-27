# Initialize a Worker

## Overview

This command is used to initialize a new platform worker.

|                        |                                                                                                                                                                                     |
|------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command name           | worker init                                                                                                                                                                         |
| Abbreviation           | worker i                                                                                                                                                                            |
| **Command options:**   |                                                                                                                                                                                     |
| `--force`              | \[Default: false] Whether to overwrite existing files.                                                                                                                              |
| `--no-test`            | \[Default: false] Whether to skip tests generation.                                                                                                                                 |
| **Command arguments:** |                                                                                                                                                                                     |
| action                 | The action that will trigger the worker (BEFORE_DOWNLOAD, AFTER_DOWNLOAD, BEFORE_UPLOAD, AFTER_CREATE, AFTER_BUILD_INFO_SAVE, AFTER_MOVE, GENERIC_EVENT, BEFORE_CREATE_TOKEN, ...). |
| worker-name            | The name of the worker                                                                                                                                                              |

## Example

Initialize a new BEFORE_DOWNLOAD worker named `my-worker`.

```
jfrog worker init BEFORE_DOWNLOAD my-worker
```