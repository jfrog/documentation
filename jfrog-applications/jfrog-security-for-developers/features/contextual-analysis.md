# Contextual Analysis

The JFrog security research team is dedicated to exploring the intricacies of vulnerabilities, analyzing new attack methods, and crafting advanced techniques to determine their applicability.

Shifting left using Contextual Analysis enables you to eliminate false positive reports on vulnerabilities that are not applicable. This process involves automated scanners running on top of the container to find reachable paths for the analyzed vulnerabilities. Xray automatically validates some high and very high-impact vulnerabilities, such as vulnerabilities that have prerequisites for exploitations, and provides contextual analysis information for these vulnerabilities, to assist you in figuring out which vulnerabilities are applicable to a specific artifact.

**What are the benefits of Vulnerability Contextual Analysis?**

* Analyzes the finished code the way an attacker would. Know what issues are exploitable and their potential impact.
* Tests an issue in the context of the complete artifact, also within a build or Release Bundle.
* Enables action and remediation in the context of the actual artifact, build or Release Bundle.

[Contextual Analysis in the JFrog CLI for Xray](../../jfrog-cli/cli-for-jfrog-security/)

[Contextual Analysis in your IDE](../../ide/)

[Contextual Analysis in your Git Repositories (FrogBot)](../../frogbot/)

