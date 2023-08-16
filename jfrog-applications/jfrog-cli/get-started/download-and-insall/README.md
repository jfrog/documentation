# Download and Insall

To download the executable, please visit the [JFrog CLI Download Site](https://www.jfrog.com/getcli/).

You can also download the sources from the [JFrog CLI Project](https://github.com/JFrog/jfrog-cli-go) on GitHub where you will also find instructions on how to build JFrog CLI.

The legacy name of JFrog CLI's executable is _**jfrog**_. In an effort to make the CLI usage easier and more convenient, we recently exposed a series of new installers, which install JFrog CLI with the new _**jf**_ executable name. For backward compatibility, the old installers will remain available. We recommend however migrating to the newer _**jf**_ executable name.

**JFrog CLI v2 "jf" installers**

The following installers are available for JFrog CLI v2. These installers make JFrog CLI available through the _**jf**_ executable.

**Debian**

```bash
wget -qO - https://releases.jfrog.io/artifactory/jfrog-gpg-public/jfrog\_public\_gpg.key | sudo apt-key add -
echo "deb https://releases.jfrog.io/artifactory/jfrog-debs xenial contrib" | sudo tee -a /etc/apt/sources.list;
apt update;
apt install -y jfrog-cli-v2-jf;
```

**RPM**

```bash
echo "\[jfrog-cli\]" > jfrog-cli.repo;
echo "name=jfrog-cli" >> jfrog-cli.repo;
echo "baseurl=https://releases.jfrog.io/artifactory/jfrog-rpms" >> jfrog-cli.repo;
echo "enabled=1" >> jfrog-cli.repo;
rpm --import https://releases.jfrog.io/artifactory/jfrog-gpg-public/jfrog\_public\_gpg.key
sudo mv jfrog-cli.repo /etc/yum.repos.d/;
yum install -y jfrog-cli-v2-jf;
```

**Homebrew**

```
brew install jfrog-cli
```

**Install with cUrl**

```
url -fL https://install-cli.jfrog.io | sh
```

**Download with cUrl**

```
curl -fL https://getcli.jfrog.io/v2-jf | sh
```

**NPM**

```
npm install -g jfrog-cli-v2-jf
```

**Docker**

```
Slim: docker run releases-docker.jfrog.io/jfrog/jfrog-cli-v2-jf jf -v Full: docker run releases-docker.jfrog.io/jfrog/jfrog-cli-full-v2-jf jf -v
```

**Powershell**

```
powershell: "Start-Process -Wait -Verb RunAs powershell '-NoProfile iwr https://releases.jfrog.io/artifactory/jfrog-cli/v2-jf/\[RELEASE\]/jfrog-cli-windows-amd64/jf.exe -OutFile $env:SYSTEMROOT\\system32\\jf.exe'"
```

**Chocolatey**

```
choco install: jfrog-cli-v2-jf
```

**JFrog CLI v2 "jfrog" installers**

The following installers are available for JFrog CLI v2. These installers make JFrog CLI available through the _**jfrog**_ executable.

**Debian**

```
wget -qO - https://releases.jfrog.io/artifactory/jfrog-gpg-public/jfrog\_public\_gpg.key | sudo apt-key add - echo "deb https://releases.jfrog.io/artifactory/jfrog-debs xenial contrib" | sudo tee -a /etc/apt/sources.list; apt update; apt install -y jfrog-cli-v2;
```

**RPM**

```
echo "\[jfrog-cli\]" > jfrog-cli.repo; echo "name=jfrog-cli" >> jfrog-cli.repo; echo "baseurl=https://releases.jfrog.io/artifactory/jfrog-rpms" >> jfrog-cli.repo; echo "enabled=1" >> jfrog-cli.repo; rpm --import https://releases.jfrog.io/artifactory/jfrog-gpg-public/jfrog\_public\_gpg.key sudo mv jfrog-cli.repo /etc/yum.repos.d/; yum install -y jfrog-cli-v2;
```

**Homebrew**

```
brew install jfrog-cli
```

**Download with Curl**

```
curl -fL https://getcli.jfrog.io/v2 | sh
```

**NPM**

```
npm install -g jfrog-cli-v2
```

**Docker**

```
Slim: docker run releases-docker.jfrog.io/jfrog/jfrog-cli-v2 jfrog -v Full: docker run releases-docker.jfrog.io/jfrog/jfrog-cli-full-v2 jfrog -v
```

**Chocolatey**

```
choco install jfrog-cli
```

**JFrog CLI v1 (legacy) installers**

The following installations are available for JFrog CLI v1. These installers make JFrog CLI available through the _**jfrog**_ executable.

**Debian**

```
wget -qO - https://releases.jfrog.io/artifactory/jfrog-gpg-public/jfrog\_public\_gpg.key | sudo apt-key add - echo "deb https://releases.jfrog.io/artifactory/jfrog-debs xenial contrib" | sudo tee -a /etc/apt/sources.list; apt update; apt install -y jfrog-cli;
```

**RPM**

```
echo "\[jfrog-cli\]" > jfrog-cli.repo; echo "name=jfrog-cli" >> jfrog-cli.repo; echo "baseurl=https://releases.jfrog.io/artifactory/jfrog-rpms" >> jfrog-cli.repo; echo "enabled=1" >> jfrog-cli.repo; rpm --import https://releases.jfrog.io/artifactory/jfrog-gpg-public/jfrog\_public\_gpg.key sudo mv jfrog-cli.repo /etc/yum.repos.d/; yum install -y jfrog-cli;
```

**Download with cUrl**

```
curl -fL https://getcli.jfrog.io | sh
```

**NPM**

```
npm install -g jfrog-cli-go
```

**Docker**

```
Slim: docker run releases-docker.jfrog.io/jfrog/jfrog-cli jfrog -v Full: docker run releases-docker.jfrog.io/jfrog/jfrog-cli-full jfrog -v
```

**Go**

```
GO111MODULE=on go get github.com/jfrog/jfrog-cli; if \[ -z "$GOPATH" \] then binPath="$HOME/go/bin"; else binPath="$GOPATH/bin"; fi; mv "$binPath/jfrog-cli" "$binPath/jfrog"; echo "$($binPath/jfrog -v) is installed at $binPath";
```
