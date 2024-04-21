
# Managing Configuration Entities

JFrog CLI offers a set of commands for managing Artifactory configuration entities.

## Creating Users

This command allows creating a bulk of users. The details of the users are provided in a CSV format file. Here's the file format.

```
"username","password","email"
"username1","password1","john@c.com"
"username2","password1","alice@c.com"
```

> **Note**: The first line in the CSV is cells' headers. It is mandatory and is used by the command to map the cell value to the users' details.

The CSV can include additional columns, with different headers, which will be ignored by the command.

### Commands Params

|                   |                                                                                                                                             |
|-------------------|---------------------------------------------------------------------------------------------------------------------------------------------|
| Command-name      | rt users-create                                                                                                                             |
| Abbreviation      | rt uc                                                                                                                                       |
| Command options   |                                                                                                                                             |
| --server-id       | <p>[Optional]<br>Artifactory server ID configured using the config command.</p>                                                             |
| --csv             | <p>[Mandatory]<br>Path to a CSV file with the users' details. The first row of the file should include the name,password,email headers.</p> |
| --replace         | <p>[Optional]<br>Set to true if you'd like existing users or groups to be replaced.</p>                                                     |
| --users-groups    | <p>[Optional]<br>A list of comma-separated(,) groups for the new users to be associated to.</p>                                             |
| Command arguments | The command accepts no arguments                                                                                                            |

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

|                   |                                                                                                                                                                                |
|-------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command-name      | rt users-delete                                                                                                                                                                |
| Abbreviation      | rt udel                                                                                                                                                                        |
| Command options   |                                                                                                                                                                                |
| --server-id       | <p>[Optional]<br>Artifactory server ID configured using the config command.</p>                                                                                                |
| --csv             | <p>[Optional]<br>Path to a csv file with the usernames to delete. The first row of the file is the reserved for the cells' headers. It must include the "username" header.</p> |
| Command arguments |                                                                                                                                                                                |
| users list        | comma-separated(,) list of usernames to delete. If the --csv command option is used, then this argument becomes optional.                                                      |

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

|                   |                                                                                 |
|-------------------|---------------------------------------------------------------------------------|
| Command-name      | rt group-create                                                                 |
| Abbreviation      | rt gc                                                                           |
| Command options   |                                                                                 |
| --server-id       | <p>[Optional]<br>Artifactory server ID configured using the config command.</p> |
| Command arguments |                                                                                 |
| group name        | The name of the group to create.                                                |

### Example

Create a new group name **reviewers** .

```
jf rt group-create reviewers
```

## Adding Users to Groups

This command adds a list fo existing users to a group.

### Commands Params

|                   |                                                                                 |
|-------------------|---------------------------------------------------------------------------------|
| Command-name      | rt group-add-users                                                              |
| Abbreviation      | rt gau                                                                          |
| Command options   |                                                                                 |
| --server-id       | <p>[Optional]<br>Artifactory server ID configured using the config command.</p> |
| Command arguments |                                                                                 |
| group name        | The name of the group to add users to.                                          |
| users list        | Comma-seperated list of usernames to add to the specified group.                |

### Example

Add to group reviewers the users with the following usernames: u1, u2 and u3.

```
jf rt group-add-users "reviewers" "u1,u2,u3"
```

## Deleting Groups

This command deletes a group.

### Commands Params

|                   |                                                                                 |
|-------------------|---------------------------------------------------------------------------------|
| Command-name      | rt group-delete                                                                 |
| Abbreviation      | rt gdel                                                                         |
| Command options   |                                                                                 |
| --server-id       | <p>[Optional]<br>Artifactory server ID configured using the config command.</p> |
| Command arguments |                                                                                 |
| group name        | The name of the group to delete.                                                |

### Example

Delete the **reviewers** group.

```
jf rt group-delete "reviewers"
```

## Managing Repositories

JFrog CLI offers a set of commands for managing Artifactory repositories. You can create, update and delete repositories. To make it easier to manage repositories, the commands which create and update the repositories accept a pre-defined configuration template file. This template file can also include variables, which can be later replaced with values, when creating or updating the repositories. The configuration template file is created using the **jf rt repo-template** command.

### Creating or Configuration Template

