# Cleaning Up Unreferenced Files from a Git LFS Repository
## Overview

This command is used to clean up files from a Git LFS repository. This deletes all files from a Git LFS repository, which are no longer referenced in a corresponding Git repository.

## Commands Params

|                   |                                                                                                                                              |
| ----------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| Command name      | rt git-lfs-clean                                                                                                                             |
| Abbreviation      | rt glc                                                                                                                                       |
| Command options   |                                                                                                                                              |
| --refs            | \[Default: refs/remotes/\*] List of Git references in the form of "ref1,ref2,..." which should be preserved.                                 |
| --repo            | \[Optional] Local Git LFS repository in Artifactory which should be cleaned. If omitted, the repository is detected from the Git repository. |
| --quiet           | \[Default: false] Set to true to skip the delete confirmation message.                                                                       |
| --dry-run         | \[Default: false] If true, cleanup is only simulated. No files are actually deleted.                                                         |
| Command arguments | If no arguments are passed in, the command assumes the .git repository is located at current directory.                                      |
| path to .git      | Path to the directory which includes the .git directory.                                                                                     |

## Examples
### Example 1

Cleans up Git LFS files from Artifactory, using the configuration in the .git directory located at the current directory.

```
 jf rt glc
```

### Example 2

Cleans up Git LFS files from Artifactory, using the configuration in the .git directory located inside the path/to/git/config directory.

```
jf rt glc path/to/git/config
```
