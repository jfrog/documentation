# CLI for JFrog Release Lifecycle Management

## Overview

This page describes how to use JFrog CLI with [JFrog Release Lifecycle Management](https://jfrog.com/help/r/jfrog-artifactory-documentation/jfrog-release-lifecycle-management-solution).

Read more about JFrog CLI [here](https://jfrog.com/help/r/jfrog-cli).

***

**Note**

> JFrog Release Lifecycle Management is only available since [Artifactory 7.63.2](https://jfrog.com/help/r/jfrog-release-information/artifactory-7.63.2-cloud).

***

### Syntax

When used with JFrog Release Lifecycle Management, JFrog CLI uses the following syntax:

```
$ jf command-name global-options command-options arguments
```

### Commands

The following sections describe the commands available in JFrog CLI for use with JFrog Release Lifecycle Management.

### Creating a release bundle from builds or from existing release bundles

This commands allows creating a release bundle from one of two sources:

1.  Published build infos. To use, provide the `--builds` option, which accepts a path to a JSON file of the following syntax:

    ```json
    {
      "builds": [
        {
          "name": "build name",
          "number": "build number",
          "project": "project"
        }
      ]
    }
    ```

    `number` is optional, latest build will be used if empty.

    `project` is optional, default project will be used if empty.
2.  Existing release bundles. To use, provide the `--release-bundles` option, which accepts a path to a JSON file of the following syntax:

    ```json
    {
      "releaseBundles": [
        {
          "name": "release bundle name",
          "version": "release bundle version",
          "project": "project"
        }
      ]
    }
    ```

    `project` is optional, default project will be used if empty.
    
    
#### Commands Params

|                        |                                                                                                                                          |
| ---------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name           | release-bundle-create                                                                                                                    |
| Abbreviation           | rbc                                                                                                                                      |
| Command options        |                                                                                                                                          |
| --builds               | <p>[Optional]<br><br>Path to a JSON file containing information of the source builds from which to create a release bundle.</p>          |
| --project              | <p>[Optional]<br><br>Project key associated with the Release Bundle version.</p>                                                         |
| --release-bundles      | <p>[Optional]<br><br>Path to a JSON file containing information of the source release bundles from which to create a release bundle.</p> |
| --server-id            | <p>[Optional]<br><br>Platform server ID configured using the config command.</p>                                                         |
| --signing-key          | <p>[Mandatory]<br><br>The GPG/RSA key-pair name given in Artifactory.</p>                                                                |
| --sync                 | <p>[Default: false]<br><br>Set to true to run synchronously.</p>                                                                         |
| Command arguments      |                                                                                                                                          |
| release bundle name    | Name of the newly created Release Bundle.                                                                                                |
| release bundle version | Version of the newly created Release Bundle.                                                                                             |

#### Example 1

Create a release bundle with name "myApp" and version "1.0.0", with signing key pair "myKeyPair". The release bundle will include artifacts of the builds that were provided in the builds spec.

```
jf rbc --builds=/path/to/builds-spec.json --signing-key=myKeyPair myApp 1.0.0
```

#### Example 2

Create a release bundle with name "myApp" and version "1.0.0", with signing key pair "myKeyPair". The release bundle will include artifacts of the release bundles that were provided in the release bundles spec.

```
jf rbc --release-bundles=/path/to/release-bundles-spec.json --signing-key=myKeyPair myApp 1.0.0
```

#### Example 3

Create a release bundle synchronously with name "myApp" and version "1.0.0", in project "project0", with signing key pair "myKeyPair". The release bundle will include artifacts of the release bundles that were provided in the release bundles spec.

```
jf rbc --release-bundles=/path/to/release-bundles-spec.json --signing-key=myKeyPair --sync=true --project=project0 myApp 1.0.0
```

### Promoting a release bundle

This commands allows promoting a release bundle to a target environment.

#### Commands Params

|                        |                                                                                                                                                                                                                             |
| ---------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name           | release-bundle-promote                                                                                                                                                                                                      |
| Abbreviation           | rbp                                                                                                                                                                                                                         |
| Command options        |                                                                                                                                                                                                                             |
| --overwrite            | <p>[Default: false]<br><br>Set to true to replace artifacts with the same name but a different checksum if such already exist at the promotion targets. By default, the promotion is stopped in a case of such conflict</p> |
| --project              | <p>[Optional]<br><br>Project key associated with the Release Bundle version.</p>                                                                                                                                            |
| --server-id            | <p>[Optional]<br><br>Platform server ID configured using the config command.</p>                                                                                                                                            |
| --signing-key          | <p>[Mandatory]<br><br>The GPG/RSA key-pair name given in Artifactory.</p>                                                                                                                                                   |
| --sync                 | <p>[Default: false]<br><br>Set to true to run synchronously.</p>                                                                                                                                                            |
| Command arguments      |                                                                                                                                                                                                                             |
| release bundle name    | Name of the Release Bundle to promote.                                                                                                                                                                                      |
| release bundle version | Version of the Release Bundle to promote.                                                                                                                                                                                   |
| environment            | Name of the target environment for the promotion.                                                                                                                                                                           |


#### Example 1

Promote a release bundle named "myApp" version "1.0.0" to environment "PROD". Use signing key pair "myKeyPair".

```
jf rbp --signing-key=myKeyPair myApp 1.0.0 PROD
```

#### Example 2

Promote a release bundle synchronously to environment "PROD". The release bundle is named "myApp", version "1.0.0", of project "project0". Use signing key pair "myKeyPair" and overwrite at conflict.

```
jf rbp --signing-key=myKeyPair --project=project0 --overwrite=true --sync=true myApp 1.0.0 PROD
```

### Distributing a release bundle

This command distributes a release bundle to an edge node.

|                        |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
|------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command-name           | release-bundle-distribute                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| Abbreviation           | rbd                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| Command options        |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| --create-repo          | <p>[Default: false]<br><br>Set to true to create the repository on the edge if it does not exist.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| --dry-run              | <p>[Default: false]<br><br>Set to true to disable communication with JFrog Distribution.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| --dist-rules           | <p>[Optional]<br><br>Path to a file, which includes the Distribution Rules in a JSON format. See the "Distribution Rules Structure" bellow.</p>                                                                                                                                                                                                                                                                                                                                                                                            |
| --site                 | <p>[Default: *]<br><br>Wildcard filter for site name.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| --city                 | <p>[Default: *]<br><br>Wildcard filter for site city name.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| --country-codes        | <p>[Default: *]<br><br>Semicolon-separated list of wildcard filters for site country codes.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| --insecure-tls         | <p>[Default: false]<br><br>Set to true to skip TLS certificates verification.</p>                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| --mapping-pattern      | <p>[Optional]<br><br>Specify along with 'mapping-target' to distribute artifacts to a different path on the edge node. You can use wildcards to specify multiple artifacts.</p>                                                                                                                                                                                                                                                                                                                                                            |
| --mapping-target       | <p>[Optional]<br><br>The target path for distributed artifacts on the edge node. If not specified, the artifacts will have the same path and name on the edge node, as on the source Artifactory server. For flexibility in specifying the distribution path, you can include [placeholders](https://www.jfrog.com/confluence/display/CLI/CLI+for+JFrog+Artifactory#CLIforJFrogArtifactory-UsingPlaceholders) in the form of {1}, {2} which are replaced by corresponding tokens in the pattern path that are enclosed in parenthesis.</p> |
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


#### Example 1
Distribute the release bundle named myApp with version 1.0.0. Use the distribution rules defined in the specified file.

	jf rbd --dist-rules=/path/to/dist-rules.json myApp 1.0.0

#### Example 2

Distribute the release bundle named myApp with version 1.0.0 using the default distribution rules.
Map files under the 'source' directory to be placed under the 'target' directory.

	jf rbd --dist-rules=/path/to/dist-rules.json --mapping-pattern="(*)/source/(*)" --mapping-target="{1}/target/{2}" myApp 1.0.0
