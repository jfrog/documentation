# Release Lifecycle Management

## Overview

This page describes how to use JFrog CLI with [Release Lifecycle Management](https://jfrog.com/help/r/jfrog-artifactory-documentation/jfrog-release-lifecycle-management-solution).

***

**Note**

> Release Lifecycle Management is only available since [Artifactory 7.63.2](https://jfrog.com/help/r/jfrog-release-information/artifactory-7.63.2-cloud).

***

## Syntax

When used with JFrog Release Lifecycle Management, JFrog CLI uses the following syntax:

```
$ jf command-name global-options command-options arguments
```

## Creating a release bundle

The **create** command allows creating a release bundle using [file specs](using-file-specs.md). 
The file spec may be of one of the following creation sources:

1. Published build infos.
    ```
    {
      "files": [
        {
          "build": "<build-name>/<build-number>",
          "includeDeps": "[true/false]",
          "project": "<project-key>"
        },
        ...
      ]
    }
    ```
    `<build-number>` is optional, latest build will be used if empty.
    `includeDeps` is optional, false by default.
    `project` is optional, default project will be used if empty.

2. Existing release bundles.
    ```
    {
      "files": [
        {
          "bundle": "<bundle-name>/<bundle-version>",
          "project": "<project-key>"
        },
        ...
      ]
    }
    ```
   `project` is optional, default project will be used if empty.

3. A pattern of artifacts in Artifactory.
    ```
    {
      "files": [
        {
          "pattern": "repo/path/*",
          "exclusions": ["excluded",...],
          "props": "key1=value1;key2=value2;key3=value3",
          "excludeArtifacts": "key1=value1;key2=value2;key3=value3",
          "recursive": "[true/false]"
        },
        ...
      ]
    }
    ```
   Only `pattern` is mandatory. `recursive` is true by default.

4. AQL query.
    ```
    {
      "files": [
        {
          "aql": {
            "items.find": {
              "repo": "<repo>",
              "path": "<path>",
              "name": "<file>"
            }
          }
        }
      ]
    }
    ```
   Only a single AQL query may be provided.
    
    
### Commands Params

|                        |                                                                                                                                                                                                                  |
|------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command-name           | release-bundle-create                                                                                                                                                                                            |
| Abbreviation           | rbc                                                                                                                                                                                                              |
| Command options        |                                                                                                                                                                                                                  |
| --project              | <p>[Optional]<br>Project key associated with the created Release Bundle version.</p>                                                                                                                         |
| --server-id            | <p>[Optional]<br>Platform server ID configured using the config command.</p>                                                                                                                                 |
| --signing-key          | <p>[Mandatory]<br>The GPG/RSA key-pair name given in Artifactory.</p>                                                                                                                                        |
| --spec                 | <p>[Optional]<br>Path to a File Spec.</p>                                                                                                                                                                    |
| --spec-vars            | <p>[Optional]<br>List of variables in the form of "key1=value1;key2=value2;..." (wrapped by quotes) to be replaced in the File Spec. In the File Spec, the variables should be used as follows: ${key1}.</p> |
| --sync                 | <p>[Optional]<br>Set to true to run synchronously.</p>                                                                                                                                                 |
| Command arguments      |                                                                                                                                                                                                                  |
| release bundle name    | Name of the newly created Release Bundle.                                                                                                                                                                        |
| release bundle version | Version of the newly created Release Bundle.                                                                                                                                                                     |

### Examples
#### Example 1

Create a release bundle with name "myApp" and version "1.0.0", using signing key pair "myKeyPair". The release bundle will include artifacts corresponding to the creation source in the provided file spec.

```
jf rbc --spec=/path/to/spec.json --signing-key=myKeyPair myApp 1.0.0
```

#### Example 2

Create a release bundle synchronously, in project "project0".

```
jf rbc --spec=/path/to/spec.json --signing-key=myKeyPair --sync=true --project=project0 myApp 1.0.0
```

#### Example 3

Create a release bundle using file spec variables.

```
jf rbc --spec=/path/to/spec.json --spec-vars="key1=value1" --signing-key=myKeyPair myApp 1.0.0
```

## Promoting a release bundle

This command allows promoting a release bundle to a target environment.

### Commands Params

|                        |                                                                                                                                                                                                                                                                                                                       |
|------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command-name           | release-bundle-promote                                                                                                                                                                                                                                                                                                |
| Abbreviation           | rbp                                                                                                                                                                                                                                                                                                                   |
| Command options        |                                                                                                                                                                                                                                                                                                                       |
| --input-repos          | <p>[Optional]<br>A list of semicolon-separated repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. If one or more repositories are specifically included, all other repositories are excluded.</p> |
| --exclude-repos        | <p>[Optional]<br>A list of semicolon-separated repositories to exclude from the promotion.</p>                                                                                                                                                                                                                    |
| --project              | <p>[Optional]<br>Project key associated with the Release Bundle version.</p>                                                                                                                                                                                                                                      |
| --server-id            | <p>[Optional]<br>Platform server ID configured using the config command.</p>                                                                                                                                                                                                                                      |
| --signing-key          | <p>[Mandatory]<br>The GPG/RSA key-pair name given in Artifactory.</p>                                                                                                                                                                                                                                             |
| --sync                 | <p>[Optional]<br>Set to true to run synchronously.</p>                                                                                                                                                                                                                                                      |
| Command arguments      |                                                                                                                                                                                                                                                                                                                       |
| release bundle name    | Name of the Release Bundle to promote.                                                                                                                                                                                                                                                                                |
| release bundle version | Version of the Release Bundle to promote.                                                                                                                                                                                                                                                                             |
| environment            | Name of the target environment for the promotion.                                                                                                                                                                                                                                                                     |

### Examples
#### Example 1

Promote a release bundle named "myApp" version "1.0.0" to environment "PROD". Use signing key pair "myKeyPair".

```
jf rbp --signing-key=myKeyPair myApp 1.0.0 PROD
```

#### Example 2

Promote a release bundle synchronously to environment "PROD". The release bundle is named "myApp", version "1.0.0", of project "project0". Use signing key pair "myKeyPair".

```
jf rbp --signing-key=myKeyPair --project=project0 --sync=true myApp 1.0.0 PROD
```

#### Example 3

Promote a release bundle while including certain repositories.

```
jf rbp --signing-key=myKeyPair --include-repos="generic-local;my-repo" myApp 1.0.0 PROD
```

#### Example 4

Promote a release bundle while excluding certain repositories.

```
jf rbp --signing-key=myKeyPair --exclude-repos="generic-local;my-repo" myApp 1.0.0 PROD
```

## Distributing a release bundle

This command distributes a release bundle to an edge node.

|                        |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
|------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command-name           | release-bundle-distribute                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| Abbreviation           | rbd                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| Command options        |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| --city                 | <p>[Default: *]<br>Wildcard filter for site city name.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| --country-codes        | <p>[Default: *]<br>Semicolon-separated list of wildcard filters for site country codes.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| --create-repo          | <p>[Optional]<br>Set to true to create the repository on the edge if it does not exist.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| --dist-rules           | <p>[Optional]<br>Path to a file, which includes the Distribution Rules in a JSON format. See the "Distribution Rules Structure" bellow.</p>                                                                                                                                                                                                                                                                                                                                                                                            |
| --dry-run              | <p>[Optional]<br>Set to true to disable communication with JFrog Distribution.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| --mapping-pattern      | <p>[Optional]<br>Specify along with 'mapping-target' to distribute artifacts to a different path on the edge node. You can use wildcards to specify multiple artifacts.</p>                                                                                                                                                                                                                                                                                                                                                            |
| --mapping-target       | <p>[Optional]<br>The target path for distributed artifacts on the edge node. If not specified, the artifacts will have the same path and name on the edge node, as on the source Artifactory server. For flexibility in specifying the distribution path, you can include [placeholders](https://www.jfrog.com/confluence/display/CLI/CLI+for+JFrog+Artifactory#CLIforJFrogArtifactory-UsingPlaceholders) in the form of {1}, {2} which are replaced by corresponding tokens in the pattern path that are enclosed in parenthesis.</p> |
| --max-wait-minutes     | <p>[Default: 60]<br>Max minutes to wait for sync distribution.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| --project              | <p>[Optional]<br>Project key associated with the Release Bundle version.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| --server-id            | <p>[Optional]<br>Platform server ID configured using the config command.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| --site                 | <p>[Default: *]<br>Wildcard filter for site name.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| --sync                 | <p>[Optional]<br>Set to true to run synchronously.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| Command arguments      |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| release bundle name    | Name of the release bundle to distribute.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| release bundle version | Version of the release bundle to distribute.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |

**Distribution Rules Structure**
   ```json
   {
    "distribution_rules": [
       {
          "site_name": "DC-1",
          "city_name": "New-York",
          "country_codes": ["1"]
       },
       {
          "site_name": "DC-2",
          "city_name": "Tel-Aviv",
          "country_codes": ["972"]
       }
    ]
   }
   ```
The Distribution Rules format also supports wildcards. For example:
   ```json
   {
    "distribution_rules": [
       {
          "site_name": "",
          "city_name": "",
          "country_codes": ["*"]
       }
    ]
   }
   ```

### Examples
#### Example 1
Distribute the release bundle named myApp with version 1.0.0. Use the distribution rules defined in the specified file.

	jf rbd --dist-rules=/path/to/dist-rules.json myApp 1.0.0

#### Example 2

Distribute the release bundle named myApp with version 1.0.0 using the default distribution rules.
Map files under the 'source' directory to be placed under the 'target' directory.

	jf rbd --dist-rules=/path/to/dist-rules.json --mapping-pattern="(*)/source/(*)" --mapping-target="{1}/target/{2}" myApp 1.0.0

#### Example 3
Synchronously distribute a release bundle associated with project "proj" 

	jf rbd --dist-rules=/path/to/dist-rules.json --sync --project="proj" myApp 1.0.0


## Deleting release bundle locally
This command allows deleting all release bundle promotions to an environment or deleting a release bundle locally altogether.
Deleting locally means distributions of the release bundle will not be deleted.

|                        |                                                                                                                                        |
|------------------------|----------------------------------------------------------------------------------------------------------------------------------------|
| Command-name           | release-bundle-delete-local                                                                                                            |
| Abbreviation           | rbdell                                                                                                                                 |
| Command options        |                                                                                                                                        |
| --project              | <p>[Optional]<br>Project key associated with the Release Bundle version.</p>                                                       |
| --quiet                | <p>[Default: $CI]<br>Set to true to skip the delete confirmation message.</p>                                                      |
| --server-id            | <p>[Optional]<br>Platform server ID configured using the config command.</p>                                                       |
| --sync                 | <p>[Optional]<br>Set to true to run synchronously.</p>                                                                       |
| Command arguments      |                                                                                                                                        |
| release bundle name    | Name of the release bundle to distribute.                                                                                              |
| release bundle version | Version of the release bundle to distribute.                                                                                           |
| environment            | If provided, all promotions to this environment are deleted. Otherwise, the release bundle is deleted locally with all its promotions. |

### Examples
#### Example 1
Locally delete the release bundle named myApp with version 1.0.0 altogether.

	jf rbdell myApp 1.0.0

#### Example 2
Delete a release bundle locally altogether. Run the command synchronously and skip the confirmation message.

	jf rbdell --quiet --sync myApp 1.0.0

#### Example 3
Delete all promotions of the release bundle to environment "PROD".

	jf rbdell myApp 1.0.0 PROD


## Deleting release bundle remotely
This command will delete distributions of a release bundle from edge nodes.

|                        |                                                                                                                                                 |
|------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------|
| Command-name           | release-bundle-delete-remote                                                                                                                    |
| Abbreviation           | rbdelr                                                                                                                                          |
| Command options        |                                                                                                                                                 |
| --city                 | <p>[Default: *]<br>Wildcard filter for site city name.</p>                                                                                  |
| --country-codes        | <p>[Default: *]<br>Semicolon-separated list of wildcard filters for site country codes.</p>                                                 |
| --dist-rules           | <p>[Optional]<br>Path to a file, which includes the Distribution Rules in a JSON format. See the "Distribution Rules Structure" bellow.</p> |
| --dry-run              | <p>[Optional]<br>Set to true to disable communication with JFrog Distribution.</p>                                                    |
| --max-wait-minutes     | <p>[Default: 60]<br>Max minutes to wait for sync distribution.</p>                                                                          |
| --project              | <p>[Optional]<br>Project key associated with the Release Bundle version.</p>                                                                |
| --quiet                | <p>[Default: $CI]<br>Set to true to skip the delete confirmation message.</p>                                                               |
| --server-id            | <p>[Optional]<br>Platform server ID configured using the config command.</p>                                                                |
| --site                 | <p>[Default: *]<br>Wildcard filter for site name.</p>                                                                                       |
| --sync                 | <p>[Optional]<br>Set to true to run synchronously.</p>                                                                                |
| Command arguments      |                                                                                                                                                 |
| release bundle name    | Name of the release bundle to distribute.                                                                                                       |
| release bundle version | Version of the release bundle to distribute.                                                                                                    |

### Examples
#### Example 1
Delete the distributions of release bundle named myApp with version 1.0.0 from edge nodes matching the provided distribution rules defined in the specified file.

	jf rbd --dist-rules=/path/to/dist-rules.json myApp 1.0.0

#### Example 2
Delete the distributions of the release bundle associated with project "proj" from the provided edge nodes. Run the command synchronously and skip the confirmation message.

	jf rbd --dist-rules=/path/to/dist-rules.json --project="proj" --quiet --sync myApp 1.0.0