This is an interactive command, which creates a configuration template file. This file should be used as an argument for the **jf rt repo-create** or the **jf rt repo-update** commands.

When using this command to create the template, you can also provide replaceable variable, instead of fixes values. Then, when the template is used to create or update repositories, values can be provided to replace the variables in the template.

#### Commands Params

|                   |                                                                                                               |
|-------------------|---------------------------------------------------------------------------------------------------------------|
| Command-name      | rt repo-template                                                                                              |
| Abbreviation      | rt rpt                                                                                                        |
| Command options   | The command has no options.                                                                                   |
| Command arguments |                                                                                                               |
| template path     | Specifies the local file system path for the template file created by the command. The file should not exist. |

#### Example

Create a configuration template, with a variable for the repository name. Then, create a repository using this template, and provide repository name to replace the variable.

```
$ jf rt repo-template template.json

Select the template type (press Tab for options): create
Insert the repository key > ${repo-name}
Select the repository class (press Tab for options): local
Select the repository's package type (press Tab for options): generic
You can type ":x" at any time to save and exit.
Select the next configuration key (press Tab for options): :x
[Info] Repository configuration template successfully created at template.json.
$
$ jf rt repo-create template.json --vars "repo-name=my-repo"
[Info] Creating local repository...
[Info] Done creating repository.
```

### Creating / Updating Repositories

These two commands create a new repository and updates an existing a repository. Both commands accept as an argument a configuration template, which can be created by the **jf rt repo-template** command. The template also supports variables, which can be replaced with values, provided when it is used.

#### Commands Params

|                   |                                                                                                                                                                                                               |
|-------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command-name      | rt repo-create / rt repo-update                                                                                                                                                                               |
| Abbreviation      | rt rc / rt ru                                                                                                                                                                                                 |
| Command options   |                                                                                                                                                                                                               |
| --server-id       | <p>[Optional]<br>Artifactory server ID configured using the config command.</p>                                                                                                                               |
| --vars            | <p>[Optional]<br>List of semicolon-separated(;) variables in the form of "key1=value1;key2=value2;..." to be replaced in the template. In the template, the variables should be used as follows: ${key1}.</p> |
| Command arguments |                                                                                                                                                                                                               |
| template path     | Specifies the local file system path for the template file to be used for the repository creation. The template can be created using the "jf rt rpt" command.                                                 |

#### Examples
##### Example 1

Create a repository, using the **template.json** file previously generated by the **repo-template** command.

```
jf rt repo-create template.json
```

##### Example 2

Update a repository, using the **template.json** file previously generated by the **repo-template** command.

```
jf rt repo-update template.json
```

##### Example 3

Update a repository, using the **template.json** file previously generated by the **repo-template** command. Replace the repo-name variable inside the template with a name for the updated repository.

```
jf rt repo-update template.json --vars "repo-name=my-repo"
```

### Deleting Repositories

This command permanently deletes a repository, including all of its content.

#### Commands Params

|                   |                                                                                                            |
|-------------------|------------------------------------------------------------------------------------------------------------|
| Command name      | rt repo-delete                                                                                             |
| Abbreviation      | rt rdel                                                                                                    |
| Command options   |                                                                                                            |
| --server-id       | <p>[Optional]<br>Artifactory server ID configured using the config command.</p>                            |
| --quiet           | <p>[Default: $CI]<br>Set to true to skip the delete confirmation message.</p>                              |
| Command arguments |                                                                                                            |
| repository key    | Specifies the repositories that should be removed. You can use wildcards to specify multiple repositories. |

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

|                   |                                                                                                               |
|-------------------|---------------------------------------------------------------------------------------------------------------|
| Command-name      | rt replication-template                                                                                       |
| Abbreviation      | rt rplt                                                                                                       |
| Command options   | The command has no options.                                                                                   |
| Command arguments |                                                                                                               |
| template path     | Specifies the local file system path for the template file created by the command. The file should not exist. |

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

