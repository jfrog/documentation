# Managing Configuration Entities

JFrog CLI offers a set of commands for managing Artifactory configuration entities.

## Creating Users

This command allows creating a bulk of users. The details of the users are provided in a CSV format file. Here's the file format.&#x20;

```
"username","password","email"
"username1","password1","john@c.com"
"username2","password1","alice@c.com"
```

> **Note**: The first line in the CSV is cells' headers. It is mandatory and is used by the command to map the cell value to the users' details.

The CSV can include additional columns, with different headers, which will be ignored by the command.

### Commands Params

|                        |                                                                                                                                             |
| ---------------------- | ------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name           | rt users-create                                                                                                                             |
| Abbreviation           | rt uc                                                                                                                                       |
| **Command options:**   |                                                                                                                                             |
| `--server-id`          | <p>[Optional]<br>Artifactory Server ID configured using the 'jf config' command.</p>                                                        |
| `--csv`                | <p>[Mandatory]<br>Path to a CSV file with the users' details. The first row of the file should include the name,password,email headers.</p> |
| `--replace`            | <p>[Optional]<br>Set to true if you'd like existing users or groups to be replaced.</p>                                                     |
| `--users-groups`       | <p>[Optional]<br>A list of comma-separated(,) groups for the new users to be associated to.</p>                                             |
| **Command arguments:** | The command accepts no arguments                                                                                                            |

### Example

Create new users according to details defined in the path/to/users.csv file.

```
jf rt users-create --csv path/to/users.csv
```

## Deleting Users

This command allows deleting a bulk of users. The command a list of usernames to delete. The list can be either provided as a comma-seperated argument, or as a CSV file, which includes one column with the usernames. Here's the CSV format.

```
"username"
"username1"
"username2"
"username2"
```

The first line in the CSV is cells' headers. It is mandatory and is used by the command to map the cell value to the users' details.

The CSV can include additional columns, with different headers, which will be ignored by the command.

### Commands Params

|                        |                                                                                                                                                                                |
| ---------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Command-name           | rt users-delete                                                                                                                                                                |
| Abbreviation           | rt udel                                                                                                                                                                        |
| **Command options:**   |                                                                                                                                                                                |
| `--server-id`          | <p>[Optional]<br>Artifactory Server ID configured using the 'jf config' command.</p>                                                                                           |
| `--csv`                | <p>[Optional]<br>Path to a csv file with the usernames to delete. The first row of the file is the reserved for the cells' headers. It must include the "username" header.</p> |
| **Command arguments:** |                                                                                                                                                                                |
| users list             | comma-separated(,) list of usernames to delete. If the --csv command option is used, then this argument becomes optional.                                                      |

### Examples

#### Example 1

Delete the users according to the usernames defined in the path/to/users.csv file.

```
jf rt users-delete --csv path/to/users.csv
```

#### Example 2

Delete the users according to the u1, u2 and u3 usernames.

```
jf rt users-delete "u1,u2,u3"
```

## Creating Groups

This command creates a new users group.

### Commands Params

|                        |                                                                                      |
| ---------------------- | ------------------------------------------------------------------------------------ |
| Command-name           | rt group-create                                                                      |
| Abbreviation           | rt gc                                                                                |
| **Command options:**   |                                                                                      |
| `--server-id`          | <p>[Optional]<br>Artifactory Server ID configured using the 'jf config' command.</p> |
| **Command arguments:** |                                                                                      |
| group name             | The name of the group to create.                                                     |

### Example

Create a new group name **reviewers** .

```
jf rt group-create reviewers
```

## Adding Users to Groups

This command adds a list fo existing users to a group.

### Commands Params

|                        |                                                                                      |
| ---------------------- | ------------------------------------------------------------------------------------ |
| Command-name           | rt group-add-users                                                                   |
| Abbreviation           | rt gau                                                                               |
| **Command options:**   |                                                                                      |
| `--server-id`          | <p>[Optional]<br>Artifactory Server ID configured using the 'jf config' command.</p> |
| **Command arguments:** |                                                                                      |
| group name             | The name of the group to add users to.                                               |
| users list             | Comma-seperated list of usernames to add to the specified group.                     |

### Example

Add to group reviewers the users with the following usernames: u1, u2 and u3.

