# Infrastructure as Code (IaC)

With the rise in the use of Infrastructure-as-Code (IaC) files, the likelihood of human error is higher than ever. Securing your IaC files by checking the configurations is critical to keeping your cloud deployment safe and secure. JFrog's IaC security scanner is a vital tool and provides a comprehensive, proactive solution to your IaC security concerns.

What do we look for?

* Insufficient access restrictions to services (public access to repositories, publicly accessible clusters, globally readable/deletable/writeable buckets, use of admin roles in ECS services, IAM users with privileged access to all resources, enforce authorization for all API Gateway methods)
* Insecure use of credentials (use of hardcoded credentials)
* Allowing weak crypto algorithms (use of weak cipher suites)
* Running batches in privileged mode
* Enforcement of secure communication (listening to HTTP, unencrypted communications)
* Wildcard actions in Glue policies
* Missing logging (e.g., found CloudTrail trails with logging disabled)
* Disabled upgrades (e.g., RDS database instance with disabled minor engine upgrades)
* Data at rest encryption enablement for Kinesis streams

[IaC in the JFrog CLI for Xray](../../jfrog-cli/cli-for-jfrog-security/)

[IaC in your IDE](../../ide/)

[IaC in your Git Repositories (Frogbot)](../../frogbot/)

