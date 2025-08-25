# Cleaning Up Unreferenced Files from a Git LFS Repository

## Overview

Use this command to clean up a Git LFS repository. This command deletes all files from a Git LFS repository in Artifactory that are no longer referenced in a corresponding Git repository.

## Commands Params

|                        |                                                                                                                                              |
| ---------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| Command name           | rt git-lfs-clean                                                                                                                             |
| Abbreviation           | rt glc                                                                                                                                       |
| **Command options:**   |                                                                                                                                              |
| `--refs`               | \[Default: refs/remotes/\*] List of Git references in the form of "ref1,ref2,..." which should be preserved.                                 |
| `--repo`               | \[Optional] Local Git LFS repository in Artifactory which should be cleaned. If omitted, the repository is detected from the Git repository. |
| `--quiet`              | \[Default: false] Set to true to skip the delete confirmation message.                                                                       |
| `--dry-run`            | \[Default: false] If true, cleanup is only simulated. No files are actually deleted.                                                         |
| **Command arguments:** | If no arguments are passed in, the command assumes the .git repository is located at current directory.                                      |
| path to .git           | Path to the directory which includes the .git directory.                                                                                     |

### Examples

#### Example 1: Clean Up Using the Current Directory

This example cleans up Git LFS files from Artifactory, using the configuration from the `.git` directory located in the current directory.

```
jf rt glc
```

#### Example 2: Clean Up Using a Specific Path

This example cleans up Git LFS files from Artifactory, using the configuration from the `.git` directory located inside the `path/to/git/config` directory.

```
jf rt glc path/to/git/config
```