|                   |                                                                                                                                                                                                               |
|-------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command-name      | replication-create                                                                                                                                                                                            |
| Abbreviation      | rt rplc                                                                                                                                                                                                       |
| Command options   |                                                                                                                                                                                                               |
| --server-id       | <p>[Optional]<br>Artifactory server ID configured using the config command.</p>                                                                                                                               |
| --vars            | <p>[Optional]<br>List of semicolon-separated(;) variables in the form of "key1=value1;key2=value2;..." to be replaced in the template. In the template, the variables should be used as follows: ${key1}.</p> |
| Command arguments |                                                                                                                                                                                                               |
| template path     | Specifies the local file system path for the template file to be used for the replication job creation. The template can be created using the "jf rt rplt" command.                                           |

#### Examples
##### Example 1

Create a replication job, using the **template.json** file previously generated by the **replication-template** command.

```
jf rt rplc template.json
```

##### Example 2

Update a replication job, using the **template.json** file previously generated by the **replication-template** command. Replace the source and target variables inside the template with the names of the replication source and target repositories.

```
jf rt rplc template.json --vars "source=my-source-repo;target=my-target-repo"
```

### Deleting Replication jobs

This command permanently deletes a replication jobs from a repository.

#### Commands Params

|                   |                                                                                 |
|-------------------|---------------------------------------------------------------------------------|
| Command name      | rt replication-delete                                                           |
| Abbreviation      | rt rpldel                                                                       |
| Command options   |                                                                                 |
| --server-id       | <p>[Optional]<br>Artifactory server ID configured using the config command.</p> |
| --quiet           | <p>[Default: $CI]<br>Set to true to skip the delete confirmation message.</p>   |
| Command arguments |                                                                                 |
| repository key    | The repository from which the replications will be deleted.                     |

#### Example

Delete a repository from Artifactory.

```
jf rt rpldel my-repo-name
```

## Managing Permission Targets

JFrog CLI offers commands creating, updating and deleting permission targets in Artifactory. To make it easier to create and update permission targets, the commands which create and update the permission targets accept a pre-defined configuration template file. This template file can also include variables, which can be later replaced with values, when creating or updating the permission target. The configuration template file is created using the **jf rt permission-target-template** command.

### Creating a Configuration Template

This command creates a configuration template file, which will be used as an argument for the **jf rt permission-target-create** and **jf rt permission-target-update** commands.

|                   |                                                                                                               |
|-------------------|---------------------------------------------------------------------------------------------------------------|
| Command-name      | rt permission-target-template                                                                                 |
| Abbreviation      | rt ptt                                                                                                        |
| Command options   | The command has no options.                                                                                   |
| Command arguments |                                                                                                               |
| template path     | Specifies the local file system path for the template file created by the command. The file should not exist. |

### Creating / Updating Permission Targets

These commands create/update a permission target. The commands accept as an argument a configuration template, which should be created by the **jf rt permission-target-template** command beforehand. The template also supports variables, which can be replaced with values, provided when it is used. 

|                   |                                                                                                                                                                                                               |
|-------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Command-name      | permission-target-create / permission-target-update                                                                                                                                                           |
| Abbreviation      | rt ptc / rt ptu                                                                                                                                                                                               |
| Command arguments |                                                                                                                                                                                                               |
| template path     | Specifies the local file system path for the template file to be used for the permission target creation or update. The template should be created using the "jf rt ptt" command.                             |
| Command-name      | permission-target-create / permission-target-update                                                                                                                                                           |
| Abbreviation      | rt ptc / rt ptu                                                                                                                                                                                               |
| Command options   |                                                                                                                                                                                                               |
| --server-id       | <p>[Optional]<br>Artifactory server ID configured using the config command.</p>                                                                                                                               |
| --vars            | <p>[Optional]<br>List of semicolon-separated(;) variables in the form of "key1=value1;key2=value2;..." to be replaced in the template. In the template, the variables should be used as follows: ${key1}.</p> |

### Deleting Permission Targets

This command permanently deletes a permission target.

|                        |                                                                                 |
|------------------------|---------------------------------------------------------------------------------|
| Command name           | rt permission-target-delete                                                     |
| Abbreviation           | rt ptdel                                                                        |
| Command options        |                                                                                 |
| --server-id            | <p>[Optional]<br>Artifactory server ID configured using the config command.</p> |
| --quiet                | <p>[Default: $CI]<br>Set to true to skip the delete confirmation message.</p>   |
| Command arguments      |                                                                                 |
| permission target name | The permission target that should be removed.                                   |
