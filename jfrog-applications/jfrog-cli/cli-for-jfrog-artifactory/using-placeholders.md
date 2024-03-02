# Using Placeholders

## Overview
The JFrog CLI offers enormous flexibility in how you **download, upload**, **copy**, or **move** files through the use of wildcard or regular expressions with placeholders.

Any wildcard enclosed in parentheses in the source path can be matched with a corresponding placeholder in the target path to determine the name of the artifact once uploaded.

## Examples
### Example 1: Upload all files to the target repository

For each .tgz file in the source directory, create a corresponding directory with the same name in the target repository and upload it there. For example, a file named **froggy.tgz** should be uploaded to **my-local-rep/froggy**. **froggy** will be created in a folder in Artifactory).

```
jf rt u "(*).tgz" my-local-repo/{1}/ --recursive=false
```

### Example 2: Upload all files sharing the same prefix to the target repository

Upload all files whose name begins with "frog" to folder **frogfiles** in the target repository, but append its name with the text "-up". For example, a file called **froggy.tgz** should be renamed **froggy.tgz-up**.

```
jf u "(frog*)" my-local-repo/frogfiles/{1}-up --recursive=false
```

### Example 3: Upload all files to corresponding directories according to extension type

Upload all files in the current directory to the **my-local-repo** repository and place them in directories that match their file extensions.

```
jf rt u "(*).(*)" my-local-repo/{2}/{1}.{2} --recursive=false
```

### Example 4: Copy all zip files to target repository and append with an extension

Copy all zip files under /rabbit in the **source-frog-repo** repository into the same path in the **target-frog-repo** repository and append the copied files' names with ".cp".

```
jf rt cp "source-frog-repo/rabbit/(*.zip)" target-frog-repo/rabbit/{1}.cp
```
