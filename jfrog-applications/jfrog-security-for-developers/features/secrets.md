# Secrets

JFrog's secrets detection searches for known structures and completely random credentials (using suspicious variable matching), ensuring that our detection engines generate minimal false positives.

JFrog Security uses a constantly updated list of more than 150 specific types of credentials. In addition, JFrog Security uses a proprietary generic secrets matcher, for the best coverage possible. It also scans for issues in the certificates used in the software, such as expired or weak certificates.

JFrog Secrets detection can detect the following types of secrets:

* [Access tokens (keys)](https://jfrog.com/help/r/6nte66fuu2ZQMB2dfriysg/9wGJsqb\~chJSbw\~zb\_\_7fg?section=UUID-9380cbae-3ae5-b761-96d0-64b171ecd499\_N1722188830347)
* [Certificates/private keys](https://jfrog.com/help/r/6nte66fuu2ZQMB2dfriysg/9wGJsqb\~chJSbw\~zb\_\_7fg?section=UUID-9380cbae-3ae5-b761-96d0-64b171ecd499\_N1722188840357)
* [High entropy secrets](https://jfrog.com/help/r/6nte66fuu2ZQMB2dfriysg/9wGJsqb\~chJSbw\~zb\_\_7fg?section=UUID-9380cbae-3ae5-b761-96d0-64b171ecd499\_N1722188849592)
* [URL secrets](https://jfrog.com/help/r/6nte66fuu2ZQMB2dfriysg/9wGJsqb\~chJSbw\~zb\_\_7fg?section=UUID-9380cbae-3ae5-b761-96d0-64b171ecd499\_N1722188859111)&#x20;

### Supported File Types:

In the IDE and Frogbot only textual files are scanned.

In the CLI there are commands such as ‘jf audit’ that scan source code and look for secrets in textual files, and other commands such as ‘jf docker scan’ that scan both binary and textual files.



[Secrets Detection in the JFrog CLI for Xray](../../jfrog-cli/cli-for-jfrog-security/)

[Secrets Detection in your IDE](../../ide/)

[Secrets Detection in your Git Repositories ](../../frogbot/)[(Frogbot)](../../frogbot/)
