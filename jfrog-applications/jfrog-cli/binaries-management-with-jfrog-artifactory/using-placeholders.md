# Using Placeholders

## Overview

JFrog CLI provides flexibility in how you download, upload, copy, or move files through the use of wildcard or regular expressions with placeholders.

Any wildcard expression enclosed in parentheses in the source path can be matched with a corresponding placeholder (e.g., `{1}`) in the target path. This allows you to dynamically determine the destination path and filename for an artifact.

### Examples

#### Example 1: Upload Files into Dynamically Named Directories

This example uploads each `.tgz` file from the source into a new directory in the target repository that matches the file's base name. For example, the file `froggy.tgz` is uploaded to `my-local-repo/froggy/`, which creates the `froggy` folder in Artifactory.

```
jf rt u "(*).tgz" my-local-repo/{1}/ --recursive=false
```

#### Example 2: Upload Files with Modified Names

This example uploads all files with names that begin with "frog" to the `frogfiles` folder, appending the suffix `-up` to each filename. For example, a file named `froggy.tgz` is renamed to `froggy.tgz-up`.

```
jf rt u "(frog*)" my-local-repo/frogfiles/{1}-up --recursive=false
```

#### Example 3: Organize Uploaded Files by Extension

This example uploads all files from the current directory into the `my-local-repo` repository, placing them into subdirectories named after their file extensions.

```
jf rt u "(*).(*)" my-local-repo/{2}/{1}.{2} --recursive=false
```

#### Example 4: Copy Files and Append a Suffix

This example copies all `.zip` files from the `rabbit/` directory in the `source-frog-repo` to the same path in the `target-frog-repo` and appends `.cp` to each copied filename.

```
jf rt cp "source-frog-repo/rabbit/(*.zip)" target-frog-repo/rabbit/{1}.cp
```
