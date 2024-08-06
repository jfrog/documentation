# Using File Specs
## Overview

To achieve complex file manipulations you may require several CLI commands. For example, you may need to upload several different sets of files to different repositories. To simplify the implementation of these complex manipulations, you can apply JFrog CLI **download**, **upload**, **move**, **copy** and **delete** commands with JFrog Artifactory using **--spec** option to replace the inline command arguments and options. Similarly, you can **create and update release bundles** by providing the `--spec` command option. Each command uses an array of file specifications in JSON format with a corresponding schema as described in the sections below. Note that if any of these commands are issued using both inline options and the file specs, then the inline options override their counterparts specified in the file specs.

## File Spec Schemas

### Copy and Move Commands Spec Schema

The file spec schema for the copy and move commands is as follows:

```
{
  "files": [
  {
    "pattern" or "aql": "[Mandatory]",
    "target": "[Mandatory]",
    "props": "[Optional]",
    "excludeProps": "[Optional]",
    "recursive": "[Optional, Default: 'true']",
    "flat": "[Optional, Default: 'false']",
    "exclusions": [
      "Optional",
      "Applicable only when 'pattern' is specified"
    ],
    "archiveEntries": "[Optional]",
    "build": "[Optional]",
    "bundle": "[Optional]",
    "validateSymlinks": "[Optional]",
    "sortBy": "[Optional]",
    "sortOrder": "[Optional, Default: 'asc']",
    "limit": "[Optional],
    "offset": [Optional] }
  ]
}
```

### Download Command Spec Schema

The file spec schema for the download command is as follows:

```
{
  "files": [
  {
    "pattern" or "aql": "[Mandatory]",
    "target": "[Optional]",
    "props": "[Optional]",
    "excludeProps": "[Optional]",
    "recursive": "[Optional, Default: 'true']",
    "flat": "[Optional, Default: 'false']",
    "exclusions": [
      "Optional",
      "Applicable only when 'pattern' is specified"
    ],
    "archiveEntries": "[Optional]",
    "build": "[Optional]",
    "bundle": "[Optional]",
    "sortBy": "[Optional]",
    "sortOrder": "[Optional, Default: 'asc']",
    "limit": [Optional],
    "offset": [Optional] }
  ]
}
```

### Create and Update Release Bundle V1 Commands Spec Schema

The file spec schema for the create and update release bundle v1 commands is as follows:

```
{
"files": [
  {
    "pattern" or "aql": "[Mandatory]",
    "pathMapping": "[Optional, Applicable only when 'aql' is specified]",
    "target": "[Optional]",
    "props": "[Optional]",
    "targetProps": "[Optional]",
    "excludeProps": "[Optional]",
    "recursive": "[Optional, Default: 'true']",
    "flat": "[Optional, Default: 'false']",
    "exclusions": [
      "Optional",
      "Applicable only when 'pattern' is specified"
    ],
    "archiveEntries": "[Optional]",
    "build": "[Optional]",
    "bundle": "[Optional]",
    "sortBy": "[Optional]",
    "sortOrder": "[Optional, Default: 'asc']",
    "limit": [Optional],
    "offset": [Optional] }
  ]
}
```

### Upload Command Spec Schema

The file spec schema for the upload command is as follows:

```
{
  "files": [
  {
    "pattern": "[Mandatory]",
    "target": "[Mandatory]",
    "targetProps": "[Optional]",
    "recursive": "[Optional, Default: 'true']",
    "flat": "[Optional, Default: 'true']",
    "regexp": "[Optional, Default: 'false']",
    "ant": "[Optional, Default: 'false']",
    "archive": "[Optional, Must be: 'zip']",
    "exclusions": [
      "Optional"
    ] }
  ]
}
```

### Search, Set-Props and Delete Commands Spec Schema

The file spec schema for the search and delete commands are as follows:

```
{
  "files": [
  {
    "pattern" or "aql": "[Mandatory]",
    "props": "[Optional]",
    "excludeProps": "[Optional]",
    "recursive": "[Optional, Default: 'true']",
    "exclusions": [
      "Optional",
      "Applicable only when 'pattern' is specified"
    ],
    "archiveEntries": "[Optional]",
    "build": "[Optional]",
    "bundle": "[Optional]",
    "sortBy": "[Optional]",
    "sortOrder": "[Optional, Default: 'asc']",
    "limit": [Optional],
    "offset": [Optional] }
  ]
}
```

### Examples

