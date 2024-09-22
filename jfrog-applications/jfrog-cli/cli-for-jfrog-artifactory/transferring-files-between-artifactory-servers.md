
# Transferring Files Between Artifactory Servers

## Overview

The transfer-files command allows transferring (copying) all the files stored in one Artifactory instance to a different Artifactory instance. The command allows transferring the files stored in a single or multiple repositories. The command expects the relevant repository to already exist on the target instance and have the same name and type as the repositories on the source.

## Limitations

1. Artifacts in remote repositories caches are not transferred.
2. The files transfer process allows transferring files that were created or modified on the source instance after the process started. However, files that were deleted on the source instance after the process started, are not deleted on the target instance by the process.
3. The files transfer process allows transferring files that were created or modified on the source instance after the process started. The custom properties of those files are also updated on the target instance. However, if only the custom properties of those file were modified on the source, but not the files' content, the properties are not modified on the target instance by the process.
4. The source and target repositories should have the same name and type.
5. Since the file are pushed from the source to the target instance, the source instance must have network connection to the target.

## Before You Begin

1. Ensure that you can log in to the UI of both the source and target instances with users that have admin permissions and that you have the connection details (including credentials) to both instances.
2. Ensure that all the repositories on source Artifactory instance which files you'd like to transfer, also exist on the target instance, and have the same name and type on both instances.
3. Ensure that JFrog CLI is installed on a machine that has network access to both the source and target instances.

## Running the Transfer Process

### Step 1 - Set Up the Source Instance for Files Transfer

To set up the source instance for files transfer, you must install the **data-transfer** user plugin in the primary node of the source instance. This section guides you through the installation steps.