```
jf rt group-add-users "reviewers" "u1,u2,u3"
```

## Deleting Groups

This command deletes a group.

### Commands Params

|                        |                                                                                      |
| ---------------------- | ------------------------------------------------------------------------------------ |
| Command-name           | rt group-delete                                                                      |
| Abbreviation           | rt gdel                                                                              |
| **Command options:**   |                                                                                      |
| `--server-id`          | <p>[Optional]<br>Artifactory Server ID configured using the 'jf config' command.</p> |
| **Command arguments:** |                                                                                      |
| group name             | The name of the group to delete.                                                     |

### Example

Delete the **reviewers** group.

```
jf rt group-delete "reviewers"
```

## Managing Repositories

JFrog CLI offers a set of commands for managing Artifactory repositories. You can create, update and delete repositories. To make it easier to manage repositories, the commands which create and update the repositories accept a pre-defined configuration template file. This template file can also include variables, which can be later replaced with values, when creating or updating the repositories. The configuration template file is created using the **jf rt repo-template** command.





Creating repository Configuration Template\



The **jf rt repo-template** (or jf rt rpt) command provides interactive prompts for building a JSON configuration template.&#x20;

On running jf rt repo-template \<filename>.json, the CLI prompts as follows:

## Template Type&#x20;

Select the template type. Following are the possible options:

* **create:** Template for creating a new repository
* **update:** Template for updating an existing repository

```
 Select the template type (press Tab for options):
```



For example

```
 Select the template type (press Tab for options): create
```

\


## Repository Key&#x20;

Enter a unique identifier for the repository key. This is the key field in the final JSON.

```
 Insert the repository key >
```

For example

```
 Insert the repository key > npm-local
```

\


**Note:** If you want to reuse the template for creating multiple repositories, use a variable as follows:

```
 Insert the repository key >  ${repo-key-var}
```

## Repository Class

Select the repository class. Following are the possible options:

* **local:** A physical, locally-managed repository into which you can deploy artifacts
* **remote:** A caching proxy for a repository managed at a remote URL
  * **Note:** For remote repositories, you need to enter a remote repository url
* **virtual:** An Aggregation of several repositories with the same package type under a common URL.
* **federated:** A Federation is a collection of repositories of Federated type in different JPDs that are automatically configured for full bi-directional mirroring

```
 Select the repository class (press Tab for options): 
```

For example

```
 Select the repository class (press Tab for options): local
```

## Repository Package Type

Select the repository package type. Following are the possible options:

* alpine
* bower
* chef
* cocoapods
* composer
* conan
* cran
* debian
* docker
* gems
* generic
* gitlfs
* go
* gradle
* helm
* ivy
* maven
* npm
* nuget
* opkg
* pypi
* puppet
* rpm
* sbt
* vagrant
* Yum

**Note:** After selecting the repository package type, you can exit by entering :x or proceed to make advanced configurations.

## Optional Configurations

\
For additional optional configurations to fine-tune the repository's behavior, configure the following:

This table explains the optional keys available for configuring your desired repository in JFrog Artifactory.



