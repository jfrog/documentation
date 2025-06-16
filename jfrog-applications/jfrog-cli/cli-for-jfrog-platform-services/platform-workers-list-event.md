# List available events

## Overview

This command list all the available events on the platform.

|                        |                           |
|------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command name           | worker list-event  |
| Abbreviation           | worker le |
| **Command options:**   |  |
| `--server-id`          | \[Optional] Server ID configured using the config command. |
| `--timeout-ms`         | \[Default: 5000] The request timeout in milliseconds. |
| `--project-key`            | \[Optional] List events available to a specific project.|

## Example

List event supported by a server identified by `my-server`.

```
jf worker list-event --server-id my-server
```