1. Install JFrog CLI on the primary node machine of the source instance as described [here](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/cli-for-jfrog-artifactory#installing-the-data-transfer-user-plugin-on-the-source-machine-manually).
2.  Configure the connection details of the source Artifactory instance with your admin credentials by running the following command from the terminal.

    ```
    jf c add source-server
    ```
3. Ensure that the **JFROG\_HOME** environment variable is set and holds the value of JFrog installation directory. It usually points to the **/opt/jfrog** directory. In case the variable isn't set, set its value to point to the correct directory as described in the JFrog Product Directory Structure article.
4. If the source instance has internet access, you can install the **data-transfer** user plugin on the source machine automatically by running the following command from the terminal `jf rt transfer-plugin-install source-server`. If however the source instance has no internet access, install the plugin manually as described [here](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/cli-for-jfrog-artifactory#installing-the-data-transfer-user-plugin-on-the-source-machine-manually).

### Step 2 - Push the Files from the Source to the Target Instance

Install JFrog CLI on any machine that has access to both the source and the target JFrog instances. To do this, follow the steps described [here](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/cli-for-jfrog-artifactory/transferring-files-between-artifactory-servers#installing-jfrog-cli-on-a-machine-with-network-access-to-the-source-and-target-machines).

Run the following command to start pushing the files from all the repositories in source instance to the target instance.

```
jf rt transfer-files source-server target-server
```

This command may take a few days to push all the files, depending on your system size and your network speed. While the command is running, It displays the transfer progress visually inside the terminal.

![transfer-files-progress](../../.gitbook/assets/transfer-4.png)

If you're running the command in the background, you use the following command to view the transfer progress.

```
jf rt transfer-files --status
```

![transfer-files-progress](../../.gitbook/assets/transfer-5.png)

In case you do not wish to transfer the files from all repositories, or wish to run the transfer in phases, you can use the `--include-repos` and `--exclude-repos` command options. Run the following command to see the usage of these options.

```
jf rt transfer-files -h
```

If the traffic between the source and target instance needs to be routed through an HTTPS proxy, refer to [this](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/cli-for-jfrog-artifactory#routing-the-traffic-from-the-source-to-the-target-through-an-https-proxy) section.

You can stop the transfer process by hitting on **CTRL+C** if the process is running in the foreground, or by running the following command, if you're running the process in the background.

```
jf rt transfer-files --stop
```

The process will continue from the point it stopped when you re-run the command.

While the file transfer is running, monitor the load on your source instance, and if needed, reduce the transfer speed or increase it for better performance. For more information, see the [Controlling the File Transfer Speed](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/cli-for-jfrog-artifactory#controlling-the-file-transfer-speed) section.

A path to an errors summary file will be printed at the end of the run, referring to a generated CSV file. Each line on the summary CSV represents an error of a file that failed to be transferred. On subsequent executions of the `jf rt transfer-files` command, JFrog CLI will attempt to transfer these files again.

Once the `jf rt transfer-files` command finishes transferring the files, you can run it again to transfer files which were created or modified during the transfer. You can run the command as many times as needed. Subsequent executions of the command will also attempt to transfer files failed to be transferred during previous executions of the command.

**Note:**

> Read more about how the transfer files works in [this](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli/cli-for-jfrog-cloud-transfer#how-does-file-transfer-work) section.

***

## Installing the data-transfer User Plugin on the Source Machine Manually

To install the **data-transfer** user plugin on the source machine manually, follow these steps.

1. Download the following two files from a machine that has internet access. Download **data-transfer.jar** from https://releases.jfrog.io/artifactory/jfrog-releases/data-transfer/\[RELEASE]/lib/data-transfer.jar and **dataTransfer.groovy** from https://releases.jfrog.io/artifactory/jfrog-releases/data-transfer/\[RELEASE]/dataTransfer.groovy
2. Create a new directory on the primary node machine of the source instance and place the two files you downloaded inside this directory.
3.  Install the **data-transfer** user plugin by running the following command from the terminal. Replace the _\[plugin files dir]_ token with the full path to the directory which includes the plugin files you downloaded.

    ```
    jf rt transfer-plugin-install source-server --dir "[plugin files dir]"
    ```

## Installing JFrog CLI on the Source Instance Machine

Install JFrog CLI on your source instance by using one of the \[#JFrog CLI Installers]. For example:

```sh
curl -fL https://install-cli.jfrog.io | sh
```

**Note**

If the source instance is running as a docker container, and you're not able to install JFrog CLI while inside the container, follow these steps.

1. Connect to the host machine through the terminal.
2.  Download the JFrog CLI executable into the correct directory by running this command:

    `curl -fL https://getcli.jfrog.io/v2-jf | sh`
3.  Copy the JFrog CLI executable you've just downloaded into the container, by running the following docker command. Make sure to replace _\[the container name]_ with the name of the container.

    `docker cp jf [the container name]:/usr/bin/jf`
4.  Connect to the container and run the following command to ensure JFrog CLI is installed:

    `jf -v`

## How Does Files Transfer Work?

### Files Transfer Phases

The `jf rt transfer-files` command pushes the files from the source instance to the target instance as follows:

* The files are pushed for each repository, one by one in sequence.
* For each repository, the process includes the following three phases:
  * **Phase 1** pushes all the files in the repository to the target.
  * **Phase 2** pushes files which have been created or modified after phase 1 started running (diffs).
  * **Phase 3** attempts to push files which failed to be transferred in earlier phases (Phase 1 or Phase 2) or in previous executions of the command.
* If Phase 1 finished running for a specific repository, and you run the `jf rt transfer-files` command again, only **Phase 2** and **Phase 3** will be triggered. You can run the `jf rt transfer-files` as many times as needed, till you are ready to move your traffic to the target instance permanently. In any subsequent run of the command, **Phase 2** will transfer the newly created and modified files and **Phase 3** will retry transferring files which failed to be transferred in previous phases and also in previous runs of the command.

**Using Replication**

> To help reduce the time it takes for Phase 2 to run, you may configure Event Based Push Replication for some or all of the local repositories on the source instance. With Replication configured, when files are created or updated on the source repository, they are immediately replicated to the corresponding repository on the target instance. The replication can be configured at any time. Before, during or after the files transfer process.

***

### Files Transfer State

You can run the `jf rt transfer-files` command multiple times. This is needed to allow transferring files which have been created or updated after previous command executions. To achieve this, JFrog CLI stores the current state of the files transfer process in a directory named **transfer** located under the JFrog CLI home directory. You can usually find this directory at this location `~/.jfrog/transfer`.

JFrog CLI uses the state stored in this directory to avoid repeating transfer actions performed in previous executions of the command. For example, once **Phase 1** is completed for a specific repository, subsequent executions of the command will skip **Phase 1** and run **Phase 2** and **Phase 3** only.

In case you'd like to ignore the stored state, and restart the files transfer from scratch, you can add the `--ignore-state` option to the `jf rt transfer-files` command.

## Installing JFrog CLI on a Machine with Network Access to the Source and Target Machines

It is recommended to run the `transfer-files` command from a machine that has network access to the source Artifactory URL. This allows spreading the transfer load on all the Artifactory cluster nodes. This machine should also have network access to the target Artifactory URL.

Follows these steps to installing JFrog CLI on that machine.

1.  Install JFrog CLI by using one of the \[#JFrog CLI Installers]. For example:

    ```sh
    curl -fL https://install-cli.jfrog.io | sh
    ```

2. If your source instance is accessible only through an HTTP/HTTPS proxy, set the proxy environment variable as described \[#here].
3.  Configure the connection details of the source Artifactory instance with your admin credentials. Run the following command and follow the instructions.

    ```
    jf c add source-server
    ```
4.  Configure the connection details of the target Artifactory instance as follows.

    ```
    jf c add target-server
    ```

## Controlling the File Transfer Speed

The `jf rt transfer-files` command pushes the binaries from the source instance to the target instance. This transfer can take days, depending on the size of the total data transferred, the network bandwidth between the source and the target instance, and additional factors.

Since the process is expected to run while the source instance is still being used, monitor the instance to ensure that the transfer does not add too much load to it. Also, you might decide to increase the load for faster a transfer rate, while you monitor the transfer. This section describes how to control the file transfer speed.

By default, the `jf rt transfer-files` command uses 8 working threads to push files from the source instance to the target instance. Reducing this value will cause slower transfer speed and lower load on the source instance, and increasing it will do the opposite. We therefore recommend increasing it gradually. This value can be changed while the `jf rt transfer-files` command is running. There's no need to stop the process to change the total of working threads. The new value set will be cached by JFrog CLI and also used for subsequent runs from the same machine. To set the value, simply run the following interactive command from a new terminal window on the same machine which runs the `jf rt transfer-files` command.

```
jf rt transfer-settings
```

**Build-info repositories**

> When transferring files in build-info repositories, JFrog CLI limits the total of working threads to 8. This is done in order to limit the load on the target instance while transferring build-info.

***

## Routing the Traffic from the Source to the Target Through an HTTPS Proxy

The `jf rt transfer-files` command pushes the files directly from the source to the target instance over the network. In case the traffic from the source instance needs to be routed through an HTTPS proxy, follow these steps.

1. Define the proxy details in the source instance UI as described in the [Managing Proxies documentation](https://jfrog.com/help/r/jfrog-platform-administration-documentation/managing-proxies).
2.  When running the `jf rt transfer-files` command, add the `--proxy-key` option to the command, with Proxy Key you configured in the UI as the option value. For example, if the Proxy Key you configured is my-proxy-key, run the command as follows:

    ```
    jf rt transfer-files my-source my-target --proxy-key my-proxy-key
    ```
