# Release Lifecycle Management

## Overview

This page describes how to use JFrog CLI with [Release Lifecycle Management](https://jfrog.com/help/r/jfrog-artifactory-documentation/jfrog-release-lifecycle-management-solution).

***

**Note**

> Release Lifecycle Management is available since [Artifactory 7.63.2](https://jfrog.com/help/r/jfrog-release-information/artifactory-7.63.2-cloud).

***

## Syntax

When used with JFrog Release Lifecycle Management, JFrog CLI uses the following syntax:

```
$ jf command-name global-options command-options arguments
```

## Create a Release Bundle v2

The **create** command allows creating a Release Bundle v2 using [file specs](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory/using-file-specs). The file spec may be of one of the following creation sources:

*   Published build info:

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

    `<build-number>` is optional; the latest build will be used if empty. `includeDeps` is optional, false by default. `project` is optional; the default project will be used if empty.
*   Existing Release Bundles:

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

    `project` is optional; the default project will be used if empty.
*   A pattern of artifacts in Artifactory:

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
*   AQL query:

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

  * Specified package (one or more):
    ```
      {
      "files": [
              {
                 "package": "catalina",
                 "version":"1.0.0",
                 "type": "maven",
                 "repoKey": "catalina-dev-maven-local"
             },
     ]
     }
    ```
  
    * A Release Bundle created from multiple sources:
      The Release Bundle can include any number of builds, artifacts (using a pattern), packages, and Release Bundles.
      However, it can include only one AQL query:
    ```
       {
        "files": [
              {
                "build": "Commons-Build/1.0.0",
                "includeDeps":"true",
                "project": "default"
              },
               {
                "bundle": "rb1/1.0.0",
                "project": "default"
              },
              {
                "bundle": "rb4/1.0.0",
                "project": "default"
              },
               {
                "pattern": "catalina-dev-maven-local/*.jar"
              },
               {
                "package": "commons",
                "version":"1.0.1",
                "type": "maven",
                "repoKey": "commons-dev-maven-local"
              },
              {
                "package": "catalina",
                "version":"1.0.0",
                "type": "maven",
                "repoKey": "catalina-dev-maven-local"
              },
              {
                "aql": {
                 "items.find": {
                  "repo":{"$eq":"catalina-dev-maven-local"},
                  "$and":[
                    {"name": {"$match":"*1.0.0.pom"}}
                  ]
                 }
              }
           }
        ]
      }
    ```


### Command Params

|                    |                                                                                                                                                                                                                                                        |
| ------------------ |--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command-name       | release-bundle-create                                                                                                                                                                                                                                  |
| Abbreviation       | rbc                                                                                                                                                                                                                                                    |
| **Command arguments:** |                                                                                                                                                                                                                                                        |
| release bundle name | Name of the newly created Release Bundle.                                                                                                                                                                                                              |
| release bundle version | Version of the newly created Release Bundle.                                                                                                                                                                                                           |
| **Command options:** |                                                                                                                                                                                                                                                        |
| `--project`        | <p>[Optional]<br>Project key associated with the created Release Bundle version.</p>                                                                                                                                                                   |
| `--server-id`      | <p>[Optional]<br>Platform Server ID configured using the 'jf config' command.</p>                                                                                                                                                                      |
| `--signing-key`    | <p>[Optional]<br>The GPG/RSA key-pair name defined in Artifactory. The <code>signing-key</code> can also be configured as an environment variable. If no key is specified, Artifactory uses a default key.</p>                                         |
| `--spec`           | <p>[Optional]<br>Path to a File Spec. If you do not define the spec, you must include the <code>build-name</code> and <code>build-number</code> as environment variables,  flags, or a combination of both (flags override environment variables).</p> |
| `--spec-vars`      | <p>[Optional]<br>List of semicolon-separated(;) variables in the form of "key1=value1;key2=value2;..." (wrapped by quotes) to be replaced in the File Spec. In the File Spec, the variables should be used as follows: ${key1}.</p>                    |
| `--build-name`     | <p>[Optional]<br>The name of the build from which to create the Release Bundle.</p>                                                                                                                                                                    |
| `--build-number`   | <p>[Optional]<br>The number of the build from which to create the Release Bundle.</p>                                                                                                                                                                  |
| `--source-type-release-bundles`| <p>[Optional]<br>One or more Release Bundles to include in the new Release Bundle in the form of "name=[rb-name], version=[rb-version];..." .</p>                                                                                                      |
| `--source-type-builds` | <p>[Optional]<br>One or more builds to include in the new Release Bundle in the form of "name=[build-name], id=[build-id], include-deps=[true/false];...".</p>                                                                                                                                                          |
| `--sync`           | <p>[Default: true]<br>Set to false to run asynchronously.</p>                                                                                                                                                                                          |


### Examples

#### Example 1

Create a Release Bundle using file spec variables.

```
jf rbc --spec=/path/to/spec.json --spec-vars="key1=value1" --signing-key=myKeyPair myApp 1.0.0
```

#### Example 2

Create a Release Bundle synchronously, in project "project0".

```
jf rbc --spec=/path/to/spec.json --signing-key=myKeyPair --sync=true --project=project0 myApp 1.0.0
```

#### Example 3

Create a Release Bundle using build name and build number variables.

```
jf rbc --build-name=Common-builds --build-number=1.0.0 myApp 1.0.0
```

#### Example 4

Create a Release Bundle from multiple builds

```
jf rbc rb3 1.0.0 --source-type-builds "name=Commons-Build, id=1.0.0, include-deps=true; name=Commons-Build, id=1.0.1"
```

#### Example 5

Create a Release Bundle from multiple existing Release Bundles.

```
jf rbc rb3 1.0.0 --project catalina --source-type-release-bundles "name=rb1, version=1.0.0; name=rb2, version=1.0.0"
```

#### Example 6

Create a Release Bundle from existing builds and Release Bundles.

```
jf rbc rb3 1.0.0 --source-type-builds "name=Commons-Build, id=1.0.0, include-deps=true; name=Commons-Build, id=1.0.1" 
--source-type-release-bundles "name=rb1, version=1.0.0; name=rb2, version=1.0.0"
```


## Promote a Release Bundle v2

This command allows promoting a Release Bundle to a target environment.

### Commands Params

|                        |                                                                                                                                                                                                                                                                                                                      |
|------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command-name           | release-bundle-promote                                                                                                                                                                                                                                                                                               |
| Abbreviation           | rbp                                                                                                                                                                                                                                                                                                                  |
| **Command arguments:** |                                                                                                                                                                                                                                                                                                                      |
| release bundle name    | Name of the Release Bundle to promote.                                                                                                                                                                                                                                                                               |
| release bundle version | Version of the Release Bundle to promote.                                                                                                                                                                                                                                                                            |
| environment            | Name of the target environment for the promotion.                                                                                                                                                                                                                                                                    |
| **Command options:**   |                                                                                                                                                                                                                                                                                                                      |
| `--input-repos`        | <p>[Optional]<br>A list of semicolon-separated(;) repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. If one or more repositories are specifically included, all other repositories are excluded.</p> |
| `--exclude-repos`      | <p>[Optional]<br>A list of semicolon-separated(;) repositories to exclude from the promotion.</p>                                                                                                                                                                                                                    |
| `--project`            | <p>[Optional]<br>Project key associated with the Release Bundle version.</p>                                                                                                                                                                                                                                         |
| `--server-id`          | <p>[Optional]<br>Platform Server ID configured using the 'jf config' command.</p>                                                                                                                                                                                                                                    |
| `--signing-key`        | <p>[Mandatory]<br>The GPG/RSA key-pair name given in Artifactory.</p>                                                                                                                                                                                                                                                |
| `--sync`               | <p>[Default: true]<br>Set to false to run asynchronously.</p>                                                                                                                                                                                                                                                        |
| `--promotion-type`    | <p>[Default: copy]<br>Specifies the promotion type. (Valid values: move / copy) .</p>                                                                                                                                                                                                                                |
 
### Examples

#### Example 1

Promote a Release Bundle named "myApp" version "1.0.0" to environment "PROD". Use signing key pair "myKeyPair".

```
jf rbp --signing-key=myKeyPair myApp 1.0.0 PROD
```

#### Example 2

Promote a Release Bundle synchronously to environment "PROD". The Release Bundle is named "myApp", version "1.0.0", of project "project0". Use signing key pair "myKeyPair".

```
jf rbp --signing-key=myKeyPair --project=project0 --sync=true myApp 1.0.0 PROD
```

#### Example 3

Promote a Release Bundle while including certain repositories.

```
jf rbp --signing-key=myKeyPair --include-repos="generic-local;my-repo" myApp 1.0.0 PROD
```

#### Example 4

Promote a Release Bundle while excluding certain repositories.

```
jf rbp --signing-key=myKeyPair --exclude-repos="generic-local;my-repo" myApp 1.0.0 PROD
```

#### Example 4

Promote a Release Bundle, using promotion type flag.

```
jf rbp --signing-key=myKeyPair --promotion-type="move" myApp 1.0.0 PROD
```

# Annotate a Release Bundle v2

This command allows you to add a single tag to a Release Bundle v2 version, and/or set or delete one or more properties.

### Commands Params

|                        |                                                                                                                                                                                                                             |
|------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command-name           | release-bundle-annotate                                                                                                                                                                                                     |
| Abbreviation           | rba                                                                                                                                                                                                                         |
| **Command arguments:** |                                                                                                                                                                                                                             |
| release bundle name    | Name of the Release Bundle to annotate.                                                                                                                                                                                     |
| `--tag`                | <p>[Optional]<br>The tag is a single free-text value limited to 128 characters, beginning and ending with an alphanumeric character ([a-z0-9A-Z]), with dashes (-), underscores (_), dots (.),and alphanumerics between.<p> |
| `--properties`         | <p>[Optional]<br>Key-value pairs separated by a semicolon (;). Keys are limited to 255 characters. Values are limited to 2400 characters.</p>                                                                               |
| `--del-prop`           | <p>[Optional]<br>Removes a key and all its associated values.</p>                                                                                                                                                           |
| `--recursive`          | <p>[Default: true]<br>Set to false(0) to don't run recursively.</p>                                                                                                                                                         |
                                                                                                                      

### Examples

#### Example 1

Add or modify a tag or property.

```
jf rba mybundle 1.0.0 --tag=release --properties "environment=production;buildNumber=1234"
```

#### Example 2

Whenever you use the --tag command option, the value you define replaces the current value.

```
jf rba mybundle 1.0.0 --tag=rejected
```
In the example above, the tag that was defined previously (release) is replaced with the new tag rejected.

#### Example 3

Whenever you use the --properties command option with an existing key, the values that you define replace the current values.

```
jf rba mybundle 1.0.0 --properties "environment=DEV,PROD,QA"
```
In the example above, the value for environment that was defined previously (production) is replaced by the values DEV, PROD, and QA.

#### Remove Tags and Properties

#### Example 1

To remove the tag, set it to empty.

```
jf rba mybundle 1.0.0 --tag=""
```

#### Example 2

To remove the values from an existing key without removing the key, leave the value empty.

```
jf rba mybundle 1.0.0 --properties "build=''"
```
In the example above, all values defined for the build key are removed but the key is retained.

#### Example 3

To remove a key and its associated values, use the --del-prop command option.

```
jf rba mybundle 1.0.0 --del-prop "environment"
```
In the example above, the environment key and all its associated values are removed.


## Distribute a Release Bundle v2

This command distributes a Release Bundle to an Edge node.

|                        |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| ---------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name           | release-bundle-distribute                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| Abbreviation           | rbd                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| **Command arguments:** |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| release bundle name    | Name of the Release Bundle to distribute.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| release bundle version | Version of the Release Bundle to distribute.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| **Command options:**   |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| `--city`               | <p>[Default: *]<br>Wildcard filter for site city name.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| `--country-codes`      | <p>[Default: *]<br>semicolon-separated(;) list of wildcard filters for site country codes.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| `--create-repo`        | <p>[Default: false]<br>Set to true to create the repository on the edge if it does not exist.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| `--dist-rules`         | <p>[Optional]<br>Path to a file, which includes the Distribution Rules in a JSON format. See the "Distribution Rules Structure" bellow.</p>                                                                                                                                                                                                                                                                                                                                                                                                       |
| `--dry-run`            | <p>[Default: false]<br>Set to true to disable communication with JFrog Distribution.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| `--mapping-pattern`    | <p>[Optional]<br>Specify along with 'mapping-target' to distribute artifacts to a different path on the Edge node. You can use wildcards to specify multiple artifacts.</p>                                                                                                                                                                                                                                                                                                                                                                       |
| `--mapping-target`     | <p>[Optional]<br>The target path for distributed artifacts on the edge node. If not specified, the artifacts will have the same path and name on the edge node, as on the source Artifactory server. For flexibility in specifying the distribution path, you can include <a href="https://www.jfrog.com/confluence/display/CLI/CLI+for+JFrog+Artifactory#CLIforJFrogArtifactory-UsingPlaceholders">placeholders</a> in the form of {1}, {2} which are replaced by corresponding tokens in the pattern path that are enclosed in parenthesis.</p> |
| `--max-wait-minutes`   | <p>[Default: 60]<br>Max minutes to wait for sync distribution.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| `--project`            | <p>[Optional]<br>Project key associated with the Release Bundle version.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| `--server-id`          | <p>[Optional]<br>Platform Server ID configured using the 'jf config' command.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| `--site`               | <p>[Default: *]<br>Wildcard filter for site name.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| `--sync`               | <p>[Default: true]<br>Set to false to run asynchronously.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |

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

Distribute the Release Bundle named myApp with version 1.0.0. Use the distribution rules defined in the specified file.

```
jf rbd --dist-rules=/path/to/dist-rules.json myApp 1.0.0
```

#### Example 2

Distribute the Release Bundle named myApp with version 1.0.0 using the default distribution rules. Map files under the `source` directory to be placed under the `target` directory.

```
jf rbd --dist-rules=/path/to/dist-rules.json --mapping-pattern="(*)/source/(*)" --mapping-target="{1}/target/{2}" myApp 1.0.0
```

#### Example 3

Synchronously distribute a Release Bundle associated with project "proj"

```
jf rbd --dist-rules=/path/to/dist-rules.json --sync --project="proj" myApp 1.0.0
```

## Delete a Release Bundle v2 locally

This command allows deleting all Release Bundle promotions to a specified environment or deleting a Release Bundle locally altogether. Deleting locally means distributions of the Release Bundle will not be deleted.

|                        |                                                                                                                                        |
| ---------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name           | release-bundle-delete-local                                                                                                            |
| Abbreviation           | rbdell                                                                                                                                 |
| **Command arguments:** |                                                                                                                                        |
| release bundle name    | Name of the Release Bundle to distribute.                                                                                              |
| release bundle version | Version of the Release Bundle to distribute.                                                                                           |
| environment            | If provided, all promotions to this environment are deleted. Otherwise, the Release Bundle is deleted locally with all its promotions. |
| **Command options:**   |                                                                                                                                        |
| `--project`            | <p>[Optional]<br>Project key associated with the Release Bundle version.</p>                                                           |
| `--quiet`              | <p>[Default: $CI]<br>Set to true to skip the delete confirmation message.</p>                                                          |
| `--server-id`          | <p>[Optional]<br>Platform Server ID configured using the 'jf config' command.</p>                                                      |
| `--sync`               | <p>[Default: true]<br>Set to false to run asynchronously.</p>                                                                          |

### Examples

#### Example 1

Locally delete the Release Bundle named myApp with version 1.0.0.

```
jf rbdell myApp 1.0.0
```

#### Example 2

Locally delete the Release Bundle named myApp with version 1.0.0. Run the command synchronously and skip the confirmation message.

```
jf rbdell --quiet --sync myApp 1.0.0
```

#### Example 3

Delete all promotions of the specified Release Bundle version to environment "PROD".

```
jf rbdell myApp 1.0.0 PROD
```

## Delete a Release Bundle v2 remotely

This command will delete distributions of a Release Bundle from a distribution target, such as an Edge node.

|                        |                                                                                                                                             |
| ---------------------- | ------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name           | release-bundle-delete-remote                                                                                                                |
| Abbreviation           | rbdelr                                                                                                                                      |
| **Command arguments:** |                                                                                                                                             |
| release bundle name    | Name of the Release Bundle to delete.                                                                                                       |
| release bundle version | Version of the Release Bundle to delete.                                                                                                    |
| **Command options:**   |                                                                                                                                             |
| `--city`               | <p>[Default: *]<br>Wildcard filter for site city name.</p>                                                                                  |
| `--country-codes`      | <p>[Default: *]<br>semicolon-separated(;) list of wildcard filters for site country codes.</p>                                              |
| `--dist-rules`         | <p>[Optional]<br>Path to a file, which includes the Distribution Rules in a JSON format. See the "Distribution Rules Structure" bellow.</p> |
| `--dry-run`            | <p>[Default: false]<br>Set to true to disable communication with JFrog Distribution.</p>                                                    |
| `--max-wait-minutes`   | <p>[Default: 60]<br>Max minutes to wait for sync distribution.</p>                                                                          |
| `--project`            | <p>[Optional]<br>Project key associated with the Release Bundle version.</p>                                                                |
| `--quiet`              | <p>[Default: $CI]<br>Set to true to skip the delete confirmation message.</p>                                                               |
| `--server-id`          | <p>[Optional]<br>Platform Server ID configured using the 'jf config' command.</p>                                                           |
| `--site`               | <p>[Default: *]<br>Wildcard filter for site name.</p>                                                                                       |
| `--sync`               | <p>[Default: true]<br>Set to false to run asynchronously.</p>                                                                               |

### Examples

#### Example 1

Delete the distributions of version 1.0.0 of the Release Bundle named myApp from Edge nodes matching the provided distribution rules defined in the specified file.

```
jf rbd --dist-rules=/path/to/dist-rules.json myApp 1.0.0
```

#### Example 2

Delete the distributions of the Release Bundle associated with project "proj" from the provided Edge nodes. Run the command synchronously and skip the confirmation message.

```
jf rbd --dist-rules=/path/to/dist-rules.json --project="proj" --quiet --sync myApp 1.0.0
```

## Export a Release Bundle v2 archive

Release Lifecycle Management supports distributing your Release Bundles to remote Edge nodes within an air-gapped environment. This use case is mainly intended for organizations that have two or more JFrog instances that have no network connection between them.

The following command allows exporting a Release Bundle as an archive to the filesystem that can be transferred to a different instance in an air-gapped environment.

|                        |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| ---------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name           | release-bundle-export                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| Abbreviation           | rbe                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| **Command arguments:** |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| release bundle name    | Name of the Release Bundle to export.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| release bundle version | Version of the Release Bundle to export.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| target pattern         | <p>The argument is optional and specifies the local file system target path.</p><p></p><p>If the target path ends with a slash, the path is assumed to be a directory. For example, if you specify the target as "repo-name/a/b/", then "b" is assumed to be a directory into which files should be downloaded.</p><p></p><p>If there is no terminal slash, the target path is assumed to be a file to which the downloaded file should be renamed. For example, if you specify the target as "a/b", the downloaded file is renamed to "b".</p> |
| **Command options:**   |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| `--project`            | <p>[Optional]<br>Project key associated with the Release Bundle version.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| `--server-id`          | <p>[Optional]<br>Platform Server ID configured using the 'jf config' command.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| mapping-pattern        | <p>[Optional]<br>Specify a list of input regex mapping pairs that define where the queried artifact is located and where it should be placed after it is imported. Use this option if the path on the target is different than the source path.</p>                                                                                                                                                                                                                                                                                             |
| mapping-target         | <p>[Optional]<br>Specify a list of output regex mapping pairs that define where the queried artifact is located and where it should be placed after it is imported. Use this option if the path on the target is different than the source path.</p>                                                                                                                                                                                                                                                                                            |
| split-count            | <p>[Optional]<br>The maximum number of parts that can be concurrently uploaded per file during a multi-part upload. Set to 0 to disable multi-part upload.</p>                                                                                                                                                                                                                                                                                                                                                                                  |
| min-split              | <p>[Optional]<br>Minimum file size in KB to split into ranges when downloading. Set to -1 for no splits</p>                                                                                                                                                                                                                                                                                                                                                                                                                                     |

#### Example

Export version 1.0.0 of the Release Bundle named "myApp":

```
jf rbe myApp 1.0.0
```

#### Example

Download the file to a specific location:

```
jf rbe myApp 1.0.0 /user/mybundle/
```

## Import a Release Bundle v2 archive

You can import a Release Bundle archive from the exported zip file.

Please note this functionality only works on Edge nodes within an air-gapped environment.

|                        |                                                                                   |
| ---------------------- | --------------------------------------------------------------------------------- |
| Command-name           | release-bundle-import                                                             |
| Abbreviation           | rbi                                                                               |
| **Command arguments:** |                                                                                   |
| path to archive        | Path to the Release Bundle archive on the filesystem.                             |
| **Command options:**   |                                                                                   |
| `--project`            | <p>[Optional]<br>Project key associated with the Release Bundle version.</p>      |
| `--server-id`          | <p>[Optional]<br>Platform Server ID configured using the 'jf config' command.</p> |

#### Example

Import version 1.0.0 of a Release Bundle named "myExportedApp":

```
jf rbi ./myExportedApp.zip
```



## Download Release Bundle v2 content

Use the following command to download the contents of a Release Bundle v2 version:

```
jf rt dl --bundle [release-bundle-name]/[release-bundle-version]
```

For more information, see [Downloading Files](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/binaries-management-with-jfrog-artifactory/generic-files#downloading-files).&#x20;
