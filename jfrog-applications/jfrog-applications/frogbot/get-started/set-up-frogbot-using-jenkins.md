# Set up Frogbot Using Jenkins

#### 🖥️ Follow these steps to install Frogbot on Jenkins

<details>

<summary>1️⃣ Install Jenkins 'Generic Webhook Trigger' plugin</summary>

From your Jenkins dashboard navigate to **Manage Jenkins** > **Manage Plugins** and select the **Available** tab. Use the search bar to find **Generic Webhook Trigger** ([more info](https://plugins.jenkins.io/generic-webhook-trigger/)).

</details>

***

<details>

<summary>2️⃣ Connect the Webhook on your Git provider</summary>



</details>

***

<details>

<summary>3️⃣ Optional - setting JobToken</summary>

* When using the plugin in several jobs, you will have the same URL trigger all jobs. If you want to trigger only a certain job you can use the **JobToken** in the URL to specify what job needs to be executed.
* Webhook URL with **JobToken** : `JENKINS_URL/generic-webhook-trigger/invoke?token=MyJobToken`
* On some Git providers the JobToken is called Secret Token.
* Read more [JobToken Docs](https://plugins.jenkins.io/generic-webhook-trigger/#plugin-content-trigger-only-specific-job)

</details>

***

<details>

<summary>4️⃣ Set up credentials</summary>

* Set up the following credentials using Jenkins credentials functionality, as **Secret Text**:
  * **JF\_URL** - JFrog Platform URL (Example: "https://acme.jfrog.io")
  * **JF\_ACCESS\_TOKEN** _or_ **JF\_USER** & **JF\_PASSWORD** - JFrog Credentials
  * **JF\_GIT\_TOKEN** - access token with read\&write access to the Git repository
* [How to use credentials with Jenkins](https://www.jenkins.io/doc/book/using/using-credentials/)

</details>

***

<details>

<summary>5️⃣ Prepare Jenkins Agent</summary>

* It is essential to have the appropriate package manager used by the scanned project installed on the Jenkins Agent. For instance, if the project uses an npm project, you need to have the npm client installed.

</details>

***

<details>

<summary>6️⃣ Scanning pull requests</summary>

Create a new pipeline job using this Jenkinsfile template.



Enable the ‘Generic Webhook Trigger’:



</details>

***

<details>

<summary>7️⃣ Scanning repository branches and fixing issues</summary>

Create a new pipeline job using this Jenkinsfile template.



</details>
