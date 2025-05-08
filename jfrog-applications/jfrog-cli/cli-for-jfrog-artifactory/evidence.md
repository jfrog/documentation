# Evidence

## Overview

This page describes how to use JFrog CLI with [Evidence](https://jfrog.com/help/r/jfrog-artifactory-documentation/promote-a-release-bundle-v2-to-a-target-environment).

***

**Note**

> JFrog CLI enables creation for custom evidence.

> For now, an evidence can be signed and verified with one key only.

> Evidence is only available with [Artifactory 7.xx.x](https://jfrog.com/help/r/jfrog-release-information/artifactory-7.63.2-cloud) or above.

***

## Syntax

When used with JFrog Evidence, JFrog CLI uses the following syntax:

```
jf evd attest --predicate file-path --predicate-type predicate-type-uri --subjects <target-path[@digest];[target-path[@digest]]> --key <local-private-key-path> [--name <evidence-name>] [--override]
```

### **Command parameters**

1. `--predicate` file-path
    
    Mandatory field.
    Is the path to a locally stored json file which contains the predicates
    ```json
    {
      // any kind of a valid json
    }
    ```
2. `--predicate-type` predicate-type-uri

    Mandatory field.
    The uri of the predicate type, ie:

        https://in-toto.io/attestation/link/v0.3
        https://in-toto.io/attestation/scai/attribute-report
        https://in-toto.io/attestation/runtime-trace/v0.1
        https://in-toto.io/attestation/test-result/v0.1
        https://in-toto.io/attestation/vulns


3. `--subjects` target-path[@digets][;]other-target-path[@digets]


    Mandatory field.

    A semicolon(;) separated list.

    Must include path.

    Digest (sha256) is optional.

    If digest provided it is verified against the paths sha256 in artifactory.

    If digest is not provided sha256 is taken from the path in artifactory.


4. `--key` local-private-key-path

    Mandatory path for a private key.

    Supported keys are:

        `rsa`
        `ed25519`
        `ecdsa`

5. `--name`

    Optional field.

    If name is provided it will be saved in artifactory with the provided name (including .json.evd suffix).

    If name is not provided it will be saved in artifactory with the default name evidence.json.evd

6. `--override`

    Optional field.
    
    If `--override` param is provided then the to be saved evidence will be saved on artifactory regardless if the file exists.
    
    If `--override` param is not provided then the to be saved evidence will fail to be saved if a file with the same name already exists on artifactory.
