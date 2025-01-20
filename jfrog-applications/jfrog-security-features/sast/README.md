# SAST

***

### What is JFrog SAST?

<table data-view="cards"><thead><tr><th align="center"></th><th></th><th></th></tr></thead><tbody><tr><td align="center"><em><strong>FAST</strong></em></td><td>Easy setup with just a few steps</td><td></td></tr><tr><td align="center"><em><strong>FOCUSED</strong></em></td><td>Deep analysis with focus on security issues</td><td></td></tr><tr><td align="center"><em><strong>LOCAL SECURE</strong></em></td><td>Work locally in your environment</td><td></td></tr></tbody></table>

The JFrog SAST solution aims to enable software developers to hunt, fix, and learn about security issues in their code while allowing them to deliver fast, quality code, and reduce issues.

**Requirements**

SAST is available for Enterprise X and Enterprise + subscriptions with Advanced Security.

### Supported Languages:

*   **JavaScript/TypeScript:** Supports JS up to ECMAScript 2016/ES7.  TypeScript is also supported.&#x20;

    Note: JSX is not supported
* **Python:** Supports Python 3.0
* **Java:** Supports all Java versions
* **Go**: Supports all Go versions

### What does JFrog SAST Scan for?

JFrog SAST scans mainly for specific sensitive operations (DB queries, OS commands, outgoing connection destinations, etc) that can be controlled by an external attacker without proper sanitation injections such as:

* SQL injections
* Command injections
* Code injections
* SSRF

It also detects cases when certain APIs (encryption, cryptographic signing, file operations, etc.) are used with parameters or under circumstances that render the API use unsafe.

SAST findings are presented in a way that will help you easily locate the vulnerable data flow in your code. The data is represented within an easy-to-use interface that enables you to track each vulnerability in the code and provides the following information per vulnerability:

* **Data Flow Analysis**: Provides information on the overall code flow and the different entry points of the vulnerability up to the execution point of the vulnerability. At JFrog we understand the developers need to see the entire picture of their code, rather than just providing the specific vulnerability found in the code. With Data Analysis Flow you will be able to follow the entire lifecycle of the vulnerability.
*   **Fix Steps**: To help you fix the security issues, the JFrog security team provides you with detailed fixes and mitigation options for the vulnerabilities. Xray empowers you to make smart choices when creating the mitigation plan and choosing the paths with the highest return on investment.

    Along with the JFrog severity given, you can make informed decisions on what vulnerabilities are a priority to fix. For example, vulnerabilities with low JFrog security severity are considered less risky, as it would be very unlikely to exploit them in the real world, or the impact of the exploitation is low.

### Why JFrog SAST?

#### Discover Zero-day Vulnerabilities in your Code

JFrog SAST is focused on finding 0-day vulnerabilities in your code projects. With a strong balance between FAST execution, QUALITY of scanners, and focusing on SECURITY, JFrog SAST keeps you secured without harming the development process.

#### Early Detection with JFrog Integrations

Incorporate JFrog's scan capabilities into your CI-CD pipeline with JFrog IDE Plugins, JFrog CLI, and Frogbot to continuously monitor your source-code projects and find issues as-early-as-possible in the SDLC.

#### Keep Your Source-Code Local and Secured

JFrog SAST is running locally on your developer's workstation --> no source code/data is leaving your premise for the analysis.

***