The following examples can help you get started using File Specs.

#### Example 1

Download all files located under the **all-my-frogs** directory in the **my-local-repo** repository to the **froggy** directory.

```
{
  "files": [
  {
    "pattern": "my-local-repo/all-my-frogs/",
    "target": "froggy/" }
  ]
}
```

#### Example 2

Download all files located under the **all-my-frogs** directory in the **my-local-repo** repository to the **froggy** directory. Download only files which are artifacts of build number 5 of build **my-build** .

```
{
  "files": [
    {
      "pattern": "my-local-repo/all-my-frogs/",
      "target": "froggy/",
      "build": "my-build/5"
    }
  ]
}
```

#### Example 3

Download all files retrieved by the AQL query to the **froggy** directory.

```
{
  "files": [
    {
      "aql": {
        "items.find": {
          "repo": "my-local-repo",
          "$or": [
            {
              "$and": [
                {
                  "path": {
                    "$match": "."
                  },
                  "name": {
                    "$match": "a1.in"
                  }
                }
              ]
            },
            {
              "$and": [
                {
                  "path": {
                    "$match": "*"
                  },
                  "name": {
                    "$match": "a1.in"
                  }
                }
              ]
            }
          ]
        }
      },
      "target": "froggy/"
    }
  ]
}
```

#### Example 4

1. All zip files located under the **resources** directory to the **zip** folder, under the **all-my-frogs** repository. AND
2. All TGZ files located under the **resources** directory to the \*\*tgz folder, under the **all-my-frogs** repository.
3. Tag all zip files with type = zip and status = ready.
4. Tag all tgz files with type = tgz and status = ready.

```
{
  "files": [
    {
      "pattern": "resources/*.zip",
      "target": "all-my-frogs/zip/",
      "props": "type=zip;status=ready"
    },
    {
      "pattern": "resources/*.tgz",
      "target": "all-my-frogs/tgz/",
      "props": "type=tgz;status=ready"
    }
  ]
}
```

#### Example 5

Upload all zip files located under the **resources** directory to the **zip** folder, under the **all-my-frogs** repository.

```
{
  "files": [
    {
      "pattern": "resources/*.zip",
      "target": "all-my-frogs/zip/"
    }
  ]
}
```

#### Example 6

Package all files located (including subdirectories) under the **resources** directory into a zip archive named **archive.zip** , and upload it into the root of the **all-my-frogs** repository.

```
{
  "files": [
    {
      "pattern": "resources/",
      "archive": "zip",
      "target": "all-my-frogs/"
    }
  ]
}
```

#### Example 7