| Configuration Key                             | Description                                                                                    | Local       | Remote      | Virtual     | Federated   |
| --------------------------------------------- | ---------------------------------------------------------------------------------------------- | ----------- | ----------- | ----------- | ----------- |
| allowAnyHostAuth                              | Allows sending credentials to any host upon redirection.                                       | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |
| archiveBrowseEnabled                          | Enables viewing archive contents (e.g., ZIPs) in the UI.                                       | ✔️          | <p><br></p> | <p><br></p> | ✔️          |
| artifactoryRequestsCanRetrieveRemoteArtifacts | Allows the virtual repository to resolve artifacts from its remote members.                    | <p><br></p> | <p><br></p> | ✔️          | <p><br></p> |
| assumedOfflinePeriodSecs                      | Sets the time (in seconds) to consider a remote repository offline after a connection failure. | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |
| blackedOut                                    | Temporarily disables the repository, blocking all traffic.                                     | ✔️          | ✔️          | <p><br></p> | ✔️          |
| blockMismatchingMimeTypes                     | Blocks caching of remote files if their MIME type is incorrect.                                | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |
| blockPushingSchema1                           | Blocks pushes from older Docker v1 clients, enforcing the v2 schema.                           | ✔️          | ✔️          | <p><br></p> | ✔️          |
| bypassHeadRequests                            | Skips initial HEAD requests and sends GET requests directly to the remote repository.          | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |
| cdnRedirect                                   | Redirects client download requests to a configured Content Delivery Network (CDN).             | ✔️          | ✔️          | <p><br></p> | ✔️          |
| checksumPolicyType                            | Defines how the server handles artifact checksums during deployment.                           | ✔️          | <p><br></p> | <p><br></p> | ✔️          |
| clientTlsCertificate                          | Specifies a client-side TLS certificate to use for authenticating to the remote repository.    | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |
| contentSynchronisation                        | Configures properties for synchronizing content from a remote Artifactory instance.            | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |
| defaultDeploymentRepo                         | Sets the default local repository for artifacts deployed to this virtual repository.           | <p><br></p> | <p><br></p> | ✔️          | <p><br></p> |
| description                                   | Provides a short, human-readable summary of the repository's purpose.                          | ✔️          | ✔️          | ✔️          | ✔️          |
| downloadRedirect                              | Redirects client downloads directly to the source URL instead of proxying through Artifactory. | ✔️          | ✔️          | <p><br></p> | ✔️          |
| enableCookieManagement                        | Enables stateful cookie handling for requests to the remote repository.                        | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |
| environment                                   | Adds a tag to classify the repository for a specific lifecycle stage (e.g., dev, qa, prod).    | ✔️          | ✔️          | ✔️          | ✔️          |
| excludesPattern                               | Defines a list of file path patterns to block from deployment.                                 | ✔️          | ✔️          | ✔️          | ✔️          |
| failedRetrievalCachePeriodSecs                | Sets the time (in seconds) to cache a "not found" response for a failed artifact download.     | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |
| fetchJarsEagerly                              | For remote Maven repositories, proactively fetches JAR files when the index is updated.        | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |
| fetchSourcesEagerly                           | For remote Maven repositories, proactively fetches source JARs when the index is updated.      | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |
| forceMavenAuthentication                      | For virtual Maven repositories, requires authentication for all requests.                      | <p><br></p> | <p><br></p> | ✔️          | <p><br></p> |
| handleReleases                                | Determines if the repository can host stable, final release versions of packages.              | ✔️          | ✔️          | <p><br></p> | ✔️          |
| handleSnapshots                               | Determines if the repository can host development or pre-release versions of packages.         | ✔️          | ✔️          | <p><br></p> | ✔️          |
| hardFail                                      | Causes requests to fail immediately upon any network error with the remote repository.         | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |
| includesPattern                               | Defines a list of file path patterns that are allowed for deployment.                          | ✔️          | ✔️          | ✔️          | ✔️          |
| keyPair                                       | Assigns a GPG key pair to the virtual repository for signing metadata files.                   | <p><br></p> | <p><br></p> | ✔️          | <p><br></p> |
| localAddress                                  | Binds outgoing connections to the remote repository to a specific local IP address.            | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |
| maxUniqueSnapshots                            | Limits the number of unique snapshot or pre-release versions stored for an artifact.           | ✔️          | <p><br></p> | <p><br></p> | ✔️          |
| missedRetrievalCachePeriodSecs                | Sets the time (in seconds) to cache a "not found" response for remote repository metadata.     | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |
| notes                                         | Offers a space for longer, more detailed information about the repository.                     | ✔️          | ✔️          | ✔️          | ✔️          |
| offline                                       | Prevents Artifactory from making any network connections to the remote repository.             | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |
| optionalIndexCompressionFormats               | Defines additional compression formats for repository index files to support various clients.  | ✔️          | <p><br></p> | ✔️          | ✔️          |
| password                                      | Sets the password for authenticating to the remote repository.                                 | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |
| pomRepositoryReferencesCleanupPolicy          | For virtual Maven repositories, controls how to handle repository references in POM files.     | <p><br></p> | <p><br></p> | ✔️          | <p><br></p> |
| primaryKeyPairRef                             | Specifies the primary GPG key pair to use for signing metadata.                                | <p><br></p> | <p><br></p> | ✔️          | <p><br></p> |
| priorityResolution                            | Gives a repository priority during package resolution within a virtual repository.             | ✔️          | ✔️          | <p><br></p> | ✔️          |
| projectKey                                    | Links the repository to a specific Project for organization and permission management.         | ✔️          | ✔️          | ✔️          | ✔️          |
| propertySets                                  | Associates required metadata fields (properties) with the repository to enforce governance.    | ✔️          | ✔️          | <p><br></p> | ✔️          |
| proxy                                         | Specifies a pre-configured network proxy to use for requests to the remote repository.         | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |
| rejectInvalidJars                             | For remote Java-based repositories, rejects and does not cache invalid or corrupt JAR files.   | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |
| remoteRepoChecksumPolicyType                  | Defines how to handle checksums for artifacts downloaded from the remote repository.           | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |
| repoLayoutRef                                 | Assigns a folder structure layout to the repository, enabling metadata parsing.                | ✔️          | ✔️          | ✔️          | ✔️          |
| repositories                                  | Defines the list of underlying repositories aggregated by the virtual repository.              | <p><br></p> | <p><br></p> | ✔️          | <p><br></p> |
| retrievalCachePeriodSecs                      | Sets the time (in seconds) to cache metadata for successfully downloaded remote artifacts.     | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |
| shareConfiguration                            | Shares the remote repository's configuration with other federated Artifactory instances.       | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |
| snapshotVersionBehavior                       | Defines how the server stores and manages snapshot versions.                                   | ✔️          | <p><br></p> | <p><br></p> | ✔️          |
| socketTimeoutMillis                           | Sets the timeout (in milliseconds) for network connections to the remote repository.           | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |
| storeArtifactsLocally                         | Controls whether artifacts downloaded from the remote repository are cached locally.           | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |
| suppressPomConsistencyChecks                  | For Maven repositories, disables validation checks on deployed POM files.                      | ✔️          | ✔️          | <p><br></p> | ✔️          |
| synchronizeProperties                         | Synchronizes artifact properties from a remote Artifactory instance.                           | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |
| unusedArtifactsCleanupEnabled                 | Enables the automatic cleanup of unused cached artifacts from the remote repository.           | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |
| unusedArtifactsCleanupPeriodHours             | Sets the time (in hours) an unused cached artifact must wait before cleanup.                   | <p><br></p> | ✔️          | <p><br></p> | <p><br></p> |

