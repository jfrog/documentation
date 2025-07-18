# Evidence Service

## Overview

This page describes how to use the JFrog CLI to create external [evidence](https://jfrog.com/help/r/jfrog-artifactory-documentation/evidence-management) files, which are then deployed to Artifactory. You can create evidence for:

* Artifacts
* Packages
* Builds
* Release Bundles v2

***

**Note**

* The Evidence service requires Artifactory 7.104.2 or above.
* The ability for users to attach external evidence to Artifactory, as described here, requires an **Enterprise+** subscription.
* The ability to collect internal evidence generated by Artifactory requires a **Pro** subscription or above. Internal evidence generated by Xray requires a **Pro X** subscription or above.
* In the current release, an evidence file can be signed with one key only.
* The maximum size evidence file supported by Artifactory is 16MB.

Click this [link](https://github.com/jfrog/Evidence-Examples/tree/main/examples) for a collection of code snippets that describe how to create evidence workflows in various tools using the JFrog CLI. &#x20;

For more information about the API used for deploying evidence to Artifactory, see [Deploy Evidence](https://jfrog.com/help/r/jfrog-rest-apis/deploy-evidence).

***

### Authentication

To deploy external evidence, use an access token or the web login mechanism for authentication. Basic authentication (username/password) is not supported.

### Syntax

JFrog CLI uses the following syntax for creating evidence:

**Artifact Evidence**

{% code overflow="wrap" %}
```
jf evd create --predicate file-path --predicate-type predicate-type-uri --subject-repo-path <target-path> --subject-sha256 <digest> --key <local-private-key-path> --key-alias <public-key-name>
```
{% endcode %}

**Package Evidence**

{% code overflow="wrap" %}
```
jf evd create --predicate file-path --predicate-type predicate-type-uri --package-name <name> --package-version <version-number> --package-repo-name <repo-name> --key <local-private-key-path> --key-alias <public-key-name>
```
{% endcode %}

**Build Evidence**

{% code overflow="wrap" %}
```
jf evd create --predicate file-path --predicate-type predicate-type-uri --build-name <name> --build-number <version-number> --key <local-private-key-path> --key-alias <public-key-name>
```
{% endcode %}

**Release Bundle v2 Evidence**

{% code overflow="wrap" %}
```
jf evd create --predicate file-path --predicate-type predicate-type-uri --release-bundle <name> --release-bundle-version <version-number> --key <local-private-key-path> --key-alias <public-key-name>
```
{% endcode %}

## **Command parameters**

1. `--predicate` file-path\
   Mandatory field.\
   Defines the path to a locally-stored, arbitrary json file that contains the predicates.

```
{
// any kind of valid json
}
```

2. `--predicate-type` predicate-type-uri\
   Mandatory field.\
   The type of predicate defined by the json file. Sample predicate type uris include:

```
 https://in-toto.io/attestation/link/v0.3
 https://in-toto.io/attestation/scai/attribute-report
 https://in-toto.io/attestation/runtime-trace/v0.1
 https://in-toto.io/attestation/test-result/v0.1
 https://in-toto.io/attestation/vulns
```

3. `--key` local-private-key-path\
   Path to a private key (see Tip below). Supported key types include:

```
 `rsa`
 `ed25519`
 `ecdsa`
```

***

**Tip**

> As an alternative to using the `--key` command parameter, you can define the key using the `JFROG_CLI_SIGNING_KEY` environment variable. If the environment variable is **not** defined, the `--key` command is mandatory.

***

***

**Note**

> Two key formats are supported: PEM and SSH

***

4. `--key-alias` RSA-1024\
   Optional case-sensitive name for the public key created from the private key. The public key is used to verify the DSSE envelope that contains the evidence.
   * If the `key-alias` is included, DSSE verification will fail if the same `key-name` is not found in Artifactory.
   * If the `key-alias` is not included, DSSE verification with the public key is not performed during creation.

***

**Tip**

> You can define a key alias using the `JFROG_CLI_KEY_ALIAS` environment variable as an alternative to using the `--key-alias` command parameter.

***

***

**Note**

> In the unlikely event the public key is deleted from Artifactory, it may take up to 4 hours for the Evidence service to clear the key from the cache. Evidence can still be signed with the deleted key during this time.

***

5. `--markdown`  md file\
   Optional path to a file that contains evidence formatted in markdown.
6. `--project` project-name\
   Optional name of the project associated with the evidence subject. This argument can be used with build, package, and Release Bundle evidence.

### Artifact command parameters

1. `--subject-repo-path` target-path\
   Mandatory field. \
   Each evidence file must have a single subject only and must include the path. Artifacts located in local repositories aggregated inside virtual repositories are supported (evidence is added to the local path).
2. `--subject-sha256` digest\
   Optional digest (sha256) of the artifact.

* If a digest is provided, it is verified against the subject's sha256 as it appears in Artifactory.
* If a digest is not provided, the sha256 is taken from the path in Artifactory.

### Package command parameters

1. `--package-name` name\
   Mandatory field.
2. `--package-version` version-number\
   Mandatory field.
3. `--package-repo-name` repo-name\
   Mandatory field.

### Build command parameters

1. `--build-name` name\
   Mandatory field unless environment variables are used (see tip below).
2. `--build-number` version-number\
   Mandatory field unless environment variables are used (see tip below).

***

**Tip**

> You can use the `FROG_CLI_BUILD_NAME` and `FROG_CLI_BUILD_NUMBER` environment variables as an alternative to the build command parameters.

***

### Release Bundle v2 command parameters

1. `--release-bundle` name\
   Mandatory field.
2. `--release-bundle-version` version-number\
   Mandatory field.

***

**Note**

> When DSSE verification is successful, the following message is displayed:

```
Evidence successfully created and verified.
```

> When DSSE verification is unsuccessful, the following message is displayed:

```
Evidence successfully created but not verified due to missing/invalid public key.
```

***

### Sample commands

**Artifact Evidence Sample**

{% code overflow="wrap" %}
```
evd create --predicate /Users/jsmith/Downloads/code-review.json --predicate-type https://in-toto.io/attestation/vulns --subject-repo-path commons-dev-generic-local/commons/file.txt --subject-sha256 69d29925ba75eca8e67e0ad99d1132b47d599c206382049bc230f2edd2d3af30 --key /Users/jsmith/Documents/keys/private.pem --key-alias xyzey
```
{% endcode %}

In the sample above, the command creates a signed evidence file with a predicate type of SLSA provenance for an artifact named **file.txt**.

**Package Evidence Sample**

{% code overflow="wrap" %}
```
evd create --predicate /Users/jsmith/Downloads/code-review.json --predicate-type https://in-toto.io/attestation/vulns --package-name DockerPackage --package-version 1.0.0 --package-repo-key local-docker --key /Users/jsmith/Documents/keys/private.pem --key-alias xyzey
```
{% endcode %}

**Build Evidence Sample**

{% code overflow="wrap" %}
```
evd create --predicate /Users/jsmith/Downloads/code-review.json --predicate-type https://in-toto.io/attestation/vulns --package-name DockerPackage --package-version 1.0.0 --package-repo-name local-docker --key /Users/jsmith/Documents/keys/private.pem --key-alias xyzey
```
{% endcode %}

**Release Bundle v2 Evidence Sample**

{% code overflow="wrap" %}
```
evd create --predicate /Users/jsmith/Downloads/code-review.json --predicate-type https://in-toto.io/attestation/vulns --release-bundle bundledemo --release-bundle-version (mandatory) 1.0.0 --key /Users/jsmith/Documents/keys/private.pem --key-alias xyzey
```
{% endcode %}
