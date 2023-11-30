# Download and Install

## JFrog CLI v2

### Overview

JFrog CLI v2 was launched in July 2021. It includes changes to the functionality and usage of some of the legacy JFrog CLI commands. The changes are the result of feedback we received from users over time through GitHub, making the usage and functionality easier and more intuitive. For example, some of the default values changed, and are now more consistent across different commands. We also took this opportunity for improving and restructuring the code, as well as replacing old and deprecated functionality.

Most of the changes included in v2 are breaking changes compared to the v1 releases. We therefore packaged and released these changes under JFrog CLI v2, allowing users to migrate to v2 only when they are ready.

New enhancements to JFrog CLI are planned to be introduced as part of V2 only. V1 receives very little development attention nowadays. We therefore encourage users who haven't yet migrated to V2, to do so.

### List of changes in JFrog CLI v2

1.  The default value of the _**--flat**_ option is now set to false for the _**jfrog rt upload**_ command.
2.  The deprecated syntax of the _**jfrog rt mvn**_ command is no longer supported. To use the new syntax, the project needs to be first configured using the **jfrog rt mvnc** command.
3.  The deprecated syntax of the _**jfrog rt gradle**_ command is no longer supported. To use the new syntax, the project needs to be first configured using the _**jfrog rt gradlec**_ command.
4.  The deprecated syntax of the **jfrog rt npm** and _**jfrog rt npm-ci**_ commands is no longer supported. To use the new syntax, the project needs to be first configured using the _**jfrog rt npmc**_ command.
5.  The deprecated syntax of the _**jfrog rt go**_ command is no longer supported. To use the new syntax, the project needs to be first configured using the _**jfrog rt go-config**_ command.
6.  The deprecated syntax of the _**jfrog rt nuget**_ command is no longer supported. To use the new syntax, the project needs to be first configured using the _**jfrog rt nugetc**_ command.
7.  All Bintray commands are removed.
8.  The _**jfrog rt config**_ command is removed and replaced by the _**jfrog config add**_ command.
9.  The _**jfrog rt use**_ command is removed and replaced with the _**jfrog config use**_.
10. The _**--props**_ command option and _**props**_ file spec property for the _**jfrog rt upload**_ command are removed, and replaced with the _**--target-props**_ command option and _**targetProps**_ file spec property respectively.
11. The following commands are removed 
    ```
    jfrog rt release-bundle-create
    jfrog rt release-bundle-delete
    jfrog rt release-bundle-distribute
    jfrog rt release-bundle-sign
    jfrog rt release-bundle-update
    ```
    and replaced with the following commands respectively 
    ```
    jfrog ds release-bundle-create
    jfrog ds release-bundle-delete
    jfrog ds release-bundle-distribute
    jfrog ds release-bundle-sign
    jfrog ds release-bundle-update
    ```
12. The _**jfrog rt go-publish**_ command now only supports Artifactory version 6.10.0 and above. Also, the command no longer accepts the target repository as an argument. The target repository should be pre-configured using the _**jfrog rt go-config**_ command.
13. The _**jfrog rt go**_ command no longer falls back to the VCS when dependencies are not found in Artifactory.
14. The _**--deps**_, _**--publish-deps**_, _**--no-registry**_ and _**--self**_ options of the _**jfrog rt go-publish**_ command are now removed.
15. The _**--apiKey**_ option is now removed. The API key should now be passed as the value of the _**--password**_ option.
16. The _**--exclude-patterns**_ option is now removed, and replaced with the _**--exclusions**_ option. The same is true for the _**excludePatterns**_ file spec property, which is replaced with the _**exclusions**_ property.
17. The _**JFROG_CLI_JCENTER\_REMOTE\_SERVER**_ and _**JFROG_CLI_JCENTER\_REMOTE\_REPO**_ environment variables are now removed and replaced with the _**JFROG_CLI_EXTRACTORS_REMOTE**_ environment variable.
18. The _**JFROG_CLI_HOME**_ environment variable is now removed and replaced with the _**JFROG_CLI_HOME_DIR**_ environment variable.
19. The _**JFROG_CLI_OFFER_CONFIG**_ environment variable is now removed and replaced with the _**CI**_ environment variable. Setting CI to true disables all prompts.
20. The directory structure is now changed when the _**jfrog rt download**_ command is used with placeholders and -_**-flat=false**_ (--flat=false is now the default). When placeholders are used, the value of the _**--flat**_ option is ignored.
21. When the **jfrog rt upload** command now uploads symlinks to Artifactory, the target file referenced by the symlink is uploaded to Artifactory with the symlink name. If the **--symlink** options is used, the symlink itself (not the referenced file) is uploaded, with the referenced file as a property attached to the file.

## Installation