\
\


After adding your desired configurations, enter :x to save the template file. Creates and saves a json template and exits the interactive terminal.

<figure><img src="https://lh7-rt.googleusercontent.com/docsz/AD_4nXcLs0RweTxrHgnLRgWVYfDcQv8yrWziUQjVMNFXfvpoRKhKAmLaTl6AmNzuce5wJXsSQkjN-UifLBXiT3E-8t3SPHBw8dWEs4NsMf7At_mMZjjqBbmuAc-ylWHsCgb3rHdeRSX-pw?key=2SCcR_Q1GYlbACRhGX5p2A" alt=""><figcaption></figcaption></figure>

\
The sample json template is as follows:\


```json
 {
  "description": "my npm local repository",
  "key": "my-npm-local", # example variable, ${repo-name}
  "packageType": "npm",
  "rclass": "local"
}
```

Reuse Template with Variables

\
Reuse the template by adding variables to the keys and provide value explicitly while executing the **jf rt repo-create command.**



For example

```
 jf rt repo-create repotemplate.json --vars "repo-name=my-npm-local"

```

If you want to pass multiple vars, enter the list of semicolon-separated(;) variables in the form of "key1=value1;key2=value2;..." (wrapped by quotes) to be replaced in the template. In the template, the variables should be used as follows: ${key1}.

```
 jf rt repo-create repotemplate.json --vars "repo-name=my-npm-local;package-type=npm;repo-type=local"
```

&#x20;&#x20;

### Creating / Updating Repositories

These two commands create a new repository and updates an existing a repository. Both commands accept as an argument a configuration template, which can be created by the **jf rt repo-template** command. The template also supports variables, which can be replaced with values, provided when it is used.

#### Commands Params

