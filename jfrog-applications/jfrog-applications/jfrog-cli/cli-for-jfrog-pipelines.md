# CLI for JFrog Pipelines

### Overview

This page describes how to use JFrog CLI with JFrog Pipelines.

Read more about JFrog CLI [here](https://jfrog-external.fluidtopics.net/r/help/JFrog-CLI/JFrog-CLI).

### Syntax

```
$ jf pl command-name arguments command-options
```

Where:

|                 |                                                                                                 |
| --------------- | ----------------------------------------------------------------------------------------------- |
| command-name    | The command to execute. Note that you can use either the full command name or its abbreviation. |
| command-options | A set of options corresponding to the command                                                   |
| arguments       | A set of arguments corresponding to the command                                                 |

***

### Commands

The following sections describe the commands available in the JFrog CLI for use with JFrog Pipelines.

#### Checking Pipelines Version

Check the version of Pipelines installation.

|                   |                                                                                                                                                    |
| ----------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command name      | pl version                                                                                                                                         |
| Abbreviation      | v                                                                                                                                                  |
| Command arguments | The command accepts no arguments.                                                                                                                  |
| Command options   |                                                                                                                                                    |
| --server-id       | <p>[Optional]<br><br>Specify Pipelines server for which the version is to be fetched. If not specified, the default configured server is used.</p> |

**Example**

```
jf pl version --server-id repo21
```

#### Getting Run Status

Get the status of the run for the specified pipeline.

|                   |                                                                                                                                |
| ----------------- | ------------------------------------------------------------------------------------------------------------------------------ |
| Command name      | pl status                                                                                                                      |
| Abbreviation      | s                                                                                                                              |
| Command arguments | The command accepts no arguments.                                                                                              |
|                   |                                                                                                                                |
| Command options   |                                                                                                                                |
| --pipeline-name   | <p>[Optional]<br><br>Name of the pipeline.</p>                                                                                 |
| --branch          | <p>[Optional]<br><br>Name of the multi branch.</p>                                                                             |
| --monitor         | <p>[Optional]<br><br>Continuous monitoring until pipeline reaches end state.<br><br>Default time is 1 hour and 30 minutes.</p> |
| --server-id       | <p>[Optional]<br><br>Name of the server.</p>                                                                                   |
| --single-branch   | <p>[Optional]<br><br>To be used when using a single branch.</p>                                                                |

**Examples**

**Single Branch**

```
jf pl status --pipeline-name myPipeline --single-branch --server-id repo21 --monitor --single-branch
```

**Multi Branch**

```
jf pl status --pipeline-name myPipeline --branch main/jobs --server-id repo21 --monitor
```

#### Triggering a Pipeline

Trigger a pipeline run.

|                   |                                                                                                                          |
| ----------------- | ------------------------------------------------------------------------------------------------------------------------ |
| Command name      | pl trigger                                                                                                               |
| Abbreviation      | t                                                                                                                        |
| Command options   |                                                                                                                          |
| --server-id       | <p>[Optional]<br><br>Name of the server.</p>                                                                             |
| --single-branch   | <p>[Optional]<br><br>To be used when using a single branch. When used, <code>branch_name</code> argument is ignored.</p> |
| Command arguments |                                                                                                                          |
| pipeline\_name    | Name of the pipeline to be triggered.                                                                                    |
| branch\_name      | Name of the multi branch.                                                                                                |

**Examples**

**Single Branch**

```
jf pl trigger myPipeline main/jobs --single-branch --server-id repo21
```

**Multi Branch**

```
jf pl trigger myPipeline main/jobs --server-id repo21
```

####

#### Syncing Pipeline Source

Perform a sync to load the latest pipeline source.

|                 |                                                                                           |
| --------------- | ----------------------------------------------------------------------------------------- |
| Command name    | pl sync                                                                                   |
| Abbreviation    | sy                                                                                        |
| Command options |                                                                                           |
| --server-id     | <p>[Optional]<br><br>Name of the server.</p>                                              |
| --repository    | <p>[Optional]<br><br>Full name of the repository where the pipeline source is stored.</p> |
| --branch        | <p>[Optional]<br><br>Name of branch that has the pipeline source.</p>                     |

**Example**

```
jf pl sync --repository jfrog/artifactory --branch main/jobs --server-id repo21
```

#### Sync Status of Pipeline Source

Get the current status of the pipeline source sync.

|                   |                                                                  |
| ----------------- | ---------------------------------------------------------------- |
| Command name      | pl sync-status                                                   |
| Abbreviation      | ss                                                               |
| Command options   |                                                                  |
| --server-id       | <p>[Optional]<br><br>Name of the server.</p>                     |
| Command arguments |                                                                  |
| full\_repo\_name  | Full name of the repository where the pipeline source is stored. |
| branch\_name      | Name of the branch.                                              |

**Example**

```
jf pl sync-status jfrog/artifactory main/jobs --server-id repo21
```