To download the executable, please visit the [JFrog CLI Download Site](https://www.jfrog.com/getcli/).

You can also download the sources from the [JFrog CLI Project](https://github.com/JFrog/jfrog-cli-go) on GitHub where you will also find instructions on how to build JFrog CLI.

The legacy name of JFrog CLI's executable is _**jfrog**_. In an effort to make the CLI usage easier and more convenient, we recently exposed a series of new installers, which install JFrog CLI with the new _**jf**_ executable name. For backward compatibility, the old installers will remain available. We recommend however migrating to the newer _**jf**_ executable name.

### JFrog CLI v2 "jf" installers

The following installers are available for JFrog CLI v2. These installers make JFrog CLI available through the _**jf**_ executable.

#### Debian

```bash
wget -qO - https://releases.jfrog.io/artifactory/jfrog-gpg-public/jfrog\_public\_gpg.key | sudo apt-key add -
echo "deb https://releases.jfrog.io/artifactory/jfrog-debs xenial contrib" | sudo tee -a /etc/apt/sources.list;
apt update;
apt install -y jfrog-cli-v2-jf;
```

#### RPM

```bash
echo "\[jfrog-cli\]" > jfrog-cli.repo;
echo "name=jfrog-cli" >> jfrog-cli.repo;
echo "baseurl=https://releases.jfrog.io/artifactory/jfrog-rpms" >> jfrog-cli.repo;
echo "enabled=1" >> jfrog-cli.repo;
rpm --import https://releases.jfrog.io/artifactory/jfrog-gpg-public/jfrog\_public\_gpg.key
sudo mv jfrog-cli.repo /etc/yum.repos.d/;
yum install -y jfrog-cli-v2-jf;
```

#### Homebrew

```
brew install jfrog-cli
```

#### Install with cUrl

```
url -fL https://install-cli.jfrog.io | sh
```

#### Download with cUrl

```
curl -fL https://getcli.jfrog.io/v2-jf | sh
```

#### NPM

```
npm install -g jfrog-cli-v2-jf
```

#### Docker

```
Slim: docker run releases-docker.jfrog.io/jfrog/jfrog-cli-v2-jf jf -v Full: docker run releases-docker.jfrog.io/jfrog/jfrog-cli-full-v2-jf jf -v
```

#### Powershell

```
powershell: "Start-Process -Wait -Verb RunAs powershell '-NoProfile iwr https://releases.jfrog.io/artifactory/jfrog-cli/v2-jf/\[RELEASE\]/jfrog-cli-windows-amd64/jf.exe -OutFile $env:SYSTEMROOT\\system32\\jf.exe'"
```

#### Chocolatey

```
choco install: jfrog-cli-v2-jf
```

### JFrog CLI v2 "jfrog" installers

The following installers are available for JFrog CLI v2. These installers make JFrog CLI available through the _**jfrog**_ executable.

#### Debian

```
wget -qO - https://releases.jfrog.io/artifactory/jfrog-gpg-public/jfrog\_public\_gpg.key | sudo apt-key add - echo "deb https://releases.jfrog.io/artifactory/jfrog-debs xenial contrib" | sudo tee -a /etc/apt/sources.list; apt update; apt install -y jfrog-cli-v2;
```

#### RPM

```
echo "\[jfrog-cli\]" > jfrog-cli.repo; echo "name=jfrog-cli" >> jfrog-cli.repo; echo "baseurl=https://releases.jfrog.io/artifactory/jfrog-rpms" >> jfrog-cli.repo; echo "enabled=1" >> jfrog-cli.repo; rpm --import https://releases.jfrog.io/artifactory/jfrog-gpg-public/jfrog\_public\_gpg.key sudo mv jfrog-cli.repo /etc/yum.repos.d/; yum install -y jfrog-cli-v2;
```

#### Homebrew

```
brew install jfrog-cli
```

#### Download with Curl

```
curl -fL https://getcli.jfrog.io/v2 | sh
```

#### NPM

```
npm install -g jfrog-cli-v2
```

#### Docker

```
Slim: docker run releases-docker.jfrog.io/jfrog/jfrog-cli-v2 jfrog -v Full: docker run releases-docker.jfrog.io/jfrog/jfrog-cli-full-v2 jfrog -v
```

#### Chocolatey

```
choco install jfrog-cli
```

### JFrog CLI v1 (legacy) installers

The following installations are available for JFrog CLI v1. These installers make JFrog CLI available through the _**jfrog**_ executable.

#### Debian

```
wget -qO - https://releases.jfrog.io/artifactory/jfrog-gpg-public/jfrog\_public\_gpg.key | sudo apt-key add - echo "deb https://releases.jfrog.io/artifactory/jfrog-debs xenial contrib" | sudo tee -a /etc/apt/sources.list; apt update; apt install -y jfrog-cli;
```

#### RPM

```
echo "\[jfrog-cli\]" > jfrog-cli.repo; echo "name=jfrog-cli" >> jfrog-cli.repo; echo "baseurl=https://releases.jfrog.io/artifactory/jfrog-rpms" >> jfrog-cli.repo; echo "enabled=1" >> jfrog-cli.repo; rpm --import https://releases.jfrog.io/artifactory/jfrog-gpg-public/jfrog\_public\_gpg.key sudo mv jfrog-cli.repo /etc/yum.repos.d/; yum install -y jfrog-cli;
```

#### Download with cUrl

```
curl -fL https://getcli.jfrog.io | sh
```

#### NPM

```
npm install -g jfrog-cli-go
```

#### Docker

```
Slim: docker run releases-docker.jfrog.io/jfrog/jfrog-cli jfrog -v Full: docker run releases-docker.jfrog.io/jfrog/jfrog-cli-full jfrog -v
```

#### Go

```
GO111MODULE=on go get github.com/jfrog/jfrog-cli; if \[ -z "$GOPATH" \] then binPath="$HOME/go/bin"; else binPath="$GOPATH/bin"; fi; mv "$binPath/jfrog-cli" "$binPath/jfrog"; echo "$($binPath/jfrog -v) is installed at $binPath";
```