|                        |                                                                                                                                                                                                               |
| ---------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name           | rt repo-create / rt repo-update                                                                                                                                                                               |
| Abbreviation           | rt rc / rt ru                                                                                                                                                                                                 |
| **Command options:**   |                                                                                                                                                                                                               |
| `--server-id`          | <p>[Optional]<br>Artifactory Server ID configured using the 'jf config' command.</p>                                                                                                                          |
| `--vars`               | <p>[Optional]<br>List of semicolon-separated(;) variables in the form of "key1=value1;key2=value2;..." to be replaced in the template. In the template, the variables should be used as follows: ${key1}.</p> |
| **Command arguments:** |                                                                                                                                                                                                               |
| template path          | Specifies the local file system path for the template file to be used for the repository creation. The template can be created using the "jf rt rpt" command.                                                 |

#### Examples

**Example 1**

Create a repository, using the **template.json** file previously generated by the **repo-template** command.

```
jf rt repo-create template.json
```

**Example 2**

Update a repository, using the **template.json** file previously generated by the **repo-template** command.

```
jf rt repo-update template.json
```

**Example 3**

Update a repository, using the **template.json** file previously generated by the **repo-template** command. Replace the repo-name variable inside the template with a name for the updated repository.

```
jf rt repo-update template.json --vars "repo-name=my-repo"
```

### Deleting Repositories

This command permanently deletes a repository, including all of its content.

#### Commands Params

|                        |                                                                                                            |
| ---------------------- | ---------------------------------------------------------------------------------------------------------- |
| Command name           | rt repo-delete                                                                                             |
| Abbreviation           | rt rdel                                                                                                    |
| **Command options:**   |                                                                                                            |
| `--server-id`          | <p>[Optional]<br>Artifactory Server ID configured using the 'jf config' command.</p>                       |
| `--quiet`              | <p>[Default: $CI]<br>Set to true to skip the delete confirmation message.</p>                              |
| **Command arguments:** |                                                                                                            |
| repository key         | Specifies the repositories that should be removed. You can use wildcards to specify multiple repositories. |

#### Example

Delete a repository from Artifactory.

```
jf rt repo-delete generic-local
```

## Managing Replications

JFrog CLI offers commands creating and deleting replication jobs in Artifactory. To make it easier to create replication jobs, the commands which creates the replication job accepts a pre-defined configuration template file. This template file can also include variables, which can be later replaced with values, when creating the replication job. The configuration template file is created using the **jf rt replication-template** command.

### Creating a Configuration Template

This command creates a configuration template file, which will be used as an argument for the **jf rt replication-create** command.

When using this command to create the template, you can also provide replaceable variable, instead of fixes values. Then, when the template is used to create replication jobs, values can be provided to replace the variables in the template.

#### Commands Params

|                        |                                                                                                               |
| ---------------------- | ------------------------------------------------------------------------------------------------------------- |
| Command-name           | rt replication-template                                                                                       |
| Abbreviation           | rt rplt                                                                                                       |
| **Command options:**   | The command has no options.                                                                                   |
| **Command arguments:** |                                                                                                               |
| template path          | Specifies the local file system path for the template file created by the command. The file should not exist. |

#### Example

Create a configuration template, with two variables for the source and target repositories. Then, create a replication job using this template, and provide source and target repository names to replace the variables.

```
$ jf rt rplt template.json
Select replication job type (press Tab for options): push
Enter source repo key > ${source}
Enter target repo key > ${target}
Enter target server id (press Tab for options): my-server-id
Enter cron expression for frequency (for example: 0 0 12 * * ? will replicate daily) > 0 0 12 * * ?
You can type ":x" at any time to save and exit.
Select the next property > :x
[Info] Replication creation config template successfully created at template.json.
$
$ jf rt rplc template.json --vars "source=generic-local;target=generic-local"
[Info] Done creating replication job.
```

### Creating Replication Jobs

This command creates a new replication job for a repository. The command accepts as an argument a configuration template, which can be created by the **jf rt replication-template** command. The template also supports variables, which can be replaced with values, provided when it is used.

#### Commands Params

|                        |                                                                                                                                                                                                               |
| ---------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name           | replication-create                                                                                                                                                                                            |
| Abbreviation           | rt rplc                                                                                                                                                                                                       |
| **Command options:**   |                                                                                                                                                                                                               |
| `--server-id`          | <p>[Optional]<br>Artifactory Server ID configured using the 'jf config' command.</p>                                                                                                                          |
| `--vars`               | <p>[Optional]<br>List of semicolon-separated(;) variables in the form of "key1=value1;key2=value2;..." to be replaced in the template. In the template, the variables should be used as follows: ${key1}.</p> |
| **Command arguments:** |                                                                                                                                                                                                               |
| template path          | Specifies the local file system path for the template file to be used for the replication job creation. The template can be created using the "jf rt rplt" command.                                           |

