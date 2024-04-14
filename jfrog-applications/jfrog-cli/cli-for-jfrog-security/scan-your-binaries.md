# Scan your Binaries

The [on-demand binary scanning](https://jfrog.com/help/r/jfrog-security-documentation/xray-on-demand-binary-scan) enables you to point to a binary in your local file system and receive a report that contains a list of vulnerabilities, licenses, and policy violations for that binary prior to uploading the binary or build to Artifactory.

### Scanning Files on the Local File System

This _**jf scan**_ command scans files on the local file system with Xray.

***

**Note**

> This command requires:

* Version 3.29.0 or above of Xray
* Version 2.1.0 or above of JFrog CLI

***

#### Commands Params

|                       |                                                                                                                                                                                                                                                                                                                                                                    |
|-----------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **Command name**      | scan                                                                                                                                                                                                                                                                                                                                                               |
| **Abbreviation**      | s                                                                                                                                                                                                                                                                                                                                                                  |
| **Command options**   |                                                                                                                                                                                                                                                                                                                                                                    |
| --server-id           | <p>[Optional]<br>Server ID configured using the <em>jf c add</em> command. If not specified, the default configured server is used.</p>                                                                                                                                                                                                                        |
| --spec                | <p>[Optional]<br>Path to a file specifying the files to scan. If the pattern argument is provided to the command, this option should not be provided.</p>                                                                                                                                                                                                      |
| --project             | <p>[Optional]<br>JFrog project key, to enable Xray to determine security violations accordingly. The command accepts this option only if the --repo-path and --watches options are not provided. If none of the three options are provided, the command will show all known vulnerabilities.</p>                                                               |
| --repo-path           | <p>[Optional]<br>Artifactory repository path in the form of &#x3C;repository>/&#x3C;path in the repository>, to enable Xray to determine violations accordingly. The command accepts this option only if the --project and --watches options are not provided. If none of the three options are provided, the command will show all known vulnerabilities.</p> |
| --watches             | <p>[Optional]<br>A comma-separated list of Xray watches, to enable Xray to determine violations accordingly. The command accepts this option only if the --project and --repo-path options are not provided. If none of the three options are provided, the command will show all known vulnerabilities.</p>                                                   |
| --licenses            | <p>[Optional]<br>Set if you also require the list of licenses to be displayed.</p>                                                                                                                                                                                                                                                                       |
| --format=json         | <p>[Optional]<br>Produces a JSON file containing the scan results.</p>                                                                                                                                                                                                                                                                                         |
| **Command arguments** |                                                                                                                                                                                                                                                                                                                                                                    |
| **Pattern**           | Specifies the local file system path to artifacts to be scanned. You can specify multiple files by using wildcards.                                                                                                                                                                                                                                                |

#### Example 1

Scans all the files located at the path/ti/files/ file-system directory using the watch1 watch defined in Xray.

```
jf s "path/to/files/" --watches "watch1"
```

#### Example 2

Scans all the files located at the path/ti/files/ file-system directory using the _watch1_ and _watch2_ Watches defined in Xray.

```
jf s "path/to/files/" --watches "watch1,watch2"
```

#### Example 3

Scans all the zip files located at the path/ti/files/ file-system directory using the _watch1_ and _watch2_ Watches defined in Xray.

```
jf s "path/to/files/*.zip" --watches "watch1,watch2"
```

#### Example 4

Scans all the tgz files located at the path/ti/files/ file-system directory using the policies defined for project-1.

```
jf s "path/to/files/*.tgz" --project "project-1"
```

#### Example 5

Scans all the tgz files located in the current directory using the policies defined for the libs-local/release-artifacts/ path in Artifactory.

```
jf s "*.tgz" --repo-path "libs-local/release-artifacts/"
```

#### Example 6

Scans all the tgz files located at the current directory. Show all known vulnerabilities, regardless of the policies defined in Xray.

```
jf s "*.tgz"
```

### Scanning Docker Containers on the Local File System

This j\_**f docker scan**\_ command scans docker containers located on the local file-system using the _**docker client**_ and _**JFrog Xray**_. The containers don't need to be deployed to Artifactory or any other container registry before it can be scanned.

***

**Note**

> This command requires:

* Version 3.40.0 or above of Xray
* Version 2.11.0 or above of JFrog CLI

***

#### Commands Params

|                       |                                                                                                                                                                                                                                                                                                                                                                    |
|-----------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **Command name**      | docker scan                                                                                                                                                                                                                                                                                                                                                        |
| **Abbreviation**      |                                                                                                                                                                                                                                                                                                                                                                    |
| **Command options**   |                                                                                                                                                                                                                                                                                                                                                                    |
| --server-id           | <p>[Optional]<br>Server ID configured using the <em>jf c add</em> command. If not specified, the default configured server is used.</p>                                                                                                                                                                                                                        |
| --project             | <p>[Optional]<br>JFrog project key, to enable Xray to determine security violations accordingly. The command accepts this option only if the --repo-path and --watches options are not provided. If none of the three options are provided, the command will show all known vulnerabilities.</p>                                                               |
| --repo-path           | <p>[Optional]<br>Artifactory repository path in the form of &#x3C;repository>/&#x3C;path in the repository>, to enable Xray to determine violations accordingly. The command accepts this option only if the --project and --watches options are not provided. If none of the three options are provided, the command will show all known vulnerabilities.</p> |
| --watches             | <p>[Optional]<br>A comma separated list of Xray watches, to enable Xray to determine violations accordingly. The command accepts this option only if the --repo-path and --repo-path options are not provided. If none of the three options are provided, the command will show all known vulnerabilities.</p>                                                 |
| --licenses            | <p>[Optional]<br>Set if you also require the list of licenses to be displayed.</p>                                                                                                                                                                                                                                                                       |
| --format=json         | <p>[Optional]<br>Produces a JSON file containing the scan results.</p>                                                                                                                                                                                                                                                                                         |
| **Command arguments** |                                                                                                                                                                                                                                                                                                                                                                    |
| **Pattern**           | Specifies the local file system path to artifacts to be scanned. You can specify multiple files by using wildcards.                                                                                                                                                                                                                                                |


#### Example 1

Scan the local _reg1/repo1/img1:1.0.0_ container and show all known vulnerabilities, regardless of the policies defined in Xray.

```
$ docker images
REPOSITORY           TAG       IMAGE ID       CREATED         SIZE
reg1/repo1/img1   1.0.0     6446ea57df7b   19 months ago   5.57MB
$ 
$ jf docker scan reg1/repo1/img1:1.0.0
```

#### Example 2

Scan the local _reg1/repo1/img1:1.0.0_ container and show all violations according to the policy associated with _my-project_ JFrog project.

```
$ docker images
REPOSITORY           TAG       IMAGE ID       CREATED         SIZE
reg1/repo1/img1   1.0.0     6446ea57df7b   19 months ago   5.57MB
$ 
$ jf docker scan reg1/repo1/img1:1.0.0 --project my-project
```

#### Example 3

Scan the local _reg1/repo1/img1:1.0.0_ container and show all violations according to the policy associated with _my-watch_ Xray Watch.

```
$ docker images
REPOSITORY           TAG       IMAGE ID       CREATED         SIZE
reg1/repo1/img1   1.0.0     6446ea57df7b   19 months ago   5.57MB
$ 
$ jf docker scan reg1/repo1/img1:1.0.0 --watches my-watch
```

#### Example 4

Scan the local _reg1/repo1/img1:1.0.0_ container and show all violations according to the policy associated with the _releases-local/app1/_ path in Artifactory.

```
$ docker images
REPOSITORY           TAG       IMAGE ID       CREATED         SIZE
reg1/repo1/img1   1.0.0     6446ea57df7b   19 months ago   5.57MB
$ 
$ jf docker scan reg1/repo1/img1:1.0.0 --repo-path releases-local/app1/
```

### Scanning Image Tarballs on the Local File System

The ‘`scan`’ command can be used to scan tarballs of Docker and OCI images on the local file system.

It requires saving the image on the file system as an uncompressed tarball using a compliant tool, and then scanning it with the ‘`jf s`’ command. The image must be saved to the file system uncompressed, in a `<name>.tar` file name.

***

**Note**

> This command requires:

* Version 3.61.5 or above of Xray.
* Version 2.14.0 or above of JFrog CLI.

***

#### Docker Client

Use Docker client ‘`docker save`’ command to save the image to the file system for scanning.

**Example**:

```
$ docker images
REPOSITORY TAG IMAGE ID CREATED SIZE
my-image 1.0.0 aaaaabbcccddd 2 months ago 1.12MB

$ docker save --output my-image-docker.tar my-image:1.0.0
$ jf s my-image-docker.tar
```

#### Skopeo

Use Skopeo CLI to save an image to the file system. Output image can be either OCI or Docker format.

**Example**:

```
$ docker images
REPOSITORY TAG IMAGE ID CREATED SIZE
my-image 1.0.0 aaaaabbcccddd 2 months ago 1.12MB

// Scan an image in Docker format
$ skopeo copy docker-daemon:my-image:1.0.0 docker-archive:my-image-docker.tar
$ jf s my-image-docker.tar

// Scan an image in OCI format
$ skopeo copy docker-daemon:my-image:1.0.0 oci-archive:my-image-oci.tar
$ jf s my-image-oci.tar
```

#### Podman

Use Podman CLI to save an image to the file system. Output image can be either OCI or Docker format.

**Example**:

```
$ podman images
REPOSITORY TAG IMAGE ID CREATED SIZE
my-image 1.0.0 aaaaabbcccddd 2 months ago 1.12MB

// Scan an image in Docker format
$ podman save --format=docker-archive -o my-image-docker.tar my-image:1.0.0
$ jf s my-image-docker.tar

// Scan an image in OCI format
$ podman save --format=oci -o my-image-oci.tar my-image:1.0.0
$ jf s my-image-oci.tar
```

#### Kaniko

Use Kaniko ‘`--tarPath’` flag to save built images to the file system, and later scan them with JFrog CLI. The example below is running Kaniko in Docker.

**Example**:

```
$ cat Dockerfile

FROM alpine:3.16

$ docker run -it --rm -v $(pwd):/workspace gcr.io/kaniko-project/executor:v1.8.1-debug -f Dockerfile --no-push --tarPath my-image.tar -d my-image:1.0 -c . --cleanup

$ jf s my-image.tar
```