Download all files located under the **all-my-frogs** directory in the **my-local-repo** repository **except** for files with .txt extension and all files inside the **all-my-frogs** directory with the props. prefix.\`

Notice that the exclude patterns do not include the repository.

```
{
    "files": [
     {
        "pattern": "my-local-repo/all-my-frogs/",
        "exclusions": ["*.txt","all-my-frog/props.*"]
     }
    ]
}
```

#### Example 8

Download The latest file uploaded to the **all-my-frogs** directory in the **my-local-repo** repository.

```
{
    "files": [
     {
        "pattern": "my-local-repo/all-my-frogs/",
        "target": "all-my-frogs/files/",
        "sortBy": ["created"],
        "sortOrder": "desc",
        "limit": 1
     }
    ]
}
```

#### Example 9

Search for the three largest files located under the **all-my-frogs** directory in the **my-local-repo** repository. If there are files with the same size, sort them "internally" by creation date.

```
{
    "files": [
     {
        "pattern": "my-local-repo/all-my-frogs/",
        "sortBy": ["size","created"],
        "sortOrder": "desc",
        "limit": 3
     }
    ]
}
```

#### Example 10

Download The second-latest file uploaded to the **all-my-frogs** directory in the **my-local-repo** repository.

```
{
    "files": [
     {
        "pattern": "my-local-repo/all-my-frogs/",
        "target": "all-my-frogs/files/",
        "sortBy": ["created"],
        "sortOrder": "desc",
        "limit": 1,
        "offset": 1
     }
    ]
}
```

#### Example 11

This example shows how to [delete artifacts in artifactory under specified path based on how old they are](https://stackoverflow.com/questions/58328701/delete-artifacts-in-artifactory-under-specified-path-based-on-how-old-they-are).

The following File Spec finds all the folders which match the following criteria:

1. They are under the my-repo repository.
2. They are inside a folder with a name that matches abc-\*-xyz and is located at the root of the repository.
3. Their name matches ver\*
4. They were created more than 7 days ago.

```
{
  "files": [
    {
      "aql": {
        "items.find": {
          "repo": "myrepo",
          "path": {"$match":"abc-*-xyz"},
          "name": {"$match":"ver*"},
          "type": "folder",
          "$or": [
            {
              "$and": [
                {
                  "created": { "$before":"7d" }
                }
              ]
            }
          ]
        }
      }
    }
  ]
}
```

#### Example 12

This example uses [Using Placeholders](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/cli-for-jfrog-artifactory#using-placeholders). For each .tgz file in the source directory, create a corresponding directory with the same name in the target repository and upload it there. For example, a file named froggy.tgz should be uploaded to my-local-rep/froggy. (froggy will be created a folder in Artifactory).

```
{
    "files": [
      {
        "pattern": "(*).tgz",
        "target": "my-local-repo/{1}/",
      }
    ]
}
```

#### Example 13

This examples uses [Using Placeholders](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/cli-for-jfrog-artifactory#using-placeholders). Upload all files whose name begins with "frog" to folder frogfiles in the target repository, but append its name with the text "-up". For example, a file called froggy.tgz should be renamed froggy.tgz-up.

```
{
    "files": [
      {
        "pattern": "(frog*)",
        "target": "my-local-repo/frogfiles/{1}-up",
        "recursive": "false"
      }
    ]
}
```

#### Example 14

The following two examples lead to the exact same outcome.\
The first one uses [Using Placeholders](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/cli-for-jfrog-artifactory#using-placeholders), while the second one does not. Both examples download all files from the generic-local repository to be under the ֿֿmy/local/path/ local file-system path, while maintaining the original Artifactory folder hierarchy. Notice the different flat values in the two examples.

```
{
    "files": [
      {
        "pattern": "generic-local/{*}",
        "target": "my/local/path/{1}",
        "flat": "true"
      }
    ]
}

{
    "files": [
      {
        "pattern": "generic-local/",
        "target": "my/local/path/",
        "flat": "false"
      }
    ]
}
```

#### Example 15

This example creates a release bundle v1 and applies "pathMapping" to the artifact paths after distributing the release bundle v1.

All occurrences of the "a1.in" file are fetched and mapped to the "froggy" repository at the edges.

1. Fetch all artifacts retrieved by the AQL query.
2. Create the release bundle v1 with the artifacts and apply the path mappings at the edges after distribution.

The "pathMapping" option is provided, allowing users to control the destination of the release bundle artifacts at the edges.

To learn more, visit the [Create Release Bundle v1 Version documentation](https://jfrog.com/help/r/jfrog-rest-apis/create-release-bundle-v1-version).
```json
{
  "files": [
    {
      "aql": {
        "items.find": {
          "repo": "my-local-repo",
          "$and": [
            {
              "name": {
                "$match": "a1.in"
              }
            },
            {
              "$or": [
                {
                  "path": {
                    "$match": "."
                  }
                },
                {
                  "path": {
                    "$match": "*"
                  }
                }
              ]
            }
          ]
        }
      },
      "pathMapping": {
        "input": "my-local-repo/(.*)",
        "output": "froggy/$1"
      }
    }
  ]
}
```

### Schema Validation

[JSON schemas](https://json-schema.org/) allow you to annotate and validate JSON files. The JFrog File Spec schema is available in the [JSON Schema Store](https://www.schemastore.org/json/) catalog and in the following link: [https://github.com/jfrog/jfrog-cli/blob/v2/schema/filespec-schema.json](https://github.com/jfrog/jfrog-cli/blob/v2/schema/filespec-schema.json).

#### Using Jetbrains IDEs (Intellij IDEA, Webstorm, Goland, etc...)?

The File Spec schema is automatically applied to the following file patterns:

\*\*/filespecs/\*.json

\*filespec\*.json

\*.filespec

#### Using Visual Studio Code?

To apply the File Spec schema validation, install the [JFrog VS-Code extension](https://marketplace.visualstudio.com/items?itemName=JFrog.jfrog-vscode-extension).

Alternatively, copy the following to your settings.json file:

**settings.json**

```
"json.schemas": [
  {
    "fileMatch": ["**/filespecs/*.json", "\*filespec\*.json", "*.filespec"],
    "url": "https://raw.githubusercontent.com/jfrog/jfrog-cli/v2/schema/filespec-schema.json"
  }
]
```