#### Examples

**Example 1**

Create a replication job, using the **template.json** file previously generated by the **replication-template** command.

```
jf rt rplc template.json
```

**Example 2**

Update a replication job, using the **template.json** file previously generated by the **replication-template** command. Replace the source and target variables inside the template with the names of the replication source and target repositories.

```
jf rt rplc template.json --vars "source=my-source-repo;target=my-target-repo"
```

### Deleting Replication jobs

This command permanently deletes a replication jobs from a repository.

#### Commands Params

|                        |                                                                                      |
| ---------------------- | ------------------------------------------------------------------------------------ |
| Command name           | rt replication-delete                                                                |
| Abbreviation           | rt rpldel                                                                            |
| **Command options:**   |                                                                                      |
| `--server-id`          | <p>[Optional]<br>Artifactory Server ID configured using the 'jf config' command.</p> |
| `--quiet`              | <p>[Default: $CI]<br>Set to true to skip the delete confirmation message.</p>        |
| **Command arguments:** |                                                                                      |
| repository key         | The repository from which the replications will be deleted.                          |

#### Example

Delete a repository from Artifactory.

```
jf rt rpldel my-repo-name
```

## Managing Permission Targets

JFrog CLI offers commands creating, updating and deleting permission targets in Artifactory. To make it easier to create and update permission targets, the commands which create and update the permission targets accept a pre-defined configuration template file. This template file can also include variables, which can be later replaced with values, when creating or updating the permission target. The configuration template file is created using the **jf rt permission-target-template** command.

### Creating a Configuration Template

This command creates a configuration template file, which will be used as an argument for the **jf rt permission-target-create** and **jf rt permission-target-update** commands.

|                        |                                                                                                               |
| ---------------------- | ------------------------------------------------------------------------------------------------------------- |
| Command-name           | rt permission-target-template                                                                                 |
| Abbreviation           | rt ptt                                                                                                        |
| **Command options:**   | The command has no options.                                                                                   |
| **Command arguments:** |                                                                                                               |
| template path          | Specifies the local file system path for the template file created by the command. The file should not exist. |

### Creating / Updating Permission Targets

These commands create/update a permission target. The commands accept as an argument a configuration template, which should be created by the **jf rt permission-target-template** command beforehand. The template also supports variables, which can be replaced with values, provided when it is used.

|                        |                                                                                                                                                                                                               |
| ---------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command-name           | permission-target-create / permission-target-update                                                                                                                                                           |
| Abbreviation           | rt ptc / rt ptu                                                                                                                                                                                               |
| **Command arguments:** |                                                                                                                                                                                                               |
| template path          | Specifies the local file system path for the template file to be used for the permission target creation or update. The template should be created using the "jf rt ptt" command.                             |
| Command-name           | permission-target-create / permission-target-update                                                                                                                                                           |
| Abbreviation           | rt ptc / rt ptu                                                                                                                                                                                               |
| **Command options:**   |                                                                                                                                                                                                               |
| `--server-id`          | <p>[Optional]<br>Artifactory Server ID configured using the 'jf config' command.</p>                                                                                                                          |
| `--vars`               | <p>[Optional]<br>List of semicolon-separated(;) variables in the form of "key1=value1;key2=value2;..." to be replaced in the template. In the template, the variables should be used as follows: ${key1}.</p> |

### Deleting Permission Targets

This command permanently deletes a permission target.

|                        |                                                                                      |
| ---------------------- | ------------------------------------------------------------------------------------ |
| Command name           | rt permission-target-delete                                                          |
| Abbreviation           | rt ptdel                                                                             |
| **Command options:**   |                                                                                      |
| `--server-id`          | <p>[Optional]<br>Artifactory Server ID configured using the 'jf config' command.</p> |
| `--quiet`              | <p>[Default: $CI]<br>Set to true to skip the delete confirmation message.</p>        |
| **Command arguments:** |                                                                                      |
| permission target name | The permission target that should be removed.                                        |
