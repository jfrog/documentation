# Secrets

JFrog's secrets detection searches for known structures and completely random credentials (using suspicious variable matching), ensuring that our detection engines generate minimal false positives.

JFrog Security uses a constantly updated list of more than 150 specific types of credentials. In addition, JFrog Security uses a proprietary generic secrets matcher, for the best coverage possible. It also scans for issues in the certificates used in the software, such as expired or weak certificates.

**Examples**:

* Use of expired certificates
* Inclusion of plaintext API keys, private keys

[Secrets Detection in the JFrog CLI for Xray](../../jfrog-applications-1/jfrog-cli/cli-for-jfrog-security/)

[Secrets Detection in your IDE](../../jfrog-applications-1/ide/)

[Secrets Detection in your Git Repositories ](../../jfrog-applications-1/frogbot/)[(Frogbot)](../../jfrog-applications-1/frogbot/)